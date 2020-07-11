package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	pb "../grpc/chatRoom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var serviceClient pb.ChatRoomServiceClient
	var userToken string

	//建立grpc server的連線
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connection : %v", err)
	}
	defer conn.Close()

	//create grpc service client
	serviceClient = pb.NewChatRoomServiceClient(conn)

	//監聽command-line輸入
	input := bufio.NewScanner(os.Stdin)

	//使用者登入
	fmt.Print("Please enter a name\n>")
	for {
		input.Scan()
		line := input.Text()
		userVo, err := serviceClient.Login(context.Background(), &pb.LoginUserVo{Name: line})
		if err != nil {
			log.Fatalf("login fail : %v", err)
		}

		if userVo.GetToken() == "" {
			fmt.Print("Login fail, Please enter a name\n>")
		} else {
			fmt.Print("Login Success, enter :exit will leave.\n>")
			userToken = userVo.GetToken()
			break
		}
	}

	//將登入後取得的token放入context並連線聊天室API
	ctx, cancel := context.WithCancel(context.Background())
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("token", userToken))
	client, err := serviceClient.Chat(ctx)
	if err != nil {
		log.Fatalf("get chat connection fail : %v", err)
	}

	//初始化同步處理
	var wg sync.WaitGroup

	//監聽server端回傳的訊息並將它輸出到畫面上
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			res, err := client.Recv()
			if err != nil {
				log.Fatalf("get chat message fail : %v", err)
			}
			fmt.Printf("\n%s\n>", res.GetContent())
		}
	}()

	//監聽使用者的輸入, 並將它發送到server端
	wg.Add(1)
	go func() {
		defer wg.Done()
		for input.Scan() {
			line := input.Text()
			if line == ":exit" {
				cancel()
				os.Exit(2)
				break
			} else {
				fmt.Print(">")
				err := client.Send(&pb.MessageVo{Content: line})
				if err != nil {
					log.Fatalf("send message fail : %v", err)
				}
			}
		}
	}()

	//等待的同步處理的工作結束
	wg.Wait()
	//關閉連線
	cancel()
}
