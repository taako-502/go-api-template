# go-api-template

[Huma](https://huma.rocks/)と`net/http`を利用した、Go製APIサーバーのテンプレートです。
`nakith-api`の構成を基に、アプリケーション固有の依存を取り除いています。

## 主な機能

- Humaによるルーティング、バリデーション、Swagger UI、OpenAPI生成
- Domain / Usecase / Infrastructure / Presentationのレイヤー分離
- 差し替え可能なRepositoryと、タスクCRUDの実装例
- `slog`による構造化ログ
- Request ID、panic recovery、CORS
- Liveness / Readinessエンドポイント
- Graceful shutdown
- pprof用の独立したデバッグサーバー
- Race detector、`go vet`、脆弱性検査を実行するCI
- Dockerfile、Dependabot設定

## セットアップ

```sh
cp .env.example .env
set -a; . ./.env; set +a
go mod tidy
make run
```

- API: `http://localhost:8080`
- Swagger UI: `http://localhost:8080/docs`
- OpenAPI: `http://localhost:8080/openapi.json`
- Liveness: `http://localhost:8080/health/live`
- Readiness: `http://localhost:8080/health/ready`
- pprof: `http://127.0.0.1:6060/debug/pprof/`

## コマンド

```sh
make run
make test
make test-race
make lint
make vuln
make openapi
make build
```

## API実行例

```sh
curl -X POST http://localhost:8080/v1/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"サンプル機能を置き換える"}'
```

設計方針とディレクトリ構成は[ARCHITECTURE.md](ARCHITECTURE.md)を参照してください。
