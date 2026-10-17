package interfacefuzz
type Bound183 interface{
 ~map[string]bool | ~func(string) bool | Pair[int,string] | ~chan string
}
type Alias183[T Bound183] = []T
func Scan183[T Bound183](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
