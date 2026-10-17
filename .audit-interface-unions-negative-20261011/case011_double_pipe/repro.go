package interfacefuzz
type Bound011 interface{
 Pair[string,int] /* variant */ || Pair[Box[int],Box[string]]
}
func Scan011[T Bound011](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
