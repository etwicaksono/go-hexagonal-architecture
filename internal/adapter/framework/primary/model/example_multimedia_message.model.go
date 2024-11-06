package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/valueobject"
)

type SendMultimediaMessageRequest struct {
	Sender   string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Message  string `json:"message" validate:"required"`
	Storage  string `json:"storage"`
	Files    []entity.MultimediaFile
}

func (s SendMultimediaMessageRequest) ToEntity() entity.SendMultimediaMessageRequest {
	return entity.SendMultimediaMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		Storage:  valueobject.MultimediaStorageFromString(s.Storage),
		Files:    s.Files,
	}
}

func FromSendMultimediaMessageRequestEntity(s entity.SendMultimediaMessageRequest) SendMultimediaMessageRequest {
	return SendMultimediaMessageRequest{
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		Storage:  s.Storage.ToString(),
		Files:    s.Files,
	}
}

type FileItem struct {
	Storage string `json:"storage"`
	File    string `json:"file"`
}

type MessageMultimediaItem struct {
	Id       string   `json:"id"`
	Sender   string   `json:"sender"`
	Receiver string   `json:"receiver"`
	Message  string   `json:"message"`
	FileUrls []string `json:"fileUrls"`
}

func FromMessageMultimediaItemEntity(s entity.MessageMultimediaItem) MessageMultimediaItem {
	var fileUrls []string

	for _, file := range s.Files {
		fileUrls = append(fileUrls, file.File)
	}

	return MessageMultimediaItem{
		Id:       s.Id,
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		FileUrls: fileUrls,
	}
}
