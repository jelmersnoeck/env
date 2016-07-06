package panic

import (
	"fmt"

	"github.com/jelmersnoeck/env"
)

type panicNotifier struct{}

func NewNotifier() env.Notifier {
	return &panicNotifier{}
}

func (n *panicNotifier) Notify(msg string, i ...interface{}) {
	panic(fmt.Sprintf(msg, i...))
}
