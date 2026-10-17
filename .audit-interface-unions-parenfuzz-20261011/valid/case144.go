package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type In interface{~[]rune | ~map[string]rune}
type C interface {
 In
~(int) | ~([3]int)
}
func Audit(xs []int) {auditSink(xs); for range xs {println("found")}}
