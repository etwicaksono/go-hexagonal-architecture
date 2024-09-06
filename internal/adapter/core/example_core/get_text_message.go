package example_core

import (
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"strconv"
	"time"
)

func (e exampleCore) GetTextMessage() ([]*entity.MessageTextItem, error) {
	var messages []*entity.MessageTextItem

	for i := 0; i < 10; i++ {
		messages = append(messages, &entity.MessageTextItem{
			Id:       fmt.Sprintf("id:%s:%s", strconv.Itoa(i), strconv.Itoa(int(time.Now().Unix()))),
			Sender:   fmt.Sprintf("sender:%s:%s", strconv.Itoa(i), strconv.Itoa(int(time.Now().Unix()))),
			Receiver: fmt.Sprintf("receiver:%s:%s", strconv.Itoa(i), strconv.Itoa(int(time.Now().Unix()))),
			Message:  fmt.Sprintf("message:%s:%s", strconv.Itoa(i), strconv.Itoa(int(time.Now().Unix()))),
		})
	}

	return messages, nil
}
