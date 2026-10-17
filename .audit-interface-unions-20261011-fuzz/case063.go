package interfacefuzz
type Embedded063 interface{~[]rune | ~map[string]rune}
type Bound063 interface{
Embedded063
 Pair[string,int] | ~map[int]int | string | ~chan int
}
type Alias063[T Bound063] = []T
func Scan063[T Bound063](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
