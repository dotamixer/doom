.PHONY: all clean idl

ifeq ($(OS), Windows_NT)
    OUTPUT=greeter-client.exe
else
    OUTPUT=greeter-client
endif

all:
	go build  -v -o ./bin/${OUTPUT} cmd/main.go

idl:
	doom-protoc generate

docker:
	docker build -t 192.168.203.41:5000/greeter-client .
	docker push 192.168.203.41:5000/greeter-client

clean:
	rm -f ./bin/${OUTPUT}