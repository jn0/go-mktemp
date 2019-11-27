package mktemp

/*
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>

static inline int getErrno(void) { return errno; }
*/
import "C"

import (
	"unsafe"
	"syscall"
	"os"
)

// int mkstemp(char *template);
func MkSTemp(template string) (file *os.File, e error) {
	buf := make([]byte, 8192)
	buf = []byte(template)
	rc := C.mkstemp((*C.char)(unsafe.Pointer(&buf[0])))
	if int(rc) == -1 {
		e = syscall.Errno(C.getErrno())
		return nil, e
	}
	return os.NewFile(uintptr(rc), string(buf)), nil
}

// Mimics `char *mktemp(char *template); // Never use this function` via MkSTemp()
func MkTemp(template string) (name string, e error) {
	file, e := MkSTemp(template)
	if e != nil {
		return
	}
	name = file.Name()
	file.Close()
	return
}

/* EOF */
