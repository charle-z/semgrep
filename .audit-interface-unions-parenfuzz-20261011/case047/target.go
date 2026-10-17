package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
(Pair[int,string]) | (*struct{X int}) | ~(func(int) bool) | interface{~int} | ~(struct{X int})
}
func Audit(xs []int) {auditSink(xs); for j:=0;j<len(xs);j++ {println(xs[j])}}
