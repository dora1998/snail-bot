# 🐌 Snail Bot

Twitterを愛してやまない電電生に向けた、課題を管理してくれるTwitter Bot

## 実行方法

`.env.sample` を `.env` としてコピーし、適宜編集したのちに以下のコマンドを実行する。

```shell script
docker-compose -f docker-compose.deps.yaml -f docker-compose.dev.yaml up -d
```

## ローカル環境での実行

```shell script
go build ./cmd/server/main.go
```
