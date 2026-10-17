package interfacefuzz
type Bound155 interface{
 string
}
func Scan155[T Bound155](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
