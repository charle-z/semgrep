package interfacefuzz
type Bound048 interface{
 Pair[Box[int],Box[string]] | ~*string | ~[]string | ~[3]string
 M() int
}
type Alias048[T Bound048] = []T
func Scan048[T Bound048](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
