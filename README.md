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

### ローカルでの起動
#### 準備
動作させたい関数をFunctions Frameworkを利用して`cmd/main.go`内で登録する。

```
// HTTPEntryPoint関数を登録する例
if err := funcframework.RegisterHTTPFunctionContext(ctx, "/http-entry-point", gcfexample.HTTPEntryPoint); err != nil {
	log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
}
```

#### 起動
```
# 端末A
# airを利用しない場合は `go run ./cmd/main.go` を実行する
$ air -c .air.toml

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.12.1 // live reload for Go apps, with Go1.14.0

mkdir /Users/moto/go/src/github.com/qushot/gcf-example/tmp
watching .
watching cmd
!exclude tmp
watching tools
building...
running...
Serving function...

# 端末B: HTTP関数にリクエスト
$ curl -X POST -d '{"name":"HTTP"}' localhost:8080/http-entry-point
Hello, HTTP!

# 端末B: Background関数(Pub/Sub)にリクエスト
$ curl -X POST -d '{"data":"'$(printf PubSub | base64)'"}' localhost:8080/background-pubsub-entry-point

# 端末A: Background関数(Pub/Sub)内で出力されたログ
2020/12/16 21:19:04 Hello, PubSub!
```

## Architecture
そのうち書く

## Testing
そのうち書く
