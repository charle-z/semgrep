package interfacefuzz
type Bound249 interface{
 Pair[Box[int],Box[string]] | ~chan int | **struct{ A int } | ~[]string | string
}
type Alias249[T Bound249] = []T
func Scan249[T Bound249](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
