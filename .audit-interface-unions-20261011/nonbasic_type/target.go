package audit
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type C interface { ~[]byte | ~map[string]bool }
func F[T C](v T) T {return v}
