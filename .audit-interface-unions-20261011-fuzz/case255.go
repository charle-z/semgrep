package interfacefuzz
type Bound255 interface{
 ~struct{ A int }
}
type Alias255[T Bound255] = []T
func Scan255[T Bound255](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
