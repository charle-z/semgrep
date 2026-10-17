package interfacefuzz
type Bound209 interface{
 *int /* variant */ | ~*string /* variant */ | ~map[string]bool /* variant */ | ~func() int
}
func Scan209[T Bound209](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
