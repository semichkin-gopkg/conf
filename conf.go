package conf

type Updater[T any] func(configuration *T)

type Builder[T any] struct {
	updaters      []Updater[T]
}

func Build[T any](updaters ...Updater[T]) T {
	return New[T]().Append(updaters...).Build()
}

func New[T any]() *Builder[T] {
	return &Builder[T]{}
}

func (c *Builder[T]) Append(updaters ...Updater[T]) *Builder[T] {
	c.updaters = append(c.updaters, updaters...)
	return c
}

func (c *Builder[T]) Prepend(updaters ...Updater[T]) *Builder[T] {
	c.updaters = append(updaters, c.updaters...)
	return c
}

func (c *Builder[T]) Build() T {
	var configuration T

	for _, update := range c.updaters {
		if update != nil {
			update(&configuration)
		}
	}

	return configuration
}
