package audit
import "time"
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
(Box[int]) /* token */ | (Box[string]) /* token */ | interface{~string} /* token */ | time.Duration
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
