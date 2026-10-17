package interfacefuzz
type Bound003 interface{
 |  ~*string | ~func() int | ~map[int]int | ~struct{ A int }
}
type Alias003[T Bound003] = []T
func Scan003[T Bound003](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
