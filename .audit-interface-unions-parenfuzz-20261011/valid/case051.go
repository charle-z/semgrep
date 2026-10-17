package audit
import "time"
type Box[T any] struct{V T}
type Pair[A,B any] struct{A A;B B}
func auditSink(v any) {}
type C interface {
interface{~int} | ~([3]int) | ~func(...int) | interface{~string} | time.Duration
}
func Audit(xs []int) {auditSink(xs); for j:=0;j<len(xs);j++ {println(xs[j])}}
