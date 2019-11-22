# 🐌 Snail Bot

![CI Status](https://github.com/dora1998/snail-bot/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/dora1998/snail-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/dora1998/snail-bot)
[![Go Report Card](https://goreportcard.com/badge/github.com/dora1998/snail-bot)](https://goreportcard.com/report/github.com/dora1998/snail-bot)

Twitterを愛してやまない電電生に向けた、課題を管理してくれるTwitter Bot

## 実行方法

`.env.sample` を `.env` としてコピーし、適宜編集したのちに以下のコマンドを実行する。

### docker-compose

```shell script
docker-compose -f docker-compose.deps.yaml -f docker-compose.dev.yaml up -d
```

### ローカル環境での実行

direnvとdotenvを入れておく。また、Go Modulesを使用できるようにしておく。

```shell script
cp .env.sample .env.local
direnv allow
go run main.go serve
```

### テスト

```shell script
go test -v ./...
```

## 使用方法

### リプライによるコマンド実行
★：Botがフォローしているアカウントからのみ実行可能

#### タスクの追加★
```
@assignment_bot 追加 [タスク名] [期限(ex.12/31)]
```

#### タスクの削除★
```
@assignment_bot 削除 [タスク名]
```
※同名のタスクがある場合、最も最近作成したものが削除される

#### タスク一覧の取得
```
@assignment_bot 一覧
```

### 課題一覧の定期ツイート
cronなどで以下を実行すると、現在の課題一覧がツイートされる

```
docker-compose -f docker-compose.prod.yaml run snail-bot /snail-bot tweet
- or -
go run main.go tweet
```