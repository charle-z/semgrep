package interfacefuzz
type Bound075 interface{
 ~func(string) bool
}
type Alias075[T Bound075] = []T
func Scan075[T Bound075](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
