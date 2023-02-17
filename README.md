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

`mbcs.ACP` is the current codepage.

#### On Windows

```
$ chcp 932
$ go run examples\AnsiToUtf8.go < testdata\jugemu-cp932.txt | nkf32 --guess
UTF-8 (LF)
```

#### On Linux

```
$ env LC_ALL=ja_JP.Shift_JIS go run examples/AnsiToUtf8.go < testdata/jugemu-cp932.txt | nkf --guess
UTF-8 (LF)
```

When OS is not Windows, the current encoding is judged with $LC_ALL and $LANG.

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

#### On Windows

```
$ chcp 932
$ go run examples\Utf8ToAnsi.go < testdata\jugemu-utf8.txt | nkf32 --guess
Shift_JIS (LF)
```

#### On Linux

```
$ env LC_ALL=ja_JP.Shift_JIS go run examples/Utf8ToAnsi.go < testdata/jugemu-utf8.txt | nkf --guess
Shift_JIS (LF)
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
    sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.NewDecoder(mbcs.ACP)))
    for sc.Scan() {
        fmt.Println(sc.Text())
    }
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

#### On Windows

```
$ chcp 932
$ go run examples\NewDecoder.go < testdata\jugemu-cp932.txt | nkf32 --guess
UTF-8 (LF)
```

#### On Linux

```
$ env LC_ALL=ja_JP.Shift_JIS go run examples/NewDecoder.go < testdata/jugemu-cp932.txt  | nkf --guess
UTF-8 (LF)
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
    sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.NewEncoder(mbcs.ACP)))
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

#### On Windows

```
$ chcp 932
$ go run examples\NewEncoder.go < testdata/jugemu-utf8.txt  | nkf32 --guess
Shift_JIS (LF)
```

#### On Linux

```
$ env LC_ALL=ja_JP.Shift_JIS go run examples/NewEncoder.go < testdata/jugemu-utf8.txt  | nkf --guess
Shift_JIS (LF)
```
