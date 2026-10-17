package interfacefuzz
type Embedded000 interface{~[]rune | ~map[string]rune}
type Bound000 interface{
Embedded000
 *int
 M() int
}
type Alias000[T Bound000] = []T
func Scan000[T Bound000](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
