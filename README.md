# 🐌 Snail Bot

![CI Status](https://github.com/dora1998/snail-bot/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/dora1998/snail-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/dora1998/snail-bot)

Twitterを愛してやまない電電生に向けた、課題を管理してくれるTwitter Bot

## 実行方法

`.env.sample` を `.env` としてコピーし、適宜編集したのちに以下のコマンドを実行する。

```shell script
docker-compose -f docker-compose.deps.yaml -f docker-compose.dev.yaml up -d
```

## ローカル環境での実行

```shell script
go run main.go serve
```
