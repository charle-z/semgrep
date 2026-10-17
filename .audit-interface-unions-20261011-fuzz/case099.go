package interfacefuzz
type Bound099 interface{
 ~map[string]bool /* variant */ | string /* variant */ | Pair[string,int] /* variant */ | ~[]string /* variant */ | ~[3]string
}
type Alias099[T Bound099] = []T
func Scan099[T Bound099](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
