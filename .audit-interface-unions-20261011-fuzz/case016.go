package interfacefuzz
type Bound016 interface{
 ~struct{ B string } | ~[]byte
 M() int
}
func Scan016[T Bound016](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
