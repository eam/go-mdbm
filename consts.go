package mdbm

/*
#cgo CFLAGS: -I/tmp/install/include
#cgo LDFLAGS: -L/tmp/install/lib64 -lmdbm

#include "mdbm.h"
*/
import "C"

const MDBM_O_CREAT = C.MDBM_O_CREAT
const MDBM_O_TRUNC = C.MDBM_O_TRUNC
const MDBM_O_FSYNC = C.MDBM_O_FSYNC
const MDBM_O_ASYNC = C.MDBM_O_ASYNC
