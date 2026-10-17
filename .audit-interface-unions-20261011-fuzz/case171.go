package interfacefuzz
type Bound171 interface{
 *int | ~*string
}
type Alias171[T Bound171] = []T
func Scan171[T Bound171](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
