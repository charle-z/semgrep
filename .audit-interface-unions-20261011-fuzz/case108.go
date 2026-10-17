package interfacefuzz
type Bound108 interface{
 ~[]byte | ~func(string) bool | ~map[string]bool | ~*string
 M() int
}
type Alias108[T Bound108] = []T
func Scan108[T Bound108](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
