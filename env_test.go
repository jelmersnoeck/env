package env_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jelmersnoeck/env"
	"github.com/stretchr/testify/assert"
)

var ev = "ENV_TEST_VALUE"

func TestString(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, "default", env.String(ev, "default"))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "value", func() {
		assert.Equal(t, "value", env.String(ev, "default"))
	})
	assert.Len(t, not.notifications, 0)
}

func TestBool(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.True(t, env.Bool(ev, true))
		assert.False(t, env.Bool(ev, false))
	})
	assert.Len(t, not.notifications, 2)

	not = envTest(ev, "true", func() {
		assert.True(t, env.Bool(ev, true))
		assert.True(t, env.Bool(ev, false))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "1", func() {
		assert.True(t, env.Bool(ev, true))
		assert.True(t, env.Bool(ev, false))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "non-true", func() {
		assert.True(t, env.Bool(ev, true))
		assert.False(t, env.Bool(ev, false))
	})
	assert.Len(t, not.notifications, 2)

	not = envTest(ev, "false", func() {
		assert.False(t, env.Bool(ev, true))
		assert.False(t, env.Bool(ev, false))
	})
	assert.Len(t, not.notifications, 0)
}

func TestDuration(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, time.Minute, env.Duration(ev, time.Minute))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "5m", func() {
		assert.Equal(t, 5*time.Minute, env.Duration(ev, time.Minute))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid duration", func() {
		assert.Equal(t, time.Minute, env.Duration(ev, time.Minute))
	})
	assert.Len(t, not.notifications, 1)
}

func TestInt64(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, int64(5), env.Int64(ev, int64(5)))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "3", func() {
		assert.Equal(t, int64(3), env.Int64(ev, int64(5)))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid", func() {
		assert.Equal(t, int64(5), env.Int64(ev, int64(5)))
	})
	assert.Len(t, not.notifications, 1)
}

func TestInt32(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, int32(5), env.Int32(ev, int32(5)))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "3", func() {
		assert.Equal(t, int32(3), env.Int32(ev, int32(5)))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid", func() {
		assert.Equal(t, int32(5), env.Int32(ev, int32(5)))
	})
	assert.Len(t, not.notifications, 1)
}

func TestInt(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, int(5), env.Int(ev, int(5)))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "3", func() {
		assert.Equal(t, int(3), env.Int(ev, int(5)))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid", func() {
		assert.Equal(t, int(5), env.Int(ev, int(5)))
	})
	assert.Len(t, not.notifications, 1)
}

func envTest(ev, val string, f func()) *mockNotifier {
	not := &mockNotifier{}
	env.UseNotifier(not)

	os.Setenv(ev, val)
	defer os.Setenv(ev, "")

	f()

	return not
}

func TestFloat64(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, float64(5), env.Float64(ev, float64(5)))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "3", func() {
		assert.Equal(t, float64(3), env.Float64(ev, float64(5)))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid", func() {
		assert.Equal(t, float64(5), env.Float64(ev, float64(5)))
	})
	assert.Len(t, not.notifications, 1)
}

func TestFloat32(t *testing.T) {
	not := envTest(ev, "", func() {
		assert.Equal(t, float32(5), env.Float32(ev, float32(5)))
	})
	assert.Len(t, not.notifications, 1)

	not = envTest(ev, "3", func() {
		assert.Equal(t, float32(3), env.Float32(ev, float32(5)))
	})
	assert.Len(t, not.notifications, 0)

	not = envTest(ev, "invalid", func() {
		assert.Equal(t, float32(5), env.Float32(ev, float32(5)))
	})
	assert.Len(t, not.notifications, 1)
}

type mockNotifier struct {
	notifications []string
}

func (n *mockNotifier) Notify(msg string, i ...interface{}) {
	n.notifications = append(n.notifications, fmt.Sprintf(msg, i...))
}
