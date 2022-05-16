package pkg

type Strategy interface {
	Route(startPoint int, endPoint int)
}
