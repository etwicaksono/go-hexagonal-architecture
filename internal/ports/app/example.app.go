package app

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

type ExampleAppInterface interface {
	GetTextMessage() ([]entity.MessageTextItem, error)
}
