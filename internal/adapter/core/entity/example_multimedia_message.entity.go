package entity

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"github.com/etwicaksono/public-proto/gen/example"
)

type SendMultimediaMessageRequest struct {
	Sender   string
	Receiver string
	Message  string
	Storage  valueobject.MultimediaStorage
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
