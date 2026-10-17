package interfacefuzz
type Bound055 interface{
 **struct{ A int }
}
func Scan055[T Bound055](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
