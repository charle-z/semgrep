package interfacefuzz
type Bound017 interface[
 ~struct{ A int } | ~[]byte | Box[int]
}
func Scan017[T Bound017](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
