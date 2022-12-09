package ctx

type Ctx struct {
	Content [2]int
	Screen  [2]int
	Loading bool
}

func New() Ctx {
	ctx := Ctx{
		Screen:  [2]int{0, 0},
		Content: [2]int{0, 0},
		Loading: false,
	}

	return ctx
}
