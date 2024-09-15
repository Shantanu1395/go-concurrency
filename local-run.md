```bash
go test ./internal/scraper/... -v
go test -bench=. ./internal/scraper/...
go test -run TestMemoryLeak ./internal/scraper/...
go mod tidy
go run main.go
```