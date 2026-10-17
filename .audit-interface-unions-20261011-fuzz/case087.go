package interfacefuzz
type Bound087 interface{
 string | *int | Pair[string,int]
}
type Alias087[T Bound087] = []T
func Scan087[T Bound087](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
