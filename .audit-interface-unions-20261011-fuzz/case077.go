package interfacefuzz
type Embedded077 interface{~[]rune | ~map[string]rune}
type Bound077 interface{
Embedded077
 ~map[string]bool /* variant */ | ~*string /* variant */ | *int
}
func Scan077[T Bound077](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
