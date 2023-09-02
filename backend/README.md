# go_project

# docker操作
  - 目的:環境を汚さないことを目的とする。
    - cd /backend
    - docker build # buildの実行
    - docker-compose up -d # コンテナの立ち上げ
    - docker exec -it go_app /bin/sh # コンテナの中に入る
    - cd sample
    - go run main.go # goの実行

