# Env

[![Build Status](https://travis-ci.org/jelmersnoeck/env.svg?branch=master)](https://travis-ci.org/jelmersnoeck/env)
[![GoDoc](https://godoc.org/github.com/jelmersnoeck/env?status.svg)](https://godoc.org/github.com/jelmersnoeck/env)

Env makes it easy to get your environment variables in the correct typed format.
It also introduces a fallback system in case the environmen variable is not set
or is not valid for the requested type.

## Supported methods

Have a look at the available methods in the [documentation](https://godoc.org/github.com/jelmersnoeck/env).

## Notifiers

Notifiers can be used to take action when something goes wrong. This could be
when an environment variabe isn't set or when it doesn't match what we expect
it to be, for example when parsing durations.

There are several implementations of notifiers, the most basic one being the
`nilNotifier` which does nothing with the notifications.

Another implementation is the `panicNotifier`, which will panic when
notifications happen, terminating the application.

## Usage

### Defaults

```go
import (
    "github.com/jelmersnoeck/env"
)

func main() {
    os.Setenv("MY_ENV", "5s") // do this on launching your application

    time.Sleep(env.Duration("MY_ENV", time.Second)) // this will sleep for 5 seconds
    time.Sleep(env.Duration("MY_OTHER_ENV", time.Second)) // this will sleep for a second

    os.Setenv("WRONG_ENV", "hello")
    time.Sleep(env.Duration("WRONG_ENV", time.Minute)) // this will sleep for a minute
}
```

### Panic notifier

```go
import (
    "github.com/jelmersnoeck/env"
    "github.com/jelmersnoeck/env/notifier/panic"
)

func main() {
    env.UseNotifier(panic.NewNotifier())

    os.Setenv("MY_ENV", "5s") // do this on launching your application

    time.Sleep(env.Duration("MY_ENV", time.Second)) // this will sleep for 5 seconds
    time.Sleep(env.Duration("MY_OTHER_ENV", time.Second)) // this will panic because the env is not set

    os.Setenv("WRONG_ENV", "hello")
    time.Sleep(env.Duration("WRONG_ENV", time.Minute)) // this will panic because the env is not a duration
}
```
