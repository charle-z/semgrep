package interfacefuzz
type Bound240 interface{
 ~func() int
 M() int
}
type Alias240[T Bound240] = []T
func Scan240[T Bound240](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
