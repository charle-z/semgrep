package interfacefuzz
type Bound019 interface{
 |  *int | ~*string | ~map[string]bool | ~func() int
}
func Scan019[T Bound019](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
