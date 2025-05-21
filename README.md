# Go 學習計畫：任務管理 API（Task Manager）

本專案為期 30 天，採取兩天為一單元的高強度學習節奏，透過實作一個「任務管理系統 API」逐步掌握 Go 語言核心、框架應用與實戰技能。

> 📌 學習風格：先見林再見樹（先整體理解架構，再逐步深入各層細節）

---

## 🎯 最終目標

建立一個可用的 RESTful 任務管理 API，具備：

- 使用者註冊、登入（JWT）
- 任務 CRUD、過濾查詢
- 資料庫持久化（GORM）
- 中介層（Middleware）
- 統一回傳格式
- Docker 部署
- ✅ 涵蓋 defer、goroutine、channel、map vs slice、for-loop、函式封裝、struct 應用等 Go 核心

---

## 🗂 學習地圖（30 天進度表）

### 🧱 Week 1：Go 基礎 + 專案初始化

| 單元 | 主題 | 任務內容 | 重點語法練習 |
|------|------|----------|------------------------|
| Day 1-2 | Go 基礎語法 | 變數、struct、interface、error handling、函式寫法 | ✅ `struct`, `func`, `defer`, `for-loop` |
| Day 3-4 | 專案初始化 | `go mod`、目錄規劃、Clean Architecture 導入 | ✅ `slice` vs `map` 比較 |
| Day 5-6 | 模組封裝 | 抽出 handler / service / model 模組化 | ✅ 函式封裝（高階函式）

---

### 🌐 Week 2：Web API 開發與資料庫整合

| 單元 | 主題 | 任務內容 | 重點語法練習 |
|------|------|----------|------------------------|
| Day 7-8 | Gin 框架 | 建立 `/ping`、`/tasks` 初始 API | ✅ `for-loop` 用於列出任務 |
| Day 9-10 | GORM 整合 | 建立 Task model + MySQL CRUD | ✅ 使用 `map` 儲存錯誤訊息對應 |
| Day 11-12 | Migration & Seeder | 撰寫 `AutoMigrate()` 與假資料 | ✅ 用 `slice` 建立假任務清單

---

### 🔐 Week 3：認證、通道與測試

| 單元 | 主題 | 任務內容 | 重點語法練習 |
|------|------|----------|------------------------|
| Day 13-14 | Middleware & Logger | 撰寫 request logger、middleware 使用 | ✅ `defer` 撰寫計時器 |
| Day 15-16 | JWT 登入註冊 | 使用者登入 + token middleware 驗證 | ✅ 封裝 token 函式與結構體解析 |
| Day 17-18 | 單元測試 + 並行處理 | 使用 testify 撰寫 handler 測試 | ✅ `goroutine` 處理多任務通知

---

### 🚀 Week 4：通道與實務部署

| 單元 | 主題 | 任務內容 | 重點語法練習 |
|------|------|----------|------------------------|
| Day 19-20 | Docker 化 | 建立 Dockerfile, docker-compose | - |
| Day 21-22 | goroutine 實戰 | 排程逾期任務提醒（非同步） | ✅ `goroutine`, `channel`, `select` |
| Day 23-24 | 串接外部 API | 並行取得天氣/任務通知內容 | ✅ `goroutine` + `channel` 聚合結果 |
| Day 25-26 | 性能優化 | 加入 context timeout + pprof 分析 | ✅ 用 `defer cancel()` 控制 context |
| Day 27-28 | 整合文件 | 撰寫 API 文件與 swagger 或 postman 範例 | - |
| Day 29-30 | 上線部署 | 推上 Render / Fly.io + README 完善 | - |

---

## 📁 專案結構範例（參考）

```
task-manager-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/      # API handler
│   ├── service/      # 業務邏輯
│   ├── model/        # GORM model
│   ├── middleware/   # JWT / Logger
│   ├── util/         # 通道、goroutine 工具
│   └── repository/   # DB 抽象介面
├── database/
│   ├── migrate.go
│   └── seed.go
├── go.mod
└── Dockerfile
```

---

## 🧪 預計涵蓋技能

- Go 語法與型別系統
- Gin 框架實戰
- GORM 與資料持久化設計
- 中介層開發（middleware）
- JWT 認證
- 單元測試與模擬資料
- `defer`, `goroutine`, `channel`, `map`, `slice`, `for-loop` 熟練應用
- Docker 化與部署流程
- Clean Architecture 應用

---
