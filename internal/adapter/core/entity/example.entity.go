package entity

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/public-proto/gen/example"
)

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

func (mti MessageTextItem) ToModel() *model.MessageTextItem {
	return &model.MessageTextItem{
		Id:       mti.Id,
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
}

type SendTextMessageRequest struct {
	Sender   string
	Receiver string
	Message  string
}
