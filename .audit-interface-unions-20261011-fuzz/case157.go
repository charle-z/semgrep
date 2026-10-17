package interfacefuzz
type Bound157 interface{
 Box[string] | string | Box[int]
}
func Scan157[T Bound157](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
