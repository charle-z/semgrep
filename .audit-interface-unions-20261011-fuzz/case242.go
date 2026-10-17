package interfacefuzz
type Bound242 interface{
 Pair[int,string] /* variant */ | ~map[int]int /* variant */ | ~func(string) bool
}
func Scan242[T Bound242](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
