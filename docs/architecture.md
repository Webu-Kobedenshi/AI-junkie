# システム構成

行動記録のモックで、入力負担と記録の続けやすさを検証します。収集したデータは後のサービスでも継続利用します。

## 技術構成

| 役割     | 技術                                                      |
| -------- | --------------------------------------------------------- |
| 画面     | React・Vite・TypeScript                                   |
| スタイル | Tailwind CSS（画面はJSX、テーマ・共通設定は `style.css`） |
| API      | Go標準の `net/http`、DB接続はpgx                          |
| DB       | PostgreSQL・Docker Compose                                |
| 開発環境 | miseでツールを固定、依存管理はpnpm・Go Modules            |
| モバイル | Capacitorを検討（後続）                                   |

## 接続

開発中：

```text
ブラウザ → Vite（5173） → /api → Go（8080） → PostgreSQL（55432）
```

本番の予定：

```text
ブラウザ → Go（ビルド済み画面＋API） → PostgreSQL
```

本番はアプリとDBをComposeで起動する方針です。現在の `compose.yaml` は開発用DBのみ。本番配置・公開経路・バックアップは未整備です。

開発手順は [CONTRIBUTING.md](../CONTRIBUTING.md) を参照してください。
