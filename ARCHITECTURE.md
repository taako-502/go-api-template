# アーキテクチャー

本テンプレートは、DDDとオニオンアーキテクチャーの考え方を簡略化して採用する。
依存方向は外側から内側へ向け、ドメイン層をHTTPやデータストアの実装から分離する。

## ディレクトリ構成

```text
.
├── cmd
│   ├── api
│   └── openapi
├── docs
│   └── api
├── internal
│   ├── app
│   ├── config
│   ├── domain
│   │   └── task
│   ├── infrastructure
│   │   └── memory
│   ├── platform
│   │   └── server
│   ├── presentation
│   │   └── httpapi
│   └── usecase
│       ├── page
│       └── task_usecase
├── Dockerfile
├── Makefile
└── go.mod
```

## レイヤーの責務

- `internal/domain`
  - エンティティ、ドメインエラー、Repositoryインターフェースを定義する。
  - Huma、`net/http`、データストアなどの外部実装に依存しない。
- `internal/usecase`
  - アプリケーション固有のユースケースを実装する。
  - 原則としてドメイン層のみに依存する。
- `internal/infrastructure`
  - Repositoryや外部サービスの具体的な実装を配置する。
  - サンプルではインメモリRepositoryを実装している。
- `internal/presentation`
  - Humaと`net/http`を利用したルーティング、入出力変換、Middlewareを担当する。
- `internal/app`
  - Repository、Usecase、Handlerを組み立てるComposition Root。
- `internal/platform`
  - HTTPサーバーなど、アプリケーション全体で利用する実行基盤を配置する。
- `cmd`
  - 実行可能ファイルのエントリーポイントを配置する。

## 依存方向

```text
cmd
 ├── app ── presentation
 │    └── infrastructure
 └── platform

presentation ── usecase ── domain
infrastructure ─────────── domain
```

PostgreSQL、認証、Redis、クラウドSDKなどは、必要になった時点でInfrastructure層へ追加する。
データストアのモデルをAPIレスポンスとして直接公開しない。
