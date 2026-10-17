package interfacefuzz
type Bound036 interface{
 Pair[Box[int],Box[string]] | ~*string
 M() int
}
type Alias036[T Bound036] = []T
func Scan036[T Bound036](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
