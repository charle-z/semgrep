package interfacefuzz
type Embedded217 interface{~[]rune | ~map[string]rune}
type Bound217 interface{
Embedded217
 **struct{ A int } | Pair[int,string] | string
}
func Scan217[T Bound217](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
