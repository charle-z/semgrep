package interfacefuzz
type Embedded245 interface{~[]rune | ~map[string]rune}
type Bound245 interface{
Embedded245
 ~map[int]int
}
func Scan245[T Bound245](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
