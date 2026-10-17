package interfacefuzz
type Embedded147 interface{~[]rune | ~map[string]rune}
type Bound147 interface{
Embedded147
 Pair[Box[int],Box[string]] | Pair[string,int] | ~func() int
}
type Alias147[T Bound147] = []T
func Scan147[T Bound147](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
