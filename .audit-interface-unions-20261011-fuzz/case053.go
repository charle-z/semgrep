package interfacefuzz
type Bound053 interface{
 *int | Pair[int,string] | ~[3]string | ~func() int
}
func Scan053[T Bound053](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
