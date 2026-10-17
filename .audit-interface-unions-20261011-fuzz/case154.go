package interfacefuzz
type Embedded154 interface{~[]rune | ~map[string]rune}
type Bound154 interface{
Embedded154
 ~chan string /* variant */ | int /* variant */ | ~complex64 /* variant */ | ~[3]string /* variant */ | ~struct{ B string }
}
func Scan154[T Bound154](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
