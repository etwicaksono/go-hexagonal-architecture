package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileItem struct {
	Storage string `bson:"storage"`
	File    string `bson:"file"`
}

type MessageMultimediaItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string             `bson:"sender"`
	Receiver string             `bson:"receiver"`
	Message  string             `bson:"message"`
	Files    []FileItem         `bson:"files"`
}

func FromMessageMultimediaItemEntity(mmi entity.MessageMultimediaItem) MessageMultimediaItem {
	var files []FileItem
	for _, file := range mmi.Files {
		files = append(files, FileItem{
			Storage: file.Storage.ToString(),
			File:    file.File,
		})
	}
	messageItem := MessageMultimediaItem{
		Sender:   mmi.Sender,
		Receiver: mmi.Receiver,
		Message:  mmi.Message,
		Files:    files,
	}
	if mmi.Id != "" {
		messageItem.Id, _ = primitive.ObjectIDFromHex(mmi.Id)
	}
	return messageItem
}

func (mti MessageMultimediaItem) ToEntity() entity.MessageMultimediaItem {
	var files []entity.FileItem

	for _, file := range mti.Files {
		files = append(files, entity.FileItem{
			Storage: valueobject.MultimediaStorageFromString(file.Storage),
			File:    file.File,
		})
	}

	return entity.MessageMultimediaItem{
		Id:       mti.Id.Hex(),
		Sender:   mti.Sender,
		Receiver: mti.Receiver,
		Message:  mti.Message,
		Files:    files,
	}
}
