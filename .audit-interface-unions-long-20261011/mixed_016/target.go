package audit
type Bound interface { string | ~struct{ Field001 int } | ~struct{ Field002 int } | ~struct{ Field003 int } | ~struct{ Field004 int } | ~struct{ Field005 int } | ~struct{ Field006 int } | ~struct{ Field007 int } | ~struct{ Field008 int } | ~struct{ Field009 int } | ~struct{ Field010 int } | ~struct{ Field011 int } | ~struct{ Field012 int } | ~struct{ Field013 int } | ~struct{ Field014 int } | ~struct{ Field015 int } }
func StillScan(xs []int) {for range xs {println("seen")}}
