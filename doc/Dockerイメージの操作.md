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

