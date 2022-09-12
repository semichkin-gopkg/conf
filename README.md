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
		Append(WithFoo(5)).
		Append(WithBar(true)).
		Apply()

	log.Println(configs.Foo) // 5
	log.Println(configs.Bar) // true
}


```