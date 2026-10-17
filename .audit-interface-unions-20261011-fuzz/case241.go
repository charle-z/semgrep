package interfacefuzz
type Bound241 interface{
 ~[]string | ~struct{ A int }
}
func Scan241[T Bound241](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
