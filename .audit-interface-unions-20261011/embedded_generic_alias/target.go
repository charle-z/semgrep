package audit
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type I[T any] = Box[T]
type C interface{ I[int] | I[string] }
