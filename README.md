# golang_api_sample
Simple SNS api made by Golang.

# API docs

## user

[POST] /user

ユーザーを作成する。
params: name, password

[GET] /:user_id/followers

フォロロワー一覧を取得する

[GET] /:user_id/following

フォロー中のユーザー一覧を取得する

いいねした日記一覧を取得する

## login

[POST] /login

user_idとpasswordを受け取り、既に存在するユーザーであればアクセストークンを返す

## follower_request

[GET] /follower_requests

自分のフォローリクエストの一覧を取得する。他人のフォローリクエストの一覧は取得することができない

[POST] /:user_id/follower_requests

フォローリクエストを送信する。
すでにフォロー済みの場合や、自分のidを指定した場合は400 BadRequestを返す

[POST] /:follower_requests_id/allow

フォローリクエストを許可する


## diary

[GET] /diaries?page=1

フォローしているユーザーの日記一覧を取得する。
各日記にいいねしているかの情報も取得

[GET] /:user_id:/diaries?page=1

ユーザーごとの日記一覧を取得する


[POST] /diaries

日記を投稿する
params: content

## like

[POST] diaries/:dialy_id/like

日記にいいねを付ける
自分の日記の場合は400 BadRequestを返す


# memo
webアプリ等でユーザーからURLが見えるなら日記の一覧取得などは`/home`とか`/`にするが、ネイティブアプリではURLが見れないので開発者が分かりやすいようにリソース名を明示する

`/users/:id`のようなルーティングをする場合、net/httpだけでは難しいので独自の実装を加える。
> https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
