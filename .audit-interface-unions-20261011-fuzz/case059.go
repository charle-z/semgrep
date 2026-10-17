package interfacefuzz
type Bound059 interface{
 *int | ~map[string]bool | ~[3]string | Pair[string,int] | ~struct{ B string }
}
func Scan059[T Bound059](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
