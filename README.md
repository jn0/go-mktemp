# go-mktemp (mktemp(3) in Go)

I was really frustrated found nothing for mk[s]temp(3) in Go.

So, here is sorta leverage.

```
package mktemp // import "github.com/jn0/go-mktemp"

func MkSTemp(template string) (file *os.File, e error) // use file.Name() ...
func MkTemp(template string) (name string, e error) // file `name` is closed
```

The `MkTemp()` mimics deprecated `mktemp(3)` by wrapping kosher `mkstemp(3)` call.

# EOF #
