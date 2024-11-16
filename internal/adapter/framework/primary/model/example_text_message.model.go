package model

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

type MessageTextItem struct {
	Id       string `json:"id" form:"id"`
	Sender   string `json:"sender" form:"sender"`
	Receiver string `json:"receiver" form:"receiver"`
	Message  string `json:"message" form:"message"`
}

func FromMessageTextItemEntity(m entity.MessageTextItem) MessageTextItem {
	return MessageTextItem{
		Id:       m.Id,
		Sender:   m.Sender,
		Receiver: m.Receiver,
		Message:  m.Message,
	}
}

type SendTextMessageRequest struct {
	Sender   string `json:"sender" form:"sender" validate:"required"`
	Receiver string `json:"receiver" form:"receiver" validate:"required"`
	Message  string `json:"message" form:"message" validate:"required"`
}

func (s SendTextMessageRequest) ToEntity() entity.SendTextMessageRequest {
	return entity.SendTextMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
	}
}

func FromSendTextMessageRequestEntity(s entity.SendTextMessageRequest) SendTextMessageRequest {
	return SendTextMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
	}
}
