# プロジェクト作成

# インストール群

- sudo apt install golang
  - goのパッケージインストール

- go version 
  - versionの確認

- パッケージの管理ツール
  - go mod init go_project
  - go mod tidy

- 環境設定
  - goenv global 1.21.0
  - goenv local 1.21.0

- ルータ
  - ルーティング
  - パスパラメータ
  - クエリパラメータ
    - 例
      - [http://localhost:8080/masaya?greeting=hello]

- gin(ホットリロード可能)
  - go get github.com/gin-gonic/gin
  - gin run src/web.go
- 環境設定パス
  - GOROOT → goがインストールされた先なのでむやみやたらと変えない方が良い。
    - [/usr/local/go/] デフォルト 
  - GOPATH → プロジェクトごとのディレクトリパスを指定する
    - [export GOPATH=go_project]
    - [source ~/.bash_profile]

# goのコマンド
- go run main.go
- go build
- go test -v
  - [-v]をつけないとテストログが出力されないので注意する

