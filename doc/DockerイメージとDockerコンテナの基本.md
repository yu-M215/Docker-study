## 2.1 コンテナでアプリケーションを実行する

Dockerイメージ
Dockerコンテナを構成するファイルシステムや、実行するアプリケーションや設定をまとめたもの。
コンテナを作成するために利用されるテンプレートとなるもの。

Dockerコンテナ
Dockerイメージをもとに作成され、具現化されたファイルシステムとアプリケーションが実行されている状態。

### 2.1.1 DockerイメージとDockerコンテナの基本

`gihyodocker/echo:latest` というDockerイメージの取得。

```
docker image pull gihyodocker/echo:latest
```

ダウンロードしたイメージの実行。

```
docker container run -t -p 9000:8080 gihyodocker/echo:latest
```

リクエスト実行。
※ `curl` だとうまくいかなかったので、ブラウザでアクセスしたら期待通りの挙動をした。

```
curl http://localhost:9000
```