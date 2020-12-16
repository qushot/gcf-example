# gcf-example
ローカルでの実行やテスト、複数のエントリポイントが必要な場合の構成について、
公式のドキュメントを参考に実装したもの。

## Local Development

### Tools
#### Functions Framework
`go install github.com/GoogleCloudPlatform/functions-framework-go/funcframework`
[公式ドキュメント](https://cloud.google.com/functions/docs/running/function-frameworks?hl=ja#installing_dependencies)

#### Air
`go get -u github.com/cosmtrek/air`
[公式GitHub](https://github.com/cosmtrek/air)

### HTTP関数のローカル起動
```
# 端末A
# airを利用しない場合は `go run ./cmd/http/main.go` を実行する
$ air -c .http.air.toml

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.12.1 // live reload for Go apps, with Go1.14.0

mkdir $HOME/go/src/github.com/qushot/gcf-example/tmp
watching .
watching cmd
watching cmd/http
watching cmd/pubsub
!exclude tmp
watching tools
building...
running...
Serving function...

# 端末B
$ curl -X POST -d '{"name":"HTTP"}' localhost:8080
Hello, HTTP!
```

### Background関数(Pub/Sub)のローカル起動
```
# 端末A
# airを利用しない場合は `go run ./cmd/pubsub/main.go` を実行する
$ air -c .pubsub.air.toml


  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.12.1 // live reload for Go apps, with Go1.14.0

mkdir $HOME/go/src/github.com/qushot/gcf-example/tmp
watching .
watching cmd
watching cmd/http
watching cmd/pubsub
!exclude tmp
watching tools
building...
running...
Serving function...

# 端末B
$ curl -X POST -d '{"data":"'$(printf PubSub | base64)'"}' localhost:8080

# 端末A
2020/12/16 21:19:04 Hello, PubSub!
```

## Architecture
そのうち書く

## Testing
そのうち書く

## その他
- `cmd`配下のファイル、現状ではエントリポイントごとにディレクトリを切って`main.go`を作成しているが、
`cmd`直下に1つの`main.go`を用意し、`funcframework.RegisterXXXFunctionContext`を並べていったほうがいいのでは感がある。
そうした場合、`.air.toml`も1個でよくなりそう。
