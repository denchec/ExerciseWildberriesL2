package pkg

type Navigator struct {
	Strategy
}

func (nav *Navigator) Setstrategy(str Strategy) {
	nav.Strategy = str
}
