package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr","127.0.0.1:8080","http service address")
var upgrader = websocket.Upgrader{}




func home(w http.ResponseWriter,r *http.Request){
	homeTemplate.Execute(w,"ws://" + r.Host + "/socket_test")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/socket_test",socket_test)
	http.HandleFunc("/",home)
	log.Fatal(http.ListenAndServe(*addr,nil))
}
var homeTemplate ,_ = template.ParseFiles("client.html")
func socket_test(w http.ResponseWriter,r *http.Request) {
	c,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Print("upgrade:",err)
		return 
	}
	defer c.Close()
	for {
		//接收消息
		mt,message,err := c.ReadMessage()
		if err != nil{
			log.Println("read:",err)
		}
		log.Printf("recv:%s",message)
		
		// 返回消息
		var str string
		str = "hello,receive:" + string(message[:])
		err = c.WriteMessage(mt,[]byte(str))

		if err != nil {
			log.Println("write:",err)
			break
		}
	}
}
