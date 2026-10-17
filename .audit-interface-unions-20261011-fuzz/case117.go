package interfacefuzz
type Bound117 interface{
 ~struct{ A int } | ~[]byte | Box[int]
}
type Alias117[T Bound117] = []T
func Scan117[T Bound117](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
