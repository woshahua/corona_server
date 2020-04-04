
# Setup
## install golang
golang をローカルインストールする

```
brew install golang
```

## Git clone
`$GOPATH`を設定してgit clone

```
$ git clone https://github.com/woshahua/corona_server.git $GOPATH/src/github.com/woshahua/corona_server
```

## Run
## 依存ライブラリをダウンロード
```
go mod download
```

## create local db
```
初回は先にcoronaというdbを作る必要がある
db/配下のsqlを使って必要なtableを作成
```
## localで起動
```
ENV=development go run main.go
```

## 本番deploy
```
gcloud app deploy
```
## 本番cron task deploy
```
gcloud app deploy cron.yaml
```
