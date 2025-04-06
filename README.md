# 本リポジトリについて

本リポジトリは、Go で Web バックエンドサーバーを構築するためのテンプレートリポジトリです。

# 使用技術

- Go
- Gin
- Gorm
- Air
- MySQL
- Docker

# 環境構築

1. Docker をインストール
1. 本リポジトリをクローン、もしくは本リポジトリをテンプレートとしてリポジトリを作成
1. .env ファイルを作成 (テンプレートは後述)
1. `docker compose build` で Docker イメージをビルド

# 必要なコマンド

下記コマンドの実行で、Gin, MySQL のサーバーを起動する。

```
docker compose up
```

サーバー起動後、http://localhost:8080 で動作を確認できる。

サーバー停止は、以下のコマンドを実行する。

```
docker compose down
```

キャッシュなしでの再ビルドは、以下のコマンドを実行する。

```
docker compose build --no-cache
```

# .env ファイルテンプレート

以下のように .env ファイルを作成するとよい。

```
DB_NAME=go-template-db
DB_ROOT_PASSWORD=rootpassword
DB_USER=db-user
DB_PASSWORD=db-password
```
