package interfacefuzz
type Bound120 interface{
 ~struct{ B string }
 M() int
}
type Alias120[T Bound120] = []T
func Scan120[T Bound120](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
