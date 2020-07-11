package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "./chatRoom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

//系統訊息名稱
const systemName string = "System"

//暫存己登入的使用者, key:token, value:name
var userMap map[string]string

//暫存使用的取得聊天訊息的連線, key:token, value:stream Service
// var userStreamMap map[string]pb.ChatRoomService_ChatServer
var userStreamMap sync.Map

type server struct{}

//登入, UserVo.Token有值時表示登入成功
func (s *server) Login(ctx context.Context, in *pb.LoginUserVo) (*pb.UserVo, error) {
	//登入名稱重覆檢查, 名字不可跟系統名同名
	if in.GetName() == systemName {
		return &pb.UserVo{Name: in.GetName()}, nil
	}
	//登入名稱重覆檢查, 名字不可重覆
	for key := range userMap {
		if strings.Compare(userMap[key], in.GetName()) == 0 {
			return &pb.UserVo{Name: in.GetName()}, nil
		}
	}

	//名字不重覆時, 產生Token
	token := generateToken()
	userMap[token] = in.GetName()

	//回傳登入結果
	return &pb.UserVo{Token: token, Name: in.GetName()}, nil
}

//聊天
func (s *server) Chat(stream pb.ChatRoomService_ChatServer) error {
	//取出連線資訊檢查是否為己登入的使用者
	md, _ := metadata.FromIncomingContext(stream.Context())
	token := md["token"][0]
	userName, exist := userMap[token]
	if !exist {
		return errors.New("user not login")
	}

	//將連線放入使用者的連線池裡
	userStreamMap.Store(token, stream)

	//如果連線結束時, 刪除使用者的暫存資料及連線, 廣播離開訊息
	go func() {
		<-stream.Context().Done()
		userStreamMap.Delete(token)
		delete(userMap, token)
		doSendMessageToAll(token, systemName, fmt.Sprintf("%s exit the chat room.", userName), time.Now())
	}()

	//廣播歡迎訊息
	doSendMessageToAll(systemName, systemName, fmt.Sprintf("Welcome %s", userName), time.Now())

	//監聽使用者發送的訊息, 並廣播給其他使用者
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		doSendMessageToAll(token, userName, req.GetContent(), time.Now())
	}
}

//產出Token
func generateToken() string {
	now := time.Now().Unix()
	rand.Seed(now)
	randNum := rand.Intn(100)
	return strconv.FormatInt(now, 16) + strconv.Itoa(randNum)
}

//GetUserStream return ChatRoomService_ChatServer
//取得使用者聊天連線
func GetUserStream(token string) pb.ChatRoomService_ChatServer {
	if stream, ok := userStreamMap.Load(token); ok {
		return stream.(pb.ChatRoomService_ChatServer)
	}
	return nil
}

/*
	廣播訊息
	token : 訊息發出者的token, 不廣播
	name : 發出訊息者的名字
	content : 訊息內容
	now : 發送時間
*/
func doSendMessageToAll(token string, name string, content string, now time.Time) {
	//訊息格式 hh:mm userName: 訊息內容
	msg := fmt.Sprintf("%s %s: %s", now.AppendFormat([]byte(""), time.Kitchen), name, content)
	// fmt.Println(msg)

	//不廣播給指定的token
	for userToken := range userMap {
		if userToken == token {
			continue
		}

		stream := GetUserStream(userToken)
		if stream != nil {
			stream.Send(&pb.MessageVo{Content: msg})
		}
	}
}

func main() {
	//初始化
	userMap = make(map[string]string)

	//指定傾聽器使用tcp協定及Port
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create grpc server
	s := grpc.NewServer()
	//將grpc service api的實作注入到grpc server
	pb.RegisterChatRoomServiceServer(s, &server{})
	//供grpcurl查看grpc server上的grpc service api
	reflection.Register(s)
	//為grpc server指定傾聽器start server
	log.Fatalf("failed to serve: %v", s.Serve(lis))
}
