obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

.PHONY: obu

reciever:
	@go build -o bin/obu_data_reciever ./obu_data_reciever
	@./bin/obu_data_reciever

.PHONY: reciever
