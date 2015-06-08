package mdbm

/*
#cgo CFLAGS: -I/tmp/install/include
#cgo LDFLAGS: -L/tmp/install/lib64 -lmdbm

#include "mdbm.h"
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

MDBM_ITER new_iter() {
  MDBM_ITER iter;
  MDBM_ITER_INIT(&iter);
  return iter;
}

datum new_datum() {
  datum d;
  return d;
}

*/
import "C"
import "unsafe"
import "fmt"
import "log"

type MdbmLib struct {
	db *C.MDBM
	// maybe no iter, leave in for now
	iter *C.MDBM_ITER
}

func NewMdbm(db_file string, flags int, mode int, psize int, presize int) *MdbmLib {
//	fmt.Printf("int MDBM_O_WRONLY: %d\n\n\n", C.MDBM_O_WRONLY)
//	fmt.Printf("int MDBM_O_RDWR: %d\n\n\n", C.MDBM_O_RDWR)
	c_db, err := C.mdbm_open(C.CString(db_file), C.int(flags), C.int(mode), C.int(psize), C.int(presize))
	if unsafe.Pointer(c_db) == nil {
		log.Fatal(err)
		fmt.Printf("mdbm failed: XXXXX %")
	}
	thing := C.new_iter()
	return &MdbmLib{c_db, &thing}
}

func (db MdbmLib) GetFirst() (string, string) {
	C.mdbm_lock(db.db)
	pair, err := C.mdbm_first_r(db.db, db.iter)
	C.mdbm_unlock(db.db)
	fmt.Printf("err: ", err)
	if pair.key.dsize == 0 {
		return "dsize is zero", "yes it is"
	}
	return string(C.GoBytes(unsafe.Pointer(pair.key.dptr), pair.key.dsize)), string(C.GoBytes(unsafe.Pointer(pair.val.dptr), pair.val.dsize))
}

func (db MdbmLib) Lock() (int, error) {
	ret, err := C.mdbm_lock(db.db)
	return int(ret), err
}
func (db MdbmLib) Unlock() (int, error) {
	ret, err := C.mdbm_unlock(db.db)
	return int(ret), err
}

func (db MdbmLib) FetchNoLock(key string) string {
	C.mdbm_lock(db.db)
	d_key := C.new_datum()
	d_val := C.new_datum()
	d_key.dptr = C.CString(key)
	d_key.dsize = C.int(len(key))
	iter := C.new_iter()
	C.mdbm_fetch_r(db.db, &d_key, &d_val, &iter)
	return C.GoStringN(d_val.dptr, d_val.dsize)
}

func (db MdbmLib) Fetch(key string) string {
	C.mdbm_lock(db.db)
	d_key := C.new_datum()
	d_val := C.new_datum()
	d_key.dptr = C.CString(key)
	d_key.dsize = C.int(len(key))
	iter := C.new_iter()
	C.mdbm_lock(db.db)
	C.mdbm_fetch_r(db.db, &d_key, &d_val, &iter)
	C.mdbm_unlock(db.db)
	return C.GoStringN(d_val.dptr, d_val.dsize)
}

func (db MdbmLib) Store(key string, val string, flags int) {
	iter := C.new_iter()
	d_key := C.new_datum()
	d_val := C.new_datum()

	d_key.dptr = C.CString(key)
	d_key.dsize = C.int(len(key))

	d_val.dptr = C.CString(val)
	d_val.dsize = C.int(len(val))

	C.mdbm_lock(db.db)
	C.mdbm_store_r(db.db, &d_key, &d_val, C.int(flags), &iter)
	C.mdbm_unlock(db.db)

}

func (db MdbmLib) Keys() []string {
	iter := C.new_iter()
	d_key := C.new_datum()
	ret := make([]string, 1)
	for {
		d_key = C.mdbm_nextkey_r(db.db, &iter)
		if d_key.dsize == 0 {
			break
		}
		ret = append(ret, C.GoStringN(d_key.dptr, d_key.dsize))
	}
	return ret
}

func (db MdbmLib) Close() {
	C.mdbm_close(db.db)
}

