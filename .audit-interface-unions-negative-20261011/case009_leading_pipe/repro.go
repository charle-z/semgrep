package interfacefuzz
type Bound009 interface{
 |  ~[2]int | Box[int] | int | ~func() int | Pair[int,string]
}
type Alias009[T Bound009] = []T
func Scan009[T Bound009](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
