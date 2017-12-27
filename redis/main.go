package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
// 	fmt.Println("redigo-demo-start")
// 	// 连接
// 	c,err := redis.Dial("tcp","192.168.79.128:6379")
// 	if err != nil {
// 		fmt.Println(err)
// 		return 
// 	}
//    //set
// 	v,err := c.Do("set","set1","666")
// 	if err != nil {
// 		fmt.Println(err)
// 	    return 
// 	}
// 	fmt.Println(v)
// 	//get
// 	v,err = redis.String(c.Do("get","set1"))
// 	if err != nil{
// 		fmt.Println(err)
// 		return 
// 	}
// 	fmt.Println(v)
// 	//list
// 	c.Do("lpush","list","666")
// 	c.Do("lpush","list","777")
// 	c.Do("lpush","list","888")
	
// 	values,_ := redis.Values(c.Do("lrange","list","0","100"))

// 	for _,v := range values {
// 		fmt.Println(string(v.([]byte)))
// 	}
// 	//管道
// 	c.Send("set","name","liyang")
// 	c.Send("get","name")
// 	c.Flush()
// 	c.Receive()
// 	c.Receive()
	
// 	c.Close()

	go subscribe()
    go subscribe()
    go subscribe()
    go subscribe()
    go subscribe()
 
 c,err := redis.Dial("tcp","192.168.79.128:6379")
 if err != nil {
    fmt.Println(err)
    return
 }
 
 defer c.Close()
 
 for {
     var s string
    fmt.Scanln(&s)
    _, err := c.Do("PUBLISH", "redChatRoom", s)
     if err != nil {
       fmt.Println("pub err: ", err)
        return
   }
 }
}
func subscribe() {
	c, err := redis.Dial("tcp", "192.168.79.128:6379")
    if err != nil {
     fmt.Println(err)
        return
 	 }
    defer c.Close()

    psc := redis.PubSubConn{c}
     psc.Subscribe("redChatRoom")
    for {
         switch v := psc.Receive().(type) {
       case redis.Message:
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
       case redis.Subscription:
           fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
       case error:
          fmt.Println(v)
           return
	   }
	}
}