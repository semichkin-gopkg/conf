# Example

```go
package main

import (
	"github.com/semichkin-gopkg/conf"
	"log"
)

type Conf struct {
	Foo int
	Bar bool
}

func WithFoo(foo int) conf.Updater[Conf] {
	return func(c *Conf) {
		c.Foo = foo
	}
}

func WithBar(bar bool) conf.Updater[Conf] {
	return func(c *Conf) {
		c.Bar = bar
	}
}

func main() {
	configs := conf.New[Conf]().
		Append(WithFoo(5)).
		Append(WithBar(true)).
		Prepend(WithBar(false)).
		Build()

	log.Println(configs.Foo) // 5
	log.Println(configs.Bar) // false
}
```