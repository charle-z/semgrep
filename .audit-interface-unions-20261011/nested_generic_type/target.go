package audit
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type C interface{ Pair[Box[int],Box[string]] | Pair[Box[string],Box[int]] }
