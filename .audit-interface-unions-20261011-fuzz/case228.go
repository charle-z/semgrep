package interfacefuzz
type Bound228 interface{
 *int | ~*string | ~map[string]bool | ~func() int
 M() int
}
type Alias228[T Bound228] = []T
func Scan228[T Bound228](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
