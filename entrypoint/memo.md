ENTRYPOINT を指定しているので、 `go` を省略してコマンドを実行できる。

```
[yumatsui@YunoMacBook-Air:~/ghq/github.com/yu-M215/Docker-study/entrypoint][main]
$ docker container run ch02/golang:latest version
go version go1.19.13 linux/arm64
```

ENTRYPOINT はイメージの作成者側でコンテナの用途をある程度制限したい場合に活用できる。
※ただし、`docker container run --entrypoint` で実行時に上書きすることも可能。