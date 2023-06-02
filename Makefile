run:
	go run main.go

swager:
	swag init -g main.go --output controller/docs

all: swager run


