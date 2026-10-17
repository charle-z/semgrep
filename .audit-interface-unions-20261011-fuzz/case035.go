package interfacefuzz
type Embedded035 interface{~[]rune | ~map[string]rune}
type Bound035 interface{
Embedded035
 Pair[string,int]
}
func Scan035[T Bound035](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
