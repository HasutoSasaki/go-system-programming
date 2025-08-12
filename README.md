# go-system-programming

このリポジトリは「Goならわかるシステムプログラミング」の動かす用リポジトリです。

## 概要

ASCII.jpで連載中の記事「Goならわかるシステムプログラミング」のサンプルコードを実際に動かして学習するためのリポジトリです。

記事URL: https://ascii.jp/serialarticles/1235262/

## 構成

- `pkg/reader/` - Reader関連のサンプルコード
- `pkg/writer/` - Writer関連のサンプルコード
- `pkg/filecopy/` - ファイルコピー関連のサンプルコード
- `pkg/png_analysis/` - PNG解析関連のサンプルコード
- `pkg/randomfile/` - ランダムファイル生成関連のサンプルコード
- `pkg/writezip/` - ZIP書き込み関連のサンプルコード
- `pkg/zipdownload/` - ZIPダウンロード関連のサンプルコード
- `sample/` - サンプルファイル

## 実行方法

各パッケージのmain.goファイルを実行してください：

```bash
# 例: Readerの基本的な使い方
go run pkg/reader/main.go

# 例: Writerの基本的な使い方
go run pkg/writer/main.go
```

## 要件

- Go 1.16以上

## ライセンス

MIT License
