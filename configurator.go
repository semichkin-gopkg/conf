package configurator

type Updater[T any] func(configuration *T)

type Configurator[T any] struct {
	updaters []Updater[T]
}

func New[T any]() *Configurator[T] {
	return &Configurator[T]{}
}

func (c *Configurator[T]) Append(updaters ...Updater[T]) *Configurator[T] {
	c.updaters = append(c.updaters, updaters...)
	return c
}

func (c *Configurator[T]) Prepend(updaters ...Updater[T]) *Configurator[T] {
	c.updaters = append(updaters, c.updaters...)
	return c
}

func Some(func (int)) {

}

var a = Some(func(i int) {
	
})

func (c *Configurator[T]) Apply() T {
	var configuration T

	for _, update := range c.updaters {
		update(&configuration)
	}

	return configuration
}
