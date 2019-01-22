### gdbt - 非公式 idobata CLI クライアント

プログラマのためのSNS、Idobataの非公式クライアント。   
[npm版](https://www.npmjs.com/package/idbt)をGoで再実装、機能追加した。   
ターミナルソフトや各種IDE付属のターミナルから実行し、仕事中に流し見をすることを想定している。

### インストール

実行形式ファイルで配布(予定)。   
OSに対応した実行形式ファイルをdistフォルダから取得し、`$PATH` の通ったところに配置する。

#### アンインストール

実行形式ファイルを削除、ユーザディレクトリ直下の設定ファイルを削除する。

### デモ

発表の際に実際に使う予定。

### 使い方

```
# global option "--help"
$ gdbt --help
$ gdbt post --help
$ gdbt draft -h # shorthand

# setup
$ gdbt setup
$ gdbt init # alias
$ gdbt i    # shorthand

# room
$ gdbt room
$ gdbt r # shorthand
$ gdbt room --reload
$ gdbt room -r # shorthand
$ gdbt room --show
$ gdbt room -s # shorthand

# list
$ gdbt list
$ gdbt l # shorthand

# gdbt post
$ gdbt p # shorthand
$ gdbt post --message=<message string>
$ gdbt post -m <message string> # shorthand
$ gdbt post --draft
$ gdbt post -d # shorthand

# draft
$ gdbt draft
$ gdbt d # shorthand
$ gdbt draft --post
$ gdbt draft -p # shorthand
$ gdbt draft --show
$ gdbt draft -s # shorthand
```
