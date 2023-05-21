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