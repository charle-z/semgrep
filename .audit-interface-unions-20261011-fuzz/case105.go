package interfacefuzz
type Embedded105 interface{~[]rune | ~map[string]rune}
type Bound105 interface{
Embedded105
 ~[]string
}
type Alias105[T Bound105] = []T
func Scan105[T Bound105](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
