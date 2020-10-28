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

clean:
	rm -f ./bin/${OUTPUT}