## 1.1 Docker とは
Docker とは・・・

- コンテナ型仮想化技術を実現するために実行される常駐アプリケーション
- それを操作するためのコマンドフロントインターフェイス

上記から成るプロダクトのこと。
アプリケーションのデプロイに特化しており、コンテナを中心とした開発・運用を可能にする。

### ユースケース
1. 軽量な仮想環境として検証に利用する
アプリケーションなどの検証の際、一から環境構築を行うのは手間・面倒。
Docker なら簡単にセットアップができ、検証に必要な環境構築を簡略化できる。
また、Dockerコンテナは仮想化ソフトウェアと比較して、より少ないオーバーヘッドで動作する。

- 操作の簡便さ
- コンテナの軽量さ

→上記のような利点から、Dockerはローカル環境での開発環境の再現に用いられるようになった。

2. 本番環境への展開やアプリケーションプラットフォームとして利用する
Dockerは既存の仮想化ソフトウェアと比較して軽量に動作するため、検証環境だけでなく、実際のアプリケーションでもコンテナが利用できる。

- ローカルのDocker環境で実行しているコンテナを、別のサーバーにあるDocker環境にデプロイする
- サーバーのDocker環境で動作するコンテナをローカルに持ってくる
などが可能なため、ポータビリティに優れていると言える。

3. インストールに癖があるコマンドラインツールをDockerコンテナとして入手する
ホストを汚さずにコマンドラインツールを即座に実行できる。

4. 依存するさまざまなライブラリやツールをDockerコンテナに同梱して配布する
実行環境を問わずに高い動作再現性を持つスクリプトを実現できる。

5. HTTP負荷テストでworkerをDockerコンテナとして用意して、HTTPリクエスト数を上げる。

### Docker の苦手な部分
Docker コンテナの内部はLinux系OSのような構成をしているものが多くを占める。
が、コンテナはOSとしての振る舞いを完全に再現しているわけではない。
そのため、より厳格にLinux系OSとして振る舞う仮想環境を構築したい場合は、VMWare や VirtualBox などの仮想化ソフトウェアを利用するべき。

### 1.1.1 Dockerの歴史
2013年春
dotCloud社（現Docker社）のエンジニアであるSolomon Hykes氏がDocker をオープンソースソフトウェアとして公開。

Docker社は、オーケストレーションシステムであるFig（現在のDocker Compose）を初め、周辺ツールを次々に買収。

2014年〜
DockerのカンファレンスであるDocker Conが毎年開催されており、Dockerの推進とコミュニティの活性化に寄与している。

Docker関連のオープンソースプロダクトの開発競争も盛ん。

### 1.1.3 Dockerの基礎概念
#### コンテナ型仮想化技術
Dockerはコンテナ型仮想化技術を利用している。
コンテナ型仮想化技術自体は、Dockerが登場する前から存在しており、Docker以前ではLXC（Linux Containers）が有名だった。
Dockerは最初期はLXC、現在はrunCというランタイムを用いてコンテナ型仮想化を実現している。

コンテナ型仮想化技術では仮想化ソフトウェアなしにOSのリソースを隔離し、仮想OSにする。
この仮想OSをコンテナと呼ぶ。
コンテナを作り出すためのオーバーヘッドは、他の仮想化ソフトウェアと比較して少なく、高速に起動・終了でき、必要なマシンリソースも少なくて済む。

OS上にインストールした仮想化ソフトウェアを利用し、ハードウェアを演算により再現しゲストOSを作り出す仕組みはホストOS型の仮想化と呼ぶ。
（VMwere Player や Oracle VirtualBoxなど。。。）
コンテナ型仮想化に比べると、オーバーヘッドが大きくなりがち。

コンテナ型仮想化技術によって、コンテナを軽量に作成、利用、破棄できるのはDockerの重要な特徴の１つ。

#### アプリケーションにフォーカスしたDocker
LXCは、ホスト型仮想化技術よりパフォーマンス面で有利なため、システムコンテナとしての用途で一定の地位を確立した。
しかし、LXCでは複製したアプリケーションを別のLXCホストで実行しようとしても、LXCの設定差異により期待した動作が得られないなどの問題があった。

Docker がそのお悩みを解決します！
Dockerはアプリケーションのデプロイにフォーカスを置いており、LXCと比較して下記のような特徴がある。

- ホストに左右されない実行環境（Docker Engine による実行環境の標準化）
- DSL(Dockerfile)によるコンテナの構成やアプリケーション配置定義
- イメージのバージョン管理
- レイヤ構造を持つイメージフォーマット（差分ビルドが可能）
- Dockerレジストリ
- プログラマブルな各種API

Dockerfileによりコンテナの情報をコードで管理でき、このコードをベースに取得や配布の支援も行われており、再現性が保ちやすいのが特徴。

Dockerはアプリケーションと実行環境を同梱しているため、実行環境への依存問題を軽減している。
また、これによりアプリケーションのデプロイが簡単。
