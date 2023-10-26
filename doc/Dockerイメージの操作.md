## 2.2 Dockerイメージの操作

Docker コマンドのヘルプを出す。

```
docker help
```

Docker のコマンドラインツールはサブコマンド構成になっている。
下記のような形式でコマンドを実行する。

```
docker <COMMAND> <SUBCOMMAND>
```

イメージ操作に関するコマンドは `image` を使用する。

```
docker image --help
```

`docker image build --help` のように実行すると、サブコマンドのヘルプも参照可能。
`docker build` は `docker image build` コマンドのエイリアス。

### 2.2.1
`docker image build` は Dockerfile をもとに Docker イメージを作成するコマンド

```
docker image build -t イメージ名[:タグ名] Dockerfile配置ディレクトリのパス
```

`-t` はイメージ名とタグ名を指定するオプション。
`docker image build` では、必ず Dockerfile を与える必要があるため、ディレクトリに Dockerfile が存在しないと実行できない。

```Dockerfileがカレントディレクトリにある場合
docker image build -t example/echo:latest .
```

`-f` オプション
`docker image build` コマンドはデフォルトで `Dockerfile` という名前の Dockerfile を探す。
そのため、`Dockerfile` とは異なる名前のファイル名を利用したい場合は、 `-f` オプションを利用する。

```使用例
docker image build -f Dockerfile-test -t example/echo:latest .
```

`--pull` オプション
`docker image build` でイメージをビルドする際、 Dockerfile の `FROM` で指定されているイメージを一度レジストリからダウンロードして、それをベースイメージにして新しくイメージをビルドする。
一度取得した Docker イメージは削除しない限りホスト内に保持される。
`--pull` オプションを使用することで、 `docker image build` 時にベースイメージを強制的に再取得させることが可能。

```使用例
docker image build --pull=true -t example/echo:latest .
```

### 2.2.2 docker search -イメージの検索
Docker Hub では、ユーザーや組織が Github と同様にリポジトリを持つことができ、リポジトリでそれぞれの Docker イメージを管理している。
`docker search` コマンドを利用すると、Docker Hub のレジストリに登録されているリポジトリを検索できる。

```
docker search [options] 検索キーワード
```

`--limit` を指定することで表示件数を制限することもできる。

検索結果は STARS の降順に表示される。
`docker search` ではリポジトリの検索はできるが、Docker イメージのタグを取得することまではできない。
イメージのタグを知りたい際は、Docker Hub のリポジトリページの `Tags` から参照するか、APIを使用する。

```
curl -s 'https://hub.docker.com/v2/repositories/library/golang/tags/?page_size=10' | jq -r '.results[].name'
```

### 2.2.3 docker image pull -イメージの取得
Docker レジストリから Docker イメージをダウンロードしてくるには、 `docker image pull` を利用する。

```
docker image pull [options] リポジトリ名[:タグ名]
```

`docker image pull` でダウンロードしてきたイメージは、そのまま Docker コンテナとして利用できる。

### 2.2.4 docker image ls -イメージの一覧
`docker image ls` では、コマンドの実行対象である Docker ホストに保持されているイメージの一覧を表示する。
`docker image pull` でダウンロードしてきたイメージだけではなく、 `docker image build` でビルドしたイメージもDocker ホストのディスクに保持される。
※ Docker ホスト = Docker デーモンを実行しているホスト環境のこと。

```
docker image ls [options] [リポジトリ[:タグ]]
```

### 2.2.5 docker image tag -イメージのタグ付け
`docker image tag` は Docker イメージの特定のバージョンにタグ付けを行う。

#### Docker イメージのバージョン
Docker イメージのバージョン = イメージIDのこと

#### イメージIDへのタグ付け
`docker image tag` はイメージIDにタグ名という形で別名を与えるコマンドということになる。
タグは特定のイメージを参照しやすくするためのエイリアス。
Docker イメージはビルドの度に生成され、内容に応じてイメージIDを持ち、新しいイメージには新しいIDが振られる。

Docker イメージのタグは、ある特定のイメージIDを持つ Docker イメージを識別しやすくするために利用される。

タグを指定しないで `docker image build` を実行すると、デフォルトで `latest` のタグがつく。
タグをつけないで `docker image build` で一度ビルドしたイメージを、再度ビルドすると、イメージIDは更新されるが、タグは `latest` が引き継がれ、古いもののタグは `<none>` になる。

```
docker image tag 元イメージ名[:タグ] 新イメージ名[:タグ]
```

```使用例
docker image tag example/echo:latest example/echo:0.1.0
```

### 2.2.6 docker image push -イメージの公開
`docker image push` コマンドは保持している Docker イメージを Docker Hub などのレジストリに登録できる。

```
docker image push [options] リポジトリ名[:タグ]
```

### Docker Hub 
Docker Hub は Docker 社が管理している Docker レジストリ。
既に Docker Hub で公開されているイメージを利用するだけであれば、サインアップなしで利用可能。
