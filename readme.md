go-windows-mbcs
===============

A simple wrapper for MultiByteToWideChar and WideCharToMultiByte.

Convert from ANSI to UTF8
--------------------

```go
func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        text, err := mbcs.AtoU(sc.Bytes(), mbcs.ACP)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
        fmt.Println(text)
    }
}
```

Convert from UTF8 to ANSI
----------------

```go
func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        bytes, err := mbcs.UtoA(sc.Text(), mbcs.ACP)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
        os.Stdout.Write(bytes)
        os.Stdout.Write([]byte{'\n'})
    }
}
```
