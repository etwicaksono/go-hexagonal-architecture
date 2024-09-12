package model

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

type MessageTextItem struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
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

func FromSendTextMessageRequestEntity(s entity.SendTextMessageRequest) SendTextMessageRequest {
	return SendTextMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
	}
}
