.PHONY: all clean idl

ifeq ($(OS), Windows_NT)
    OUTPUT=greeter-server.exe
else
    OUTPUT=greeter-server
endif

all:
	go build  -v -o ./bin/${OUTPUT} cmd/main.go

idl:
	doom-protoc generate

clean:
	rm -f ./bin/${OUTPUT}
