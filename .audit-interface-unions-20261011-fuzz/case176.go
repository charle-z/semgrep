package interfacefuzz
type Bound176 interface{
 ~[2]int /* variant */ | ~chan string
 M() int
}
func Scan176[T Bound176](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
