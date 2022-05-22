# gtw
Go用のtwitter APIです。Twitter API v2で動作します。まだalpha版なので機能も少ないし、bugも少なからずあると思います。

## 準備
このライブラリーではユーザーログイン状態での動作をOAuth 1.0aで、ユーザーログインの要らないpublicなアクセスをOAuth 2.0で行っているので、認証に次の情報が必要です。
- API KeyとSecret
- Access TokenとSecret
- Beare Token
これらは[TwitterのDeveloper Portal](https://developer.twitter.com/en/apps/)から取得できます。

## インストール
普通のGoライブラリーと同様にインストールできます。具体的には下記でインストールできます。

```console
go get github.com/Michi-gi/gtw
```
## 機能
現時点で次のAPIが実装されています。
- Tweets
  - Tweets lookup
  - Manage Tweets
  - Timeline(Reverse chronological home timelineは未対応)
  - Search Tweets(Recent searchのみ。Full-archive searchは未検証)

## コード例
例えば、tweetを取得する場合は次のようにする。
```
package main
import (
    "fmt"

    "github.com/Michi-gi/gtw"
)

func main() {
    client := gtw.NewClient(
        "YOUR_API_KEY",
        "YOUR_API_KEY_SECRET",k
        "YOUR_ACCESS_TOKEN",
        "YOUR_ACCESS_TOKEN_SECRET",
        "YOUR_BEARER_TOKEN")
    
    tweet, err := client.Tweets.Get(
        "TARGET_TWEET_ID",
        []string{"attachments.media_keys", "author_id"},
        false)
    if err != nil {
        fmt.Println("Error!")
    } else {
        fmt.Println(tweet)
    }
}
```

## 利用ライブラリー
このライブラリーはOAuth 1.0aに https://github.com/dghubble/oauth1 を使っています。