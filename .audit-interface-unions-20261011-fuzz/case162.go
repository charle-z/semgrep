package interfacefuzz
type Bound162 interface{
 ~[]string | ~func(string) bool | ~[2]int
}
type Alias162[T Bound162] = []T
func Scan162[T Bound162](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
