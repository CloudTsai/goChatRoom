# go + grpc chat room example
執行該範例需要安裝go, 並使用go的指令取得grpc<br>
安裝好go再確認環境變數沒問題後, 再安裝grpc, 在CMD下執行 go get -u google.golang.org/grpc

該範例使用go + grpc來示範一個簡單的聊天室, 使用CMD(windows的命令提示字元)來執行該範例, 需啓動一支server端程式及N支client端<br>
該範例使用的prot是8081

先開聊天室server端, 在cmd裡將路徑指到該專案的grpc資料夾下, 執行go run grpcServer.go啓動server<br>
例如: C:\goChatRoom\grpc>go run grpcServer.go<br>
若是Windows系統需要等一下, 防火牆會跳訊息是否允許存取, 這時需要手動按允取存取, 按完後才算開啓成功

開好server端後, 再開聊天室client端, 在cmd裡將路徑指到該專案的client資料夾下, 執行go run client.go啓動客戶端開始聊天<br>
例如: C:\goChatRoom\client>go run client.go<br>
client端請開二個以上, 只開一個的話因為不會廣播給發言人, 使用上會像是沒有任何反應

client端看到下列的內容時, 請輸入聊天室裡使用的名字後按下enter<br>
Please enter a name<br>
&gt;

輸入完名字後看到 Login Success, enter :exit will leave. 即為登入成功<br>
等看到<br>
hh:mmAM System: Welcome name<br>
&gt;<br>
就可以開始聊天了

client端想離開聊天室時輸入:exit後按下enter即可退出聊天室, 或是按下Ctrl + C

server端想關閉時按下Ctrl + C

