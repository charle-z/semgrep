package interfacefuzz
type Bound074 interface{
 string | int | Box[string] | ~chan string | ~func() int
}
func Scan074[T Bound074](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
