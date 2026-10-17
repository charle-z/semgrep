package interfacefuzz
type Bound149 interface{
 Pair[Box[int],Box[string]] | ~chan int | ~[]string | *int | ~map[int]int
}
func Scan149[T Bound149](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
