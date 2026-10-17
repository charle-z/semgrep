package interfacefuzz
type Bound032 interface{
 ~struct{ B string } | **struct{ A int } | ~chan string
 M() int
}
func Scan032[T Bound032](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
