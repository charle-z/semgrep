package interfacefuzz
import "time"
type Bound246 interface{
 time.Duration | ~func() int
}
type Alias246[T Bound246] = []T
func Scan246[T Bound246](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
