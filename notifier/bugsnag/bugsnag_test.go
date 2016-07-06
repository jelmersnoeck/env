package bugsnag_test

import (
	"testing"

	"github.com/jelmersnoeck/env/notifier/bugsnag"
	"github.com/stretchr/testify/assert"
)

func TestBugsnag_Notify(t *testing.T) {
	n := &notifier{[]error{}}
	not := bugsnag.NewNotifier(n)
	not.Notify("my message %s", "foo")

	assert.Len(t, n.nots, 1)
	assert.Equal(t, "my message foo", n.nots[0].Error())
}

type notifier struct {
	nots []error
}

func (n *notifier) Notify(e error, r ...interface{}) error {
	n.nots = append(n.nots, e)
	return nil
}
