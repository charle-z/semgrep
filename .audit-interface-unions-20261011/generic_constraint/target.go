package audit
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type C[T interface{ ~[]byte | ~map[string]bool }] struct{V T}
