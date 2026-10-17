package interfacefuzz
type Embedded161 interface{~[]rune | ~map[string]rune}
type Bound161 interface{
Embedded161
 ~[]byte | ~[]string
}
func Scan161[T Bound161](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
