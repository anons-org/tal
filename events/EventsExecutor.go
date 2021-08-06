package events

import (
	"fmt"
	"tal/structs"
)

type EventsExecutor struct {
	
}

//收到群聊消息
func (me *EventsExecutor )OnGroupMessage(e structs.Event)  bool{
	//if event.MessageType!="group" && event.PostType!="message"{
	//	return false
	//}
	fmt.Println(fmt.Sprintf("群消息:发送人%d 消息%s",e.Sender.UserId,e.Message))
	return true;
}