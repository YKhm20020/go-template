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

# 動作確認

http://localhost:8080 で、以下のようにサンプルのメッセージが取得できていることを確認する。

```
{
"message": "Hello, Gin + MySQL + Docker + Air"
}
```

http://localhost:8080/accounts で、以下のようにサンプルデータを一覧表示できていることを確認する。

```
[
    {
        "ID": 1,
        "Name": "a1_name",
        "Password": "a1_password"
    },
    {
        "ID": 2,
        "Name": "a2_name",
        "Password": "a2_password"
    },
    {
        "ID": 3,
        "Name": "a3_name",
        "Password": "a3_password"
    },
    {
        "ID": 4,
        "Name": "a4_name",
        "Password": "a4_password"
    }
]
```

http://localhost:8080/accounts/1 で、ID が 1 のサンプルデータを表示できていることを確認する。

```
[
    {
        "ID": 1,
        "Name": "a1_name",
        "Password": "a1_password"
    },
]
```
