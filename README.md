# はじめに

書籍や動画、Qiita 記事を参考に Docker, Docker-compose を使用して Ruby on Rails の環境を構築することができたのでまとめました。

# 環境

- MacOS(Big Sur)
- Docker: 20.10.12
- Docker Compose: v2.2.3
- Ruby: 3.1.1
- Rails: 6.1.3.1

# 構築手順

## ファイル作成

```shell
# Railsアプリ用ディレクトリ作成
mkdir rails_project

# rails_projectに移動
cd rails_project

# 必要ファイルの作成
touch Dockerfile
touch docker-compose.yml
touch Gemfile
touch Gemfile.lock
```

# 設定ファイル修正

- Dockerfile

```Dockerfile
# イメージ名にRubyの実行環境のイメージを指定
FROM ruby:3.1.1

# yarnパッケージ管理ツールをインストール
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

# パッケージのリストを更新しrailsの環境構築に必要なパッケージをインストール
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs yarn

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

- docker-compose.yml

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
    command: bash -c "rm -rf tmp/pids/server.pid && bundle exec rails s -p 3000 -b '0.0.0.0'"
    # カレントディレクトリを/myappにバインドマウント
    volumes:
      - .:/myapp
    # 3000で公開して、コンテナの3000へ転送
    ports:
      - "3000:3000"
    # Webサービスを起動する前にdbサービスを起動
    depends_on:
      - db
    environment:
      WEBPACKER_DEV_SERVER_HOST: webpacker
  webpacker:
    build: .
    volumes:
      - .:/myapp
    command: ./bin/webpack-dev-server
    environment:
      WEBPACKER_DEV_SERVER_HOST: 0.0.0.0
    ports:
      - "3035:3035"
# データ永続化のためにpgdatabolのvolumeを作成し、postgresqlのデータ領域をマウント
volumes:
  pgdatavol:
```

- Gemfile

```Gemfile
source 'https://rubygems.org'
gem 'rails', '6.1.3.1'
```

- Gemfile.lock
  ※空ファイルで問題ないです。

```Gemfile.lock

```

## Rails アプリケーション作成

```zsh
# Railsアプリケーションの作成
docker-compose run web rails new . --force --database=postgresql

# イメージのビルド
docker-compose build
```

## rails プロジェクトに使用するデータベースの設定ファイルを修正

```config/database.yml
default: &default
  adapter: postgresql
  encoding: unicode
  # -------- 追加 --------
  host: db
  username: postgresql
  password: postgresql-pass
  # -------- ここまで --------
```

## データベース作成コマンド

```zsh
docker-compose run web rails db:create
```

## Scaffold にて簡易的なアプリケーション作成

```zsh
docker-compose run web bin/rails g scaffold User name:string
```

```zsh
docker-compose run web bin/rails db:migrate
```

## Rails アプリケーションの起動

```zsh
docker-compose up
# ※デタッチモード（バックグラウンド）で起動したい場合は -d オプションを追加します。
```

## アプリケーションの動作確認

下記 URL から scaffold コマンドで作成された CRUD アプリケーションが動くか確認します。
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

# Go 言語を使用してみる

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

# GET リクエストの確認

```
curl http://localhost:9000
```
