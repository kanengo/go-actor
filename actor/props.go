package actor

type Props struct {
	producer Producer
}

func PropsFromActor(actor Actor, opts ...PropsOption) *Props {
	props := &Props{producer: func() Actor {
		return actor
	}}

	for _, opt := range opts {
		opt(props)
	}

	return props
}
