package interfacefuzz
type Bound097 interface{
 Box[string] | ~[]byte | int
}
func Scan097[T Bound097](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
