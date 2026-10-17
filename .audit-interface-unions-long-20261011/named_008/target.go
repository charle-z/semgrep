package audit
type Named000 int
type Named001 int
type Named002 int
type Named003 int
type Named004 int
type Named005 int
type Named006 int
type Named007 int
type Bound interface { Named000 | Named001 | Named002 | Named003 | Named004 | Named005 | Named006 | Named007 }
func StillScan(xs []int) {for range xs {println("seen")}}
