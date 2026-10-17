package interfacefuzz
type Bound213 interface{
 ~[2]int | Pair[string,int] | Pair[int,string] | ~func(string) bool
}
type Alias213[T Bound213] = []T
func Scan213[T Bound213](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
