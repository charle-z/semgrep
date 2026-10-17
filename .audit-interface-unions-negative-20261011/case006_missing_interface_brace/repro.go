package interfacefuzz
type Bound006 interface[
 ~[]string | Pair[Box[int],Box[string]]
}
type Alias006[T Bound006] = []T
func Scan006[T Bound006](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
