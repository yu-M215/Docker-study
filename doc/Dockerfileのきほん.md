## Dockerfile の基本

### FROM

Docker イメージのベースとなるイメージを指定する。
Dockerfile でイメージをビルドする際、FROM で指定されたイメージをダウンロードしてから実行される。
（Docker Hub で公開されているイメージ）
Docker はデフォルトで、FROM の取得先として DockerHub のレジストリを参照する。
各イメージにはタグと呼ばれる識別子があり、バージョンなどを意味している。

### RUN

Docker イメージビルド時に、Docker コンテナ内で実行するコマンドを定義する。
RUN の引数には Docker コンテナ内で実行するコマンドをそのまま指定する。

### COPY

Docker を動作させているホストマシン上のファイルやディレクトリを Docker コンテナ内にコピーするためのインストラクション。
ADD は COPY とは用途が異なるため注意。

### CMD

Docker コンテナとして実行する際に、コンテナ内で実行するプロセスを指定。
イメージをビルドするための RUN に対して、CMD はコンテナ起動時に一度実行される。
RUN でアプリケーションの更新や配置、CMD でアプリケーションそのものを動作させるイメージ。

```bash
go run /echo/main.go
```

```Dockerfile
CMD ["go", "run", "/echo/main.go"]
```

1 つのコマンドを空白区切りで分割して配列化した形式で、CMD に記述する。

#### CMD の実行時上書き

CMD で指定した命令は `docker container run` の指定で実行時に上書きが可能。

```
docker container run $(docker image build -q .) echo yay
```

### ENTRYPOINT
コンテナのコマンド実行の仕方を工夫できる。
ENTRYPOINT は CMD と同じくコンテナ内で実行するプロセスを指定するためのインストラクション。
ENTRYPOINTを指定すると、CMDの引数は、ENTRYPOINTで実行するファイルへの引数となり、こんなてなが実行するデフォルトのプロセスを指定できる。

### LABEL
イメージの作者名記入などに使う。MAINTAINERというインストラクションが以前はあったが、deprecated。

### ENV
Dockerfile をもとに生成したDockerコンテナ内で使える環境変数を指定する。

### ARG
ビルド時に情報を埋め込むために使う。イメージビルドの時だけ使用できる一時的な環境変数。
ビルド時にARGに渡す値を指定している↓

```
[yumatsui@YunoMacBook-Air:~/ghq/github.com/yu-M215/Docker-study/other_instruction][main]
$ docker image build --build-arg builddate=today -t example/others .
```

環境変数にコマンドで私た値が反映されている。

```
[yumatsui@YunoMacBook-Air:~/ghq/github.com/yu-M215/Docker-study/other_instruction][main]
$ docker container run example/others
HOSTNAME=d61980168f7a
SHLVL=1
HOME=/root
BUILDFROM=from Alpine
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
PWD=/
BUILDDATE=today
```