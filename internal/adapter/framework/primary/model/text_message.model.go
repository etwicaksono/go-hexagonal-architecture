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

type MultimediaStorage string

const (
	MultimediaStorage_LOCAL MultimediaStorage = "LOCAL"
	MultimediaStorage_MINIO MultimediaStorage = "MINIO"
)

// Enum value maps for MultimediaStorage.
var (
	MultimediaStorage_self = map[int32]MultimediaStorage{
		0: MultimediaStorage_LOCAL,
		1: MultimediaStorage_MINIO,
	}
)

type SendMultimediaMessageRequest struct {
	Sender   string            `json:"sender" validate:"required"`
	Receiver string            `json:"receiver" validate:"required"`
	Message  string            `json:"message" validate:"required"`
	Storage  MultimediaStorage `json:"storage"`
	Files    []entity.MultimediaFile
}

func (s SendMultimediaMessageRequest) ToEntity() entity.SendMultimediaMessageRequest {
	return entity.SendMultimediaMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		Storage:  entity.MultimediaStorage_self[string(s.Storage)],
		Files:    s.Files,
	}
}

func FromSendMultimediaMessageRequestEntity(s entity.SendMultimediaMessageRequest) SendMultimediaMessageRequest {
	return SendMultimediaMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		Storage:  MultimediaStorage_self[int32(s.Storage)],
		Files:    s.Files,
	}
}
