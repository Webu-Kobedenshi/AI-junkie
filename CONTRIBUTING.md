# Contribute

## 技術構成

- フロントエンド: React・Vite・TypeScript
- スタイリング: Tailwind CSS
- バックエンド: Go（標準の `net/http`）、PostgreSQL接続にpgx
- DB: PostgreSQL、開発用の起動にDocker Compose
- 開発ツール: mise、フロントの依存管理にpnpm、Goの依存管理にGo Modules
- 後続: Capacitorによるモバイルアプリ化を検討

## セットアップ

Linux・macOS、Windowsの場合はWSL2のLinuxシェルを前提にしています。
Git、[mise](https://mise.jdx.dev/getting-started.html)、Docker EngineとCompose（またはDocker DesktopのWSL連携）を事前に用意してください。
必要なアプリの起動やOS設定は、各自の環境に合わせて行います。
Dockerを利用できる状態で `docker info` と `docker compose version` が成功することを確認してください。
Goのrace検証にはCコンパイラも必要です。

```bash
git clone https://github.com/Webu-Kobedenshi/AI-junkie.git
cd AI-junkie

mise trust
mise install

cp .env.example .env
```

`.env` の `POSTGRES_USER` と `POSTGRES_PASSWORD` に開発用の値を手動で設定してください。既存の `.env` がある場合はコピー不要です。

```bash
mise run install
```

`mise install` は開発ツール本体、`mise run install` はプロジェクトの依存関係を取得します。
miseをシェルへactivateしていなくても、`mise run ...` で必要なバージョンを使えます。
macOS・WSL2での手順は実環境での確認が必要です。

## 設定と認証情報

環境変数はGit対象外の `.env` で設定します。他サービスの認証情報や本番の値をコピーしないでください。

| 設定                                                  | 用途                                                   |
| ----------------------------------------------------- | ------------------------------------------------------ |
| `POSTGRES_USER` / `POSTGRES_PASSWORD` / `POSTGRES_DB` | ComposeとGo APIで共有する開発用接続先                  |
| `POSTGRES_PORT`                                       | ホスト側に公開するDBポート                             |
| `PGHOST` / `PGPORT`                                   | Go APIから接続するDBのホスト・ポート                   |
| `PGSSLMODE`                                           | 開発用のローカル接続は `disable`。本番の方針は別途決定 |
| `HTTP_ADDR`                                           | APIの待受アドレス                                      |
| `API_PROXY_TARGET`                                    | Viteが `/api` を転送する接続先                         |

## コミットメッセージ

コミットメッセージは `prefix: 日本語の変更内容` の形式で記述します。

| prefix     | 用途                           | 例                                  |
| ---------- | ------------------------------ | ----------------------------------- |
| `feat`     | 機能の追加                     | `feat: 行動記録の入力画面を追加`    |
| `fix`      | 不具合の修正                   | `fix: 日付をまたぐ記録の表示を修正` |
| `docs`     | ドキュメントの変更             | `docs: セットアップ手順を更新`      |
| `refactor` | 挙動を変えないコード整理       | `refactor: DB接続処理を共通化`      |
| `test`     | テストの追加・修正             | `test: 記録時刻の検証を追加`        |
| `chore`    | 依存更新・設定・開発環境の整備 | `chore: 開発用Goのバージョンを更新` |

- prefixは英小文字とし、コロンの後に半角スペースを入れます。
- 説明には何を変更したかを具体的に書きます。

## 関連資料

- [構成と採用理由](docs/architecture.md)
- [PRD](https://docs.google.com/document/d/1d89M3ARCa6HIeVgZHD5ESq3XOtfDZrACuZMdYc7MMeo/edit)
