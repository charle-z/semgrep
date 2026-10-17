package interfacefuzz
type Bound227 interface{
 int | ~func(string) bool | ~[3]string
}
func Scan227[T Bound227](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
