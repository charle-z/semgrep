package interfacefuzz
type Embedded007 interface{~[]rune | ~map[string]rune}
type Bound007 interface{
Embedded007
 ~chan int | string | ~complex64
}
func Scan007[T Bound007](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
