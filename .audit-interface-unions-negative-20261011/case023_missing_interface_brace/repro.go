package interfacefuzz
type Bound023 interface[
 ~func() int | ~*string | ~[]string | ~[3]string
}
func Scan023[T Bound023](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
