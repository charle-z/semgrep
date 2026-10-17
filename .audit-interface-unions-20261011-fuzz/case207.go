package interfacefuzz
type Bound207 interface{
 ~[]string | ~struct{ B string } | ~struct{ A int }
}
type Alias207[T Bound207] = []T
func Scan207[T Bound207](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
