package mktemp

import "testing"
import "os"
import "strings"

func TestMkTemp(t *testing.T) {
	t.Logf("test of MkTemp")

	var templ string = "tempFile.XXXXXX"
	const testString = "test string\n"

	file, e := MkSTemp(templ)
	if e != nil { t.Fatalf("Failed MkSTemp(%+q): %v", templ, e); }
	t.Logf("MkSTemp(%+q): file=%+q e=%v", templ, file.Name(), e)
	_, e = file.Write([]byte(testString))
	if e != nil { t.Fatalf("Cannot write to %#v", file); }
	file.Close()

	file, e = os.Open(file.Name())
	buf := make([]byte, 1024)
	l, e := file.Read(buf)
	if e != nil { t.Fatalf("Cannot read from %#v", file); }
	file.Close()
	if strings.Compare(string(buf[:l]), testString) != 0 {
		t.Fatalf("Inconsistent data in %+q: expected %+q got(%d) %+q",
			 file.Name(), testString, l, string(buf))
	}
	e = os.Remove(file.Name())
	if e != nil { t.Fatalf("Cannot remove %#v", file); }
	t.Logf("MkSTemp(%+q): name=%+q ok", templ, file.Name())

	name, e := MkTemp(templ)
	t.Logf("MkTemp(%+q): name=%+q e=%v", templ, name, e)
	if e != nil { t.Fatalf("Failed MkTemp(%+q): %v", templ, e); }
	e = os.Remove(name)
	if e != nil { t.Fatalf("Cannot remove %#v", name); }
	t.Logf("MkTemp(%+q): name=%+q ok", templ, name)

	templ = Template("tempDir")
	name, e = MkDTemp(templ)
	t.Logf("MkDTemp(%+q): name=%+q e=%v", templ, name, e)
	if e != nil { t.Fatalf("Failed MkDTemp(%+q): %v", templ, e); }
	e = os.Remove(name)
	if e != nil { t.Fatalf("Cannot remove %#v dir", name); }
	t.Logf("MkDTemp(%+q): name=%+q ok", templ, name)

	t.Logf("done")
}
