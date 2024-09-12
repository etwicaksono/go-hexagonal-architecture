package entity

import (
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

type SendTextMessageRequest struct {
	Sender   string
	Receiver string
	Message  string
}

type MultimediaStorage int32

const (
	MultimediaStorage_LOCAL MultimediaStorage = 0
	MultimediaStorage_MINIO MultimediaStorage = 1
)

// Enum value maps for MultimediaStorage.
var (
	MultimediaStorage_name = map[int32]string{
		0: "LOCAL",
		1: "MINIO",
	}
	MultimediaStorage_self = map[string]MultimediaStorage{
		"LOCAL": MultimediaStorage_LOCAL,
		"MINIO": MultimediaStorage_MINIO,
	}
)

type MultimediaFile struct {
	Filename    string
	ContentType string
	Data        []byte
}

type SendMultimediaMessageRequest struct {
	Sender   string
	Receiver string
	Message  string
	Storage  MultimediaStorage
	Files    []MultimediaFile
}

type FileItem struct {
	Storage string
	File    string
}

type MessageMultimediaItem struct {
	Id       string
	Sender   string
	Receiver string
	Message  string
	Files    []FileItem
}

func (mmi MessageMultimediaItem) ToProto() *example.MessageMultimediaItem {
	var fileUrl []string

	for _, file := range mmi.Files {
		fileUrl = append(fileUrl, file.File)
	}

	return &example.MessageMultimediaItem{
		Id:       mmi.Id,
		Sender:   mmi.Sender,
		Receiver: mmi.Receiver,
		Message:  mmi.Message,
		FileUrls: fileUrl,
	}
}
