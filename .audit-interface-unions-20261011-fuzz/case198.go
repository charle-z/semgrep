package interfacefuzz
type Bound198 interface{
 ~[2]int /* variant */ | ~*string /* variant */ | ~complex64 /* variant */ | Pair[int,string]
}
type Alias198[T Bound198] = []T
func Scan198[T Bound198](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
