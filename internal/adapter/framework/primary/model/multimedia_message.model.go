package model

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

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

type FileItem struct {
	Storage string `json:"storage"`
	File    string `json:"file"`
}

type MessageMultimediaItem struct {
	Id       string     `json:"id"`
	Sender   string     `json:"sender"`
	Receiver string     `json:"receiver"`
	Message  string     `json:"message"`
	Files    []FileItem `json:"files"`
}

func FromMessageMultimediaItemEntity(s entity.MessageMultimediaItem) MessageMultimediaItem {
	var files []FileItem

	for _, file := range s.Files {
		files = append(files, FileItem{
			Storage: file.Storage,
			File:    file.File,
		})
	}

	return MessageMultimediaItem{
		Id:       s.Id,
		Sender:   s.Sender,
		Receiver: s.Receiver,
		Message:  s.Message,
		Files:    files,
	}
}
