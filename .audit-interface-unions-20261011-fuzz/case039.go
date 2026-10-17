package interfacefuzz
type Bound039 interface{
 ~func() int | ~[3]string | ~struct{ A int } | ~func(string) bool | ~[]byte
}
type Alias039[T Bound039] = []T
func Scan039[T Bound039](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
