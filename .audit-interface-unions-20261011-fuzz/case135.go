package interfacefuzz
type Bound135 interface{
 Pair[string,int]
}
type Alias135[T Bound135] = []T
func Scan135[T Bound135](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
