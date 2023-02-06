go-windows-mbcs
===============

Convert between UTF8 and non-UTF8 character codes(ANSI) using Windows APIs: MultiByteToWideChar and WideCharToMultiByte.

Convert from ANSI to UTF8
-------------------------

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/nyaosorg/go-windows-mbcs"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        text, err := mbcs.AnsiToUtf8(sc.Bytes(), mbcs.ACP)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
        fmt.Println(text)
    }
}
```

Convert from UTF8 to ANSI
-------------------------

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/nyaosorg/go-windows-mbcs"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        bytes, err := mbcs.Utf8ToAnsi(sc.Text(), mbcs.ACP)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
        os.Stdout.Write(bytes)
        os.Stdout.Write([]byte{'\n'})
    }
}
```

Filter from ANSI/UTF8 to UTF8
-----------------------------

`Filter` is the class like bufio.Scanner but judges whther the code is UTF8 or ANSI automatically.

```go
func main() {
    filter := mbcs.NewFilter(os.Stdin, mbcs.ACP)
    for filter.Scan() {
        println(filter.Text())
    }
}
```