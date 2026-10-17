package interfacefuzz
type Embedded133 interface{~[]rune | ~map[string]rune}
type Bound133 interface{
Embedded133
 *int | ~*string | ~map[string]bool | ~func() int
}
func Scan133[T Bound133](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
