package interfacefuzz
type Embedded021 interface{~[]rune | ~map[string]rune}
type Bound021 interface{
Embedded021
 ~[2]int | Pair[Box[int],Box[string]]
}
type Alias021[T Bound021] = []T
func Scan021[T Bound021](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
