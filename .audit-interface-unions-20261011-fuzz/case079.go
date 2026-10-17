package interfacefuzz
type Bound079 interface{
 ~[2]int | Box[string] | ~struct{ B string } | ~*string | int
}
func Scan079[T Bound079](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
