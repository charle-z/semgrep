package interfacefuzz
type Bound180 interface{
 ~[]byte
 M() int
}
type Alias180[T Bound180] = []T
func Scan180[T Bound180](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
