[![Go Reference](https://pkg.go.dev/badge/github.com/nyaosorg/go-windows-mbcs.svg)](https://pkg.go.dev/github.com/nyaosorg/go-windows-mbcs)

go-windows-mbcs
===============

Convert between UTF8 and non-UTF8 character codes(ANSI) using Windows APIs: MultiByteToWideChar and WideCharToMultiByte.

Convert from ANSI-bytes to UTF8-strings
---------------------------------------

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
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```


Convert from UTF8-strings to ANSI-bytes
---------------------------------------

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
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

Use golang.org/x/text/transform
-------------------------------

### Convert from ANSI-reader to UTF8-reader

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "golang.org/x/text/transform"

    "github.com/nyaosorg/go-windows-mbcs"
)

func main() {
    sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.Decoder{CP: mbcs.ACP}))
    for sc.Scan() {
        fmt.Println(sc.Text())
    }
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

### Convert from UTF8-reader to ANSI-reader

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "golang.org/x/text/transform"

    "github.com/nyaosorg/go-windows-mbcs"
)

func main() {
    sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.Encoder{CP: mbcs.ACP}))
    for sc.Scan() {
        os.Stdout.Write(sc.Bytes())
        os.Stdout.Write([]byte{'\n'})
    }
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```
