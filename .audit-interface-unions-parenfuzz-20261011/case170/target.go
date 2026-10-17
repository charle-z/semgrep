package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
~(struct{X int}) | *(Box[int]) | ~(chan int) | (Pair[int,string])
 M() int
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
