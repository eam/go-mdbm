package mdbm

import "testing"
import "fmt"

func TestFirst(t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	key, val := x.GetFirst()
	if key != "zzz" {
		t.Fatal("first key isn't 'zzz'")
	}
	if val != "qqq" {
		t.Fatal("first value isn't 'qqq'")
	}
}


func TestFetch(t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	result := x.Fetch("zzz")
	if result != "qqq" {
        t.Fatal("fetched value isn't 'qqq'")
    }

	x.Store("aaa", "bbb2", MDBM_REPLACE)

	result = x.Fetch("aaa")
	if result != "bbb2" {
		t.Fatal("stored aaa => bbb2 didn't return bbb2")
	}
}

func TestClose(t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	x.Close()
}


func TestKeys (t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	keys := x.Keys()
	fmt.Println("TestKeys:")
	fmt.Println(keys)
}

func TestFetchMissing (t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	result := x.Fetch("This key doesn't exist")
	fmt.Println(result)
	if result != "" {
		t.Fatal("Fetching missing key not nil")
	}
}

func TestTwoLocks (t *testing.T) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	_, err := x.Lock()
	_, err = x.Unlock()
	if err != nil {
		t.Errorf("First unlock raised error")
	}
	_, err = x.Unlock()
	if err == nil {
		t.Errorf("Second unlock did NOT raise error")
	}
}


func BenchmarkFetch (b *testing.B) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	for i := 0; i < b.N; i++ {
		_ = x.Fetch("aaa")
	}
}

func BenchmarkFetchNoLock (b *testing.B) {
	x := NewMdbm("/tmp/zzz-qqq.mdbm", MDBM_O_RDWR|MDBM_O_CREAT, 0644, 0, 0)
	defer x.Close()
	for i := 0; i < b.N; i++ {
		_ = x.FetchNoLock("aaa")
	}
}
