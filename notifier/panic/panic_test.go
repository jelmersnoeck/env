package panic_test

import (
	"testing"

	"github.com/jelmersnoeck/env/notifier/panic"
	"github.com/stretchr/testify/assert"
)

func TestPanic_Notify(t *testing.T) {
	not := panic.NewNotifier()

	assert.Panics(t, func() { not.Notify("mymessage") })
}
