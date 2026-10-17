package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
~(func(int) bool) | ~[]byte | ~(struct{X int}) | ~map[string]bool | (func() int)
}
func Audit(xs []int) {auditSink(xs); for j:=0;j<len(xs);j++ {println(xs[j])}}
