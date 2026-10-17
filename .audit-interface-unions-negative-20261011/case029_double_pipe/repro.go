package interfacefuzz
type Bound029 interface{
 ~struct{ B string } || string | ~[]byte | ~map[int]int | ~chan string
}
func Scan029[T Bound029](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
