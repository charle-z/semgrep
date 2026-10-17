package interfacefuzz
type Bound165 interface{
 Box[int]
}
type Alias165[T Bound165] = []T
func Scan165[T Bound165](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
