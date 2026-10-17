package interfacefuzz
type Embedded182 interface{~[]rune | ~map[string]rune}
type Bound182 interface{
Embedded182
 int | Box[int] | ~[2]int
}
func Scan182[T Bound182](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
