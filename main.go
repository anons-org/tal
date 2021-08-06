package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"tal/events"
	"tal/structs"
)



func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {

		data,err := c.GetRawData()
		if err != nil{
			fmt.Println(err.Error())
		}

		p := &structs.Event{}
		err = json.Unmarshal([]byte(data), p)
		fmt.Println(err)
		fmt.Println(*p)

		evt:=&events.EventsExecutor{}



		if p.MessageType=="group"{

			switch p.PostType {
			case "message":
				evt.OnGroupMessage(*p)
			}


		}

		//
		//fmt.Printf("data: %v\n",string(data))
		////很关键
		////把读过的字节流重新放到body
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		//
		//fmt.Println(c.Request.Body)
		//message := c.PostForm("message")
		//nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "ok",
		})
	})




	router.Run(":5701")
}