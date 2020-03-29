
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

## docker setting
```
docker run -d \
--name corona \
-p 5888:5432 \
-e POSTGRES_USER=woshahua \
-e POSTGRES_PASSWORD=woshahua \
-e POSTGRES_DB=corona \
postgres
```

## run server 
```
go run main.go
```