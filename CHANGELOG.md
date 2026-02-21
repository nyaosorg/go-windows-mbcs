Changelog
=========

v0.4.4
------
Jun 8, 2025

- (#1) refactor: windows error handing & length check by @Snshadow
- Updated golang.org/x/sys to v0.30.0 (the latest version that can be built with Go 1.20)
- Updated golang.org/x/text to v0.22.0 (the latest version that can be built with Go 1.20)

These versions were chosen to ensure that this package remains buildable with Go 1.20, because some dependent tools are expected to support older Windows environments, including Windows 7, 8, 10, 11, and Windows Server 2008 R2 and later.

Thanks to @Snshadow

v0.4.3
------
Jun 8, 2025

- Updated `golang.org/x/sys` to v0.28.0
- Updated `golang.org/x/text` to v0.21.0
- Added benchmark code

v0.4.2
------
Apr 30, 2023

+ Fix: that transform.Transform types returned by the functions NewEncoder and NewDecoder fails when source text is larger than 4096 bytes  
   (See also [nyaosorg/nyagos issue 431](https://github.com/nyaosorg/nyagos/issues/431) )
+ Remove `Deprecated` from Filter types and methods

v0.4.1
------
Apr 25, 2023

- Unexpose types ( `Decoder` / `Encoder` ) that should not be.
    - Some packages using them could not be built on Linux even if it could be built on Windows.
    - Use the functions `NewDecoder` or `NewEncoder` instead.

v0.4.0
------
Feb 17, 2023

- Add functions: `NewEncoder(CODEPAGE)` and `NewDecoder(CODEPAGE)`
- On Linux, get the current encoding from the environment variable $LC\_ALL and $LANG.

Now supporting encoding on Linux

```
932:   japanese.ShiftJIS,
936:   simplifiedchinese.GBK,
949:   korean.EUCKR, // Unified Hangul Code
950:   traditionalchinese.Big5,
951:   traditionalchinese.Big5, // Big5-HKSCS
50222: japanese.ISO2022JP,
51932: japanese.EUCJP,
51949: korean.EUCKR,
52936: simplifiedchinese.HZGB2312,
65001: UTF8
```

v0.3.1
------
Feb 7, 2023

-  When the code page is ACP or 65001 on Linux, the text is passed without conversion instead of an error.
-  Add tests for Linux (codepage=65001 only)

v0.3.0
------
Feb 7, 2023

- Add new types: `Decoder`, `Encoder`, `AutoDecoder` as implements of [transform.Transformer](https://pkg.go.dev/golang.org/x/text@v0.6.0/transform#Transformer)
- Deprecate the type: `Filter` and its methods and constructor

v0.2.0
------
Feb 7, 2023

Rename methods:

- AtoU() -&gt; AnsiToUtf8()
- UtoA() -&gt; Utf8ToAnsi()

Old name still works, but will be phased out eventually.

v0.1.0
------
Feb 7, 2023

The version used in [nyagos 4.4.13_1](https://github.com/nyaosorg/nyagos/releases/tag/4.4.13_1)
