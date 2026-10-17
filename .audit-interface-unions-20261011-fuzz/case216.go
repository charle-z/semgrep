package interfacefuzz
type Bound216 interface{
 ~map[int]int | Pair[string,int]
 M() int
}
type Alias216[T Bound216] = []T
func Scan216[T Bound216](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
