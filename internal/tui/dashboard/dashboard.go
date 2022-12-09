package dashboard

type status int

const (
	todo status = iota
	inProg
	done
	ideas
)

type Model struct {
}
