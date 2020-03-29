
# Setup
## install golang
1.12.7 をローカルインストールする

```
brew install golang
```

## Git clone
`$GOPATH`を設定してgit clone

```
$ git clone https://github.com/woshahua/corona_server.git $GOPATH/src/github.com/woshahua/corona_server
```

## Run
### 依存ライブラリをダウンロード
```
go mod download
```
