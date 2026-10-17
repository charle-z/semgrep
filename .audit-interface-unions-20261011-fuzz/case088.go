package interfacefuzz
type Bound088 interface{
 ~[2]int /* variant */ | string /* variant */ | ~complex64 /* variant */ | ~map[int]int
 M() int
}
func Scan088[T Bound088](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
