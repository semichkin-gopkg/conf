# Example

```go
package main

import (
	"github.com/semichkin-gopkg/configurator"
	"log"
)

type Configuration struct {
	Foo int
	Bar bool
}

func WithFoo(foo int) configurator.Updater[Configuration] {
	return func(c *Configuration) {
		c.Foo = foo
	}
}

func WithBar(bar bool) configurator.Updater[Configuration] {
	return func(c *Configuration) {
		c.Bar = bar
	}
}

func main() {
	configs := configurator.New[Configuration]().
		Fix(WithFoo(2)).
		Append(WithFoo(5)).
		Append(WithBar(true)).
		Prepend(WithBar(false)).
		Apply()

	log.Println(configs.Foo) // 2
	log.Println(configs.Bar) // true
}
```