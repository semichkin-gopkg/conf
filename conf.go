package conf

type Updater[T any] func(configuration *T)

type Builder[T any] struct {
	updaters      []Updater[T]
	fixedUpdaters []Updater[T]
}

func NewBuilder[T any]() *Builder[T] {
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

func (c *Builder[T]) Fix(updaters ...Updater[T]) *Builder[T] {
	c.fixedUpdaters = append(c.fixedUpdaters, updaters...)
	return c
}

func (c *Builder[T]) Build() T {
	var configuration T

	updaters := make([]Updater[T], 0, len(c.updaters)+len(c.fixedUpdaters))
	updaters = append(updaters, c.updaters...)
	updaters = append(updaters, c.fixedUpdaters...)

	for _, update := range updaters {
		update(&configuration)
	}

	return configuration
}
