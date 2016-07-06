package env

// Notifier is an interface used when getter methods encounter an error and
// should the application. The getters will always notify the notifier upon an
// error and return a fitting value.
type Notifier interface {
	Notify(msg string, i ...interface{})
}

var (
	notifier Notifier = &nilNotifier{}
)

// UseNotifier sets the notifier being used for the package to the given
// notifier. The default notifier is a nil notifier which does nothing with the
// notifications.
func UseNotifier(n Notifier) {
	notifier = n
}

type nilNotifier struct{}

func (n *nilNotifier) Notify(msg string, i ...interface{}) {
}
