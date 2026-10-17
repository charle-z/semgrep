package interfacefuzz
type Bound184 interface{
 **struct{ A int } | ~map[string]bool | int | ~*string | *int
 M() int
}
func Scan184[T Bound184](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
