package interfacefuzz
type Bound061 interface{
 ~func() int | ~struct{ A int }
}
func Scan061[T Bound061](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
