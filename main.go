package tal

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MessageSender struct {
	UserId int `json:"user_id"`
}

type Event struct {
	MessageType string `json:"message_type"`
	PostType string `json:"post_type"`
	RawMessage string `json:"raw_message"`
	Message string `json:"message"`
	Sender MessageSender
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {

		data,err := c.GetRawData()
		if err != nil{
			fmt.Println(err.Error())
		}

		p := &Event{}
		err = json.Unmarshal([]byte(data), p)
		fmt.Println(err)
		fmt.Println(*p)

		if p.MessageType=="group"{
			fmt.Println(fmt.Sprintf("群消息:发送人%d 消息%s",p.Sender.UserId,p.Message))
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