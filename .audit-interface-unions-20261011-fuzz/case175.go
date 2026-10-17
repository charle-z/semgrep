package interfacefuzz
type Embedded175 interface{~[]rune | ~map[string]rune}
type Bound175 interface{
Embedded175
 Pair[string,int]
}
func Scan175[T Bound175](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
