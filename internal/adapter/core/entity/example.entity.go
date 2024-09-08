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

// Multimedia Message
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
	MultimediaStorage_value = map[string]int32{
		"LOCAL": 0,
		"MINIO": 1,
	}
)

type SendMultimediaMessageRequest struct {
	Sender   string
	Receiver string
	Message  string
	Storage  MultimediaStorage
	Files    [][]byte
}

type MessageMultimediaItem struct {
	Id       string
	Sender   string
	Receiver string
	Message  string
	FileUrl  string
}
