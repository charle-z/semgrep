package interfacefuzz
type Embedded049 interface{~[]rune | ~map[string]rune}
type Bound049 interface{
Embedded049
 int | ~[3]string | ~complex64 | Pair[int,string] | *int
}
func Scan049[T Bound049](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
