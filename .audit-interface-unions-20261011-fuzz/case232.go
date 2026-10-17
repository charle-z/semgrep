package interfacefuzz
import "time"
type Bound232 interface{
 time.Duration | ~*string | string
 M() int
}
func Scan232[T Bound232](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
