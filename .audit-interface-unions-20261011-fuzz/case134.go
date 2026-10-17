package interfacefuzz
type Bound134 interface{
 *int | ~chan int | ~[3]string | Pair[string,int] | ~func(string) bool
}
func Scan134[T Bound134](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
