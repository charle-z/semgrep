package audit
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type In interface{~[]rune | ~map[string]rune}
type C interface {
 In
interface{} | ~chan<- string | ~func(...int)
}
func Audit(xs []int) {auditSink(xs); for j:=0;j<len(xs);j++ {println(xs[j])}}
