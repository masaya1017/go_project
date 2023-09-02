# go_project

# docker操作
  - 目的:環境を汚さないことを目的とする。
    - cd /backend
    - docker build # buildの実行
    - docker-compose up -d # コンテナの立ち上げ
    - docker exec -it go_app /bin/sh # コンテナの中に入る
    - cd sample
    

# go実行
- プログラム単位の実行
  - go run main.go # goの実行
- テストの実行
  - go mod init local.package/main # go moduleの初期化
  - go test # テストの実行
- 
　- go mod tidy
- キャッシュ削除
  - go clean -testcache

