package interfacefuzz
type Bound012 interface{
 ~chan string | Box[string] | string
 M() int
}
type Alias012[T Bound012] = []T
func Scan012[T Bound012](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
