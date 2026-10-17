package interfacefuzz
type Bound066 interface{
 Box[int] /* variant */ | ~struct{ A int }
}
type Alias066[T Bound066] = []T
func Scan066[T Bound066](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
