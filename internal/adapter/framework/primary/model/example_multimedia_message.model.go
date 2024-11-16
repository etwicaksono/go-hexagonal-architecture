package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
)

type SendMultimediaMessageRequest struct {
	Sender   string `json:"sender" form:"sender" validate:"required"`
	Receiver string `json:"receiver" form:"receiver" validate:"required"`
	Message  string `json:"message" form:"message" validate:"required"`
	Storage  string `json:"storage" form:"storage"`
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
	Storage string `json:"storage" form:"storage"`
	File    string `json:"file" form:"file"`
}

type MessageMultimediaItem struct {
	Id       string   `json:"id" form:"id"`
	Sender   string   `json:"sender" form:"sender"`
	Receiver string   `json:"receiver" form:"receiver"`
	Message  string   `json:"message" form:"message"`
	FileUrls []string `json:"fileUrls" form:"fileUrls"`
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
