package bugsnag

import (
	"fmt"

	"github.com/jelmersnoeck/env"
)

type (
	bugsnagClient interface {
		Notify(error, ...interface{}) error
	}

	bugsnagNotifier struct {
		notifier bugsnagClient
	}
)

func NewNotifier(not bugsnagClient) env.Notifier {
	return &bugsnagNotifier{not}
}

func (n *bugsnagNotifier) Notify(msg string, i ...interface{}) {
	n.notifier.Notify(fmt.Errorf(msg, i...))
}
