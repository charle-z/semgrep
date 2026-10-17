package interfacefuzz
type Bound033 interface{
 Pair[string,int] /* variant */ | Pair[Box[int],Box[string]] /* variant */ | string /* variant */ | ~[2]int | 
}
type Alias033[T Bound033] = []T
func Scan033[T Bound033](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
