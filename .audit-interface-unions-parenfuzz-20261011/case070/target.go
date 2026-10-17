package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
~(chan int) /* token */ | *(Box[string]) /* token */ | (*struct{X int}) /* token */ | *int
 M() int
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
