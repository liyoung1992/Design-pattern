package main
import (
	"flag"
	"log"
	"net/url"
	"github.com/gorilla/websocket"
)
var addr = flag.String("addr","localhost:8080","http service addr")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme:"ws",Host:*addr,Path:"/socket_test"}

	log.Printf("connecting to %s",u.String())

	c,_,err := websocket.DefaultDialer.Dial(u.String(),nil)
	if err != nil {
		log.Fatal("dial:",err)
	}
	defer c.Close()
	// 开启线程一直发送消息
	go func() {
		defer c.Close()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	for  {
      var str string = "test"
	
	  var data []byte = []byte(str)

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println("write:", err)
		return
	}
	}

	c.Close()

}