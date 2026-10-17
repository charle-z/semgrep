package interfacefuzz
type Embedded231 interface{~[]rune | ~map[string]rune}
type Bound231 interface{
Embedded231
 ~struct{ B string } /* variant */ | **struct{ A int }
}
type Alias231[T Bound231] = []T
func Scan231[T Bound231](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
