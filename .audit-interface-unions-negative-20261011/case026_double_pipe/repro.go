package interfacefuzz
type Bound026 interface{
 *int || ~[]byte
}
func Scan026[T Bound026](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
