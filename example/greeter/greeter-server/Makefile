.PHONY: all clean idl

ifeq ($(OS), Windows_NT)
    OUTPUT=greeter-server.exe
else
    OUTPUT=greeter-server
endif

all:
	CGO_ENABLED=0 go build  -v -o ./bin/${OUTPUT} cmd/main.go

idl:
	doom-protoc generate

docker:
	docker build -t 192.168.203.41:5000/greeter-server .
	docker push 192.168.203.41:5000/greeter-server

clean:
	rm -f ./bin/${OUTPUT}
