# 各タスクの実行用
.PHONY: bingo fizzbuzz tag

bingo:
	go run cmd/bingo/main.go
fizzbuzz:
	go run cmd/fizzbuzz/main.go
tag:
	go run cmd/tag/main.go
