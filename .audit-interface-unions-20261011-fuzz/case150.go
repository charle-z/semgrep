package interfacefuzz
type Bound150 interface{
 ~[3]string
}
type Alias150[T Bound150] = []T
func Scan150[T Bound150](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
