package notifier

import (
	"fmt"

	"github.com/jelmersnoeck/env"
)

type panicNotifier struct{}

func NewPanic() env.Notifier {
	return &panicNotifier{}
}

func (n *panicNotifier) Notify(msg string, i ...interface{}) {
	panic(fmt.Sprintf(msg, i...))
}
