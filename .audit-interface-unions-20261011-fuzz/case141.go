package interfacefuzz
type Bound141 interface{
 ~[2]int | **struct{ A int }
}
type Alias141[T Bound141] = []T
func Scan141[T Bound141](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
