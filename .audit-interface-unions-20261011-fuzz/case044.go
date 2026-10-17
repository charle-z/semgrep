package interfacefuzz
type Bound044 interface{
 ~map[int]int /* variant */ | int /* variant */ | Box[int] /* variant */ | ~struct{ A int } /* variant */ | ~func(string) bool
 M() int
}
func Scan044[T Bound044](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
