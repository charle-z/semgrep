package audit
type Named000 int
type Named001 int
type Named002 int
type Named003 int
type Named004 int
type Named005 int
type Named006 int
type Named007 int
type Named008 int
type Named009 int
type Named010 int
type Named011 int
type Named012 int
type Named013 int
type Named014 int
type Named015 int
type Bound interface { Named000 | Named001 | Named002 | Named003 | Named004 | Named005 | Named006 | Named007 | Named008 | Named009 | Named010 | Named011 | Named012 | Named013 | Named014 | Named015 }
func StillScan(xs []int) {for range xs {println("seen")}}
