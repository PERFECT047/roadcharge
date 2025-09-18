obu:
	@go build -o bin/obu ./obu
	@./bin/obu

reciever:
	@go build -o bin/obu_data_reciever ./obu_data_reciever
	@./bin/obu_data_reciever

dist_calc:
	@go build -o bin/dist_calculator ./dist_calculator
	@./bin/dist_calculator

.PHONY: obu
