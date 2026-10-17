package interfacefuzz
type Bound069 interface{
 ~func() int | ~[]string | int | ~struct{ B string } | Pair[int,string]
}
type Alias069[T Bound069] = []T
func Scan069[T Bound069](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
