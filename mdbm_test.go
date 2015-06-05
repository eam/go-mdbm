package mdbm

import "testing"
import "fmt"

func TestThing(t *testing.T) {
	//x := NewMdbm("/tmp/go.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	key, val := x.GetFirst()
	fmt.Printf("\n\ngot key: %s val: %s\n\n", key, val)

	result := x.Fetch("zzz")
	fmt.Printf("fetching 'zzz' result: %s\n", result)

	x.Store("aaa", "bbb2", MDBM_REPLACE)

	result2 := x.Fetch("aaa")
	fmt.Printf("fetching 'aaa' result: %s\n", result2)
}
