package interfacefuzz
type Bound121 interface{
 Box[string] /* variant */ | **struct{ A int }
}
func Scan121[T Bound121](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
