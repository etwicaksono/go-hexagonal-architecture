package entity

import "github.com/etwicaksono/public-proto/gen/example"

type MessageTextItem struct {
	Id       string
	Sender   string
	Receiver string
	Message  string
}

func (mti MessageTextItem) ToProto() *example.MessageTextItem {
	return &example.MessageTextItem{
		Id:       mti.Id,
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
}
