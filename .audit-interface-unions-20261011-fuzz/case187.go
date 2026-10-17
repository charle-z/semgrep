package interfacefuzz
type Bound187 interface{
 ~struct{ A int } /* variant */ | ~[]byte /* variant */ | Box[int]
}
func Scan187[T Bound187](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
