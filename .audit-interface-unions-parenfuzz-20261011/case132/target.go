package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
interface{~string} | Box[map[string]int]
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
