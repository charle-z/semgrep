package interfacefuzz
type Bound253 interface{
 ~complex64 /* variant */ | ~chan string /* variant */ | Pair[Box[int],Box[string]] /* variant */ | Box[string]
}
func Scan253[T Bound253](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
