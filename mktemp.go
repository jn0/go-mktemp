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
	"path"
	"strings"
)

const Macro = "XXXXXX"

func Template(base string) string {
	tmp := os.TempDir()
	return path.Join(tmp, base + "." + Macro)
}

type Error struct {
	error
	m string
}
func (self Error) Error() string {
	return self.m
}
func newError(message string) Error {
	var e = Error{m: message}
	return e
}

// char *mkdtemp(char *template);
func MkDTemp(template string) (name string, e error) {
	if !strings.HasSuffix(template, Macro) {
		e = newError("Invalid template")
		return
	}
	buf := make([]byte, 8192)
	buf = []byte(template)
	rc := C.mkdtemp((*C.char)(unsafe.Pointer(&buf[0])))
	if rc == nil {
		e = syscall.Errno(C.getErrno())
		return
	}
	return C.GoString(rc), nil
}

// int mkstemp(char *template);
func MkSTemp(template string) (file *os.File, e error) {
	if !strings.HasSuffix(template, Macro) {
		e = newError("Invalid template")
		return
	}
	buf := make([]byte, 8192)
	buf = []byte(template)
	rc := C.mkstemp((*C.char)(unsafe.Pointer(&buf[0])))
	if int(rc) == -1 {
		e = syscall.Errno(C.getErrno())
		return
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
