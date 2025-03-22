# 本リポジトリについて

本リポジトリは、Go で Web バックエンドサーバーを構築するためのテンプレートリポジトリです。

# 使用技術

- Go
- Gin
- Air
- MySQL
- Docker

# 環境構築

1. Docker をインストール
1. 本リポジトリをクローン、もしくは本リポジトリをテンプレートとしてリポジトリを作成
1. `docker compose build` で Docker イメージをビルド

# 必要なコマンド

下記コマンドの実行で、Gin, MySQL のサーバーを起動する。

```
docker compose up
```

サーバー起動後、http://localhost:8080 で動作を確認できる。services:
app:
build:
context: .
dockerfile: docker/go/Dockerfile
depends_on: - db
volumes: - ./src/app:/src/app
ports: - "8080:8080"
tty: true

サーバー停止は、以下のコマンドを実行する。

```
docker compose down
```
