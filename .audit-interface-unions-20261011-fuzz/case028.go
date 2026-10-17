package interfacefuzz
type Embedded028 interface{~[]rune | ~map[string]rune}
type Bound028 interface{
Embedded028
 Pair[Box[int],Box[string]] | ~func(string) bool | ~func() int | ~[]string
 M() int
}
func Scan028[T Bound028](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
