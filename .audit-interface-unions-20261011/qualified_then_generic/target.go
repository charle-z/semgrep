package audit
import "time"
type Box[T any] struct{ V T }
type Pair[A, B any] struct{ A A; B B }
type C interface{ time.Duration | Box[int] }
