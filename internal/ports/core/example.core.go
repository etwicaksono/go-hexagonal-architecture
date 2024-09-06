package core

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

type ExampleCoreInterface interface {
	GetTextMessage() ([]*entity.MessageTextItem, error)
}
