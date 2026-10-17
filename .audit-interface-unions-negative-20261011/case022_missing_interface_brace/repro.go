package interfacefuzz
import "time"
type Bound022 interface[
 time.Duration /* variant */ | Pair[int,string] /* variant */ | Pair[Box[int],Box[string]]
}
func Scan022[T Bound022](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
