package audit
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type I interface{ ~int | ~string }
type C interface{ ~<-chan string | ~[]byte }
