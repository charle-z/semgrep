package interfacefuzz
type Bound060 interface{
 Box[int]
 M() int
}
type Alias060[T Bound060] = []T
func Scan060[T Bound060](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
