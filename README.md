# Env

[![Build Status](https://travis-ci.org/jelmersnoeck/env.svg?branch=master)](https://travis-ci.org/jelmersnoeck/env)
[![GoDoc](https://godoc.org/github.com/jelmersnoeck/env?status.svg)](https://godoc.org/github.com/jelmersnoeck/env)

Env allows you to get environment variables and fallback to defaults if the
environment variable is not set or not valid.

## Notifiers

Notifiers can be used to take action when something goes wrong. This could be
when an environment variabe isn't set or when it doesn't match what we expect
it to be, for example when parsing durations.

There are several implementations of notifiers, the most basic one being the
`nilNotifier` which does nothing with the notifications.

Another implementation is the `panicNotifier`, which will panic when
notifications happen, terminating the application.

## Supported methods

- String
- Bool
- Duration
- Int64
- Int32
- Int
- Float64
- Float32

## Usage

```go
func main() {
    os.Setenv("MY_ENV", "5s") // do this on launching your application

    time.Sleep(env.Duration("MY_ENV", time.Second)) // this will sleep for 5 seconds
}
```
