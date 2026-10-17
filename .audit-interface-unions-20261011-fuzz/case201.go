package interfacefuzz
type Bound201 interface{
 string | ~func(string) bool
}
type Alias201[T Bound201] = []T
func Scan201[T Bound201](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
