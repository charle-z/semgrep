package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
~([3]int) /* token */ | ([]string) /* token */ | *(Box[string]) /* token */ | (Pair[int,string])
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
