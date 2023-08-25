.PHONY: run-main run-emails

run-main:
	go run cmd/events/main.go

run-emails:
	go run cmd/emails/main.go