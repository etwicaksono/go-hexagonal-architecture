package model

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

type MessageTextItem struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func FromMessageTextItemEntity(mti entity.MessageTextItem) MessageTextItem {
	return MessageTextItem{
		Id:       mti.Id,
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
}

type SendTextMessageRequest struct {
	Sender   string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Message  string `json:"message" validate:"required"`
}

func (s SendTextMessageRequest) ToEntity() entity.SendTextMessageRequest {
	return entity.SendTextMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
	}
}

func FromSendTextMessageRequestEntity(mti entity.SendTextMessageRequest) SendTextMessageRequest {
	return SendTextMessageRequest{
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
}
