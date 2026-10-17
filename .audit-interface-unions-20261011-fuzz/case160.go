package interfacefuzz
type Bound160 interface{
 ~*string
 M() int
}
func Scan160[T Bound160](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
