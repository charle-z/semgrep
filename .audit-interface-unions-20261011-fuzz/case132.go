package interfacefuzz
type Bound132 interface{
 ~map[string]bool /* variant */ | Box[string] /* variant */ | ~complex64
 M() int
}
type Alias132[T Bound132] = []T
func Scan132[T Bound132](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
