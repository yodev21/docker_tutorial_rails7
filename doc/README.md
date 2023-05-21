# 作業手順

# ファイル作成
下記ファイルを作成

```
Dockerfile
docker-compose.yml
Gemfile
Gemfile.lock
```

```Dockerfile
# イメージ名にRuby(Ver2.6.5)の実行環境のイメージを指定
FROM ruby:2.6.5

# パッケージのリストを更新しrailsの環境構築に必要なパッケージをインストール
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs

# プロジェクト用のディレクトリを作成
RUN mkdir /myapp

# ワーキングディレクトリに設定
WORKDIR /myapp

# プロジェクトのディレクトリにコピー
COPY Gemfile /myapp/Gemfile
COPY Gemfile.lock /myapp/Gemfile.lock

# bundle install実行
RUN bundle install

# ビルドコンテキストの内容を全てmyappにコピー
COPY . /myapp
```

```docker-compose.yml
version: '3'
services:
  db:
    # postgresのイメージを取得
    image: postgres
    environment:
      POSTGRES_USER: 'postgresql'
      POSTGRES_PASSWORD: 'postgresql-pass'
    restart: always
    volumes:
      - pgdatavol:/var/lib/postgresql/data
  web:
    # Dockerfileからイメージをビルドして使用
    build: .
    # コンテナ起動時に実行
    command: bundle exec rails s -p 3000 -b '0.0.0.0'
    # カレントディレクトリを/myappにバインドマウント
    volumes:
      - .:/myapp
    # 3000で公開して、コンテナの3000へ転送
    ports:
      - "3000:3000"
    # Webサービスを起動する前にdbサービスを起動
    depends_on:
      - db
# データ永続化のためにpgdatabolのvolumeを作成し、postgresqlのデータ領域をマウント
volumes:
  pgdatavol:
```

```Gemfile
source 'https://rubygems.org'
gem 'rails', '5.2.4.2'
```

```Gemfile.lock
```

## railsアプリケーション作成

```
docker-compose run web rails new . --force --database=postgresql
```

railsプロジェクトに使用するデータベースの設定ファイルを修正

```database.yml
default: &default
  adapter: postgresql
  encoding: unicode
  # -------- 追加 --------
  host: db
  username: postgresql
  password: postgresql-pass
  # -------- ここまで --------
```

# デタッチモード（バックグラウンド）で起動

```
docker-compose up -d
```

# bundle installが反映されない場合の対応

```
docker-compose build --no-cache
```

# データベース作成コマンド

```
docker-compose run web rails db:create
```

# Scaffoldにて簡易的なアプリケーション作成

```
docker-compose run web bin/rails g scaffold User name:string
```

```
docker-compose run web bin/rails db:migrate
```

http://localhost:3000/users

# コンテナの停止

```
docker-compose stop
```

# コンテナ削除

```
docker-compose down
```

# コンテナ内に移動

```
docker-compose run web bash
```

# Go言語を使用してみる

ディレクトリ移動

```
cd doc/golang/
```

```Dockerfile
FROM golang:1.9

RUN mkdir /echo
COPY main.go /echo
CMD ["go", "run", "/echo/main.go"]
```

# イメージのビルド

```
docker image build -t example/echo:latest .
```

# イメージの確認

```
docker image ls
```

# コンテナの起動

```
docker container run -d -p 9000:8080 example/echo:latest
```

# GETリクエストの確認

```
curl http://localhost:9000
```