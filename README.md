# mdprv

Markdown ファイルをローカルの HTTP サーバーで手軽にプレビューするための CLI ツールです。

## 主な機能

- GFM（GitHub Flavored Markdown）および脚注（Footnotes）の変換
- GitHub 形式のアラート表示（Note、Tip、Important、Warning、Caution）
- highlight.js によるコードブロックのシンタックスハイライト
- 空きポートの自動割り当てによるポート衝突の防止
- 単一バイナリによる簡単なセットアップ

## インストール

Go 1.27 以降がインストールされている環境で、以下のコマンドを実行します。

```bash
go install .
```

GitHub Releases から各 OS 向けの事前ビルド済みバイナリをダウンロードして利用することもできます。

## 使い方

プレビューしたい Markdown ファイルのパスを指定して実行します。

```bash
# 基本的な実行方法
mdprv README.md

# サブコマンドを指定して実行する方法
mdprv preview README.md
```

コマンドを実行するとローカル HTTP サーバーが起動し、ブラウザでアクセスするための URL がコンソールに出力されます。

```text
2026/09/26 17:00:00 Create a preview for README.md
2026/09/26 17:00:00 Starting the server: http://127.0.0.1:54321
```

表示された URL（例: `http://127.0.0.1:54321`）をブラウザで開くことで、整形された HTML を確認できます。
サーバーを終了する場合は、ターミナルで `Ctrl + C` を押します。

## ライセンス

[MIT](LICENSE)
