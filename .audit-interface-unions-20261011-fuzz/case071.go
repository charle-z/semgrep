package interfacefuzz
type Bound071 interface{
 ~struct{ B string } | ~[3]string
}
func Scan071[T Bound071](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
