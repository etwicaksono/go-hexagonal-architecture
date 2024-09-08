package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageTextItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string             `bson:"sender"`
	Receiver string             `bson:"receiver"`
	Message  string             `bson:"message"`
}

func (mti MessageTextItem) ToEntity() entity.MessageTextItem {
	return entity.MessageTextItem{
		Id:       mti.Id.Hex(),
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
}

func FromMessageTextItemEntity(mti entity.MessageTextItem) MessageTextItem {
	messageItem := MessageTextItem{
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
	}
	if mti.Id != "" {
		messageItem.Id, _ = primitive.ObjectIDFromHex(mti.Id)
	}
	return messageItem
}

type MessageMultimediaItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string             `bson:"sender"`
	Receiver string             `bson:"receiver"`
	Message  string             `bson:"message"`
	FileUrl  string             `bson:"fileUrl"`
}

func FromMessageMultimediaItemEntity(mmi entity.MessageMultimediaItem) MessageMultimediaItem {
	messageItem := MessageMultimediaItem{
		Sender:   mmi.Sender,
		Receiver: mmi.Receiver,
		Message:  mmi.Message,
		FileUrl:  mmi.FileUrl,
	}
	if mmi.Id != "" {
		messageItem.Id, _ = primitive.ObjectIDFromHex(mmi.Id)
	}
	return messageItem
}
