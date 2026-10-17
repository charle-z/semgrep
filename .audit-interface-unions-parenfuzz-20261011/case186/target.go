package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
~(map[string]int) | (func() int) | ~map[string]bool | interface{}
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
