package configurator

type Updater[T any] func(configuration *T)

type Configurator[T any] struct {
	updaters      []Updater[T]
	fixedUpdaters []Updater[T]
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

func (c *Configurator[T]) Fix(updaters ...Updater[T]) *Configurator[T] {
	c.fixedUpdaters = append(c.fixedUpdaters, updaters...)
	return c
}

func (c *Configurator[T]) Apply() T {
	var configuration T

	updaters := make([]Updater[T], 0, len(c.updaters)+len(c.fixedUpdaters))
	updaters = append(updaters, c.updaters...)
	updaters = append(updaters, c.fixedUpdaters...)

	for _, update := range updaters {
		update(&configuration)
	}

	return configuration
}
