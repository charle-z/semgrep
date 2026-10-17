package interfacefuzz
type Bound082 interface{
 ~complex64 | ~[]byte | Pair[string,int]
}
func Scan082[T Bound082](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
