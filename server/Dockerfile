FROM golang:1.20 as builder

# airのインストール
RUN go install github.com/cosmtrek/air@latest

# 環境変数の設定
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app

# ワークディレクトリの設定
WORKDIR ${ROOTPATH}

# Goモジュールファイルのコピー
COPY go.mod go.sum .air.toml ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# ポートの公開
EXPOSE 80

# airを使ってアプリケーションを起動
CMD ["air", "-c", ".air.toml"]
