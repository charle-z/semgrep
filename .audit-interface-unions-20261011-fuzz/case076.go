package interfacefuzz
type Bound076 interface{
 *int | ~*string
 M() int
}
func Scan076[T Bound076](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
