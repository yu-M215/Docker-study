## 2.3 Docker コンテナの操作
### 2.3.1 Docker コンテナのライフサイクル

「実行中」
「停止」
「破棄」

Docker コンテナは、上記の3つの状態に分類される。
これを Docker コンテナのライフサイクルと呼ぶ。
`docker container run` で起動された直後は「実行中」の状態。

#### 実行中
`docker container run` で指定された Docker イメージをもとにコンテナが作成され、 Dockerfile の `CMD` や `ENTRYPOINT` で定義されているアプリケーションの実行を開始する。
このアプリケーションが実行中なら、 Docker コンテナは実行中にあるということになる。
実行が完了すると、停止の状態に移行する。

#### 停止
ユーザーが明示的にコンテナを停止するか、コンテナで実行されているアプリケーションが正常・異常を問わず終了した場合に、コンテナは自動的に停止する。
停止により仮想環境としての役割を終えるが、ディスクにコンテナ終了時の状態は保持されているため、停止したコンテナは再実行可能。

#### 破棄
停止したコンテナは明示的に破棄しない限り、ディスクに残り続ける。
ディスクの占有を避けるため、完全に不要なコンテナは破棄することが望ましい。
ただし、一度破棄したコンテナを再び開始することはできないので、その点は注意。

### 2.3.2 docker container run -コンテナの作成と実行
`docker container run` は Docker イメージからコンテナを作成・実行するコマンド。

```
docker container run [options] イメージ名[:タグ] [コマンド] [コマンド引数...]
```

```
docker container run [options] イメージID [コマンド] [コマンド引数...]
```

コンテナをバックグラウンドで実行するには下記のようにする。

```
docker container run -d -p 9000:8080 example/echo:latest
```

`-p` オプションでホスト側の `9000` ポートからコンテナ側の `8080` ポートへポートフォワーディングしている。

#### docker container run 時に引数を与える
`docker container run` にコマンド引数を与えることで Dockerfile で指定していた `CMD` を上書きできる。

```
$ docker image pull alpine:3.7
# docker container run -it alpine:3.7 # シェルに入る
$ docker container run -it alpine:3.7 uname -a # uname -a で上書き。
```

#### 名前付きコンテナ
`docker container run` でコンテナを実行する際、 `docker container ls` で表示される `NAMES` には適当な単語で作られた名前が自動で作られる。

コンテナの停止などコンテナを制御するコマンドを実行する際、コンテナIDかコンテナ名を指定する必要がある。
コンテナ名をつけていない場合、停止の際、 `docker container ls` で ID もしくはコンテナ名を逐一確認することになってしまう。
`--name` オプションを使用してコンテナに任意の名前をつけることで、確認の手間を省くことができる。

```
docker container run --name [コンテナ名] [イメージ名]:[タグ]
```

#### コマンド実行時の頻出オプション
`-i`
Docker 起動後にコンテナ側の標準入力を繋ぎっぱなしにする。
このため、シェルに入ってコマンドを実行することなどができる。

`-t`
疑似端末を有効にする。
`-i` がつかないと端末を起動しても入力できないので `-i` とセットで使うことが多い。

`--rm`
コンテナ終了時にコンテナを破棄する。
1回走らせればその後保持しておく必要がないコマンドラインツールなどの実行時使うのに適している。

`-v`
ホストとコンテナ間で、ディレクトリ、ファイルを共有する時に使う。

### 2.3.3 docker container ls -コンテナの一覧
`docker container ls` は実行中のコンテナおよび終了したコンテナの一覧を表示するコマンド。

```
docker container ls [options]
```

- CONTAINER ID
コンテナに付与される一意のID

- IMAGE
コンテナ作成に使用された Docker イメージ

- COMMAND
コンテナで実行されているアプリケーションのプロセス

- CREATED
コンテナが作成されてから経過した時間

- STATUS
`Up`（実行中）、`Exited`（終了）といったコンテナの実行状態

- PORTS
ホストのポートとコンテナポートの紐付け（ポートフォワーディング）

- NAMES
コンテナにつけられた名前

#### コンテナIDだけを抽出する
`-q` オプションを付与すると、コンテナID（省略形）だけを抽出できる。

```
docker container ls -q
```

#### filter を使う
`--filter` オプションを利用すると、特定の条件に一致するものだけを抽出することができる。

```
docker container ls --filter "filter名=値"
```

```使用例
docker container ls --filter "name=echo1"
```

イメージで抽出する場合は `ancestor` フィルターを利用する。

```使用例
docker container ls --filter "ancestor=example/echo"
```

#### 終了したコンテナを取得する
`-a` オプションを付与することで、終了したコンテナも含めたコンテナの一覧を取得できる。
終了したコンテナの実行時の標準出力を参照したり、再実行するようなケースで利用する。

```
docker container ls -a
```

### 2.3.4 docker container stop -コンテナの停止
`docker container stop` コマンドで実行しているコンテナを終了する。

```
docker container stop コンテナIDまたはコンテナ名
```

```使用例
$ docker container run -t -d --name echo example/echo:latest
$ docker container stop echo
```

### 2.3.5 docker container restart -コンテナの再起動
`docker container restart` で停止したコンテナを再実行できる。

```
docker container restart コンテナIDまたはコンテナ名
```

### 2.3.6 docker container rm -コンテナの破棄
停止したコンテナをディスクから完全に破棄する場合は `docker container rm` コマンドを使用する。

```
docker container rm コンテナIDまたはコンテナ名
```

`docker container rm` コマンドでは実行中のコンテナを破棄することはできない。
`-f` オプションをつけると、実行中のコンテナを停止・削除まで行うことができる。

#### docker container run --rm で停止の際にコンテナを破棄する
`docker container run --rm` で停止の際にコンテナの破棄まで行うことができる。
このオプションが使われるユースケースは、下記のように、コマンドラインツールでのDocker コンテナを利用する場合など。

```
$ echo '{"version":100}' | docker container run -i --rm gihyodocker/jq:1.5 '.version'
```

### 2.3.7 docker container logs -標準出力の取得
`docker container logs` コマンドでは、実行している特定の Docker コンテナの標準出力を表示することができる。
標準出力されているものだけが表示されているため、コンテナの中でアプリケーションがファイルに出力したようなログは表示されない。

```
docker container logs [options] コンテナIDまたはコンテナ名
```

`-f` オプションをつけると、標準出力の取得をし続ける。

### 2.3.8 docker container exec -実行中コンテナでのコマンド実行
`docker container exec` コマンドで実行している Docker コンテナの中で任意のコマンドを実行できる。

```
docker container exec [options] コンテナIDまたはコンテナ名 コンテナで実行するコマンド
```

`sh` や `bash` をコンテナで実行するコマンドで渡すと、コンテナにSSHでログインしたかのようにコンテナ内部で操作をすることもできる。

```使用例
$ docker container run -t -d --name echo --rm example/echo:latest
$ docker container exec -it echo sh
```

### 2.3.9 docker container cp -ファイルのコピー
`docker container cp` を利用すると、コンテナ間、コンテナ・ホストかんでファイルをコピーできる。
`Dockerfile` の `COPY` はイメージビルド時にホストからファイルをコピーするために利用されるが、 `docker container cp` は実行中のコンテナ間でのファイルのやり取りのために使われる。

```
docker container cp [options] コンテナIDまたはコンテナ名:コンテナ内のコピー元 ホストのコピー先
```

```
docker container cp [options] ホストのコピー元 コンテナIDまたはコンテナ名:コンテナ内のコピー先
```