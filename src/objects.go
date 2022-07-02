package src

// objects
const (
	VOID_O = iota
	NUM_O
	BOOL_O
	NAME_O
	LAMBDA_O
	LIST_O
)

type Object struct {
	Type    int
	Content interface{}
	X       int
	Y       int
}
