Changelog
=========

v0.4.4
------
Jun 8, 2025

- @Snshadow さんによるプルリクエスト #1 「WIndows APIのエラー処理と長さチェックについてのリファクタリング」のマージ
- golang.org/x/sys を v0.30.0 へ更新（Go 1.20 でビルドできる最終バージョン）
- golang.org/x/text を v0.22.0 へ更新 （Go 1.20 でビルドできる最終バージョン）
  
このパッケージを使うツールは Windows 7, 8, 10, 11, Server2008R2以降での動作も配慮する必要があるため、本パッケージも Go 1.20 でビルドできるようにする必要がありました。そのため、sys も text も最新バージョンではなく、1.20 で利用できる最終バージョンを用いることにしています

Thanks to @Snshadow

v0.4.3
------
Jun 8, 2025

- `golang.org/x/sys` を v0.28.0 へ更新
- `golang.org/x/text` を v0.21.0 へ更新
- ベンチマークコードを追加

v0.4.2
------
Apr 30, 2023

+ NewEncoder / NewDecoder 関数が返す transform.Transfomer 型が、ソーステキストが 4096 バイトよりも大きい時に失敗する問題を修正
   ( [nyaosorg/nyagos issue 431](https://github.com/nyaosorg/nyagos/issues/431) も参照のこと )
+ Filter型とそのメソッドから `Deprecated` を外した。

v0.4.1
------
Apr 25, 2023

- 公開すべきでなかった型 ( `Decoder` / `Encoder` ) を非公開 ( `_Decoder` / `_Encoder` ) とした
    - それらを使っているパッケージが Windows ではビルドできていても、Linux ではビルドできませんでした。
    - かわりに `NewDecoder` もしくは `NewEncoder` を使ってください。

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

The version used in [nyagos 4.4.13\_1](https://github.com/nyaosorg/nyagos/releases/tag/4.4.13_1)
