.PHONY: all clean idl uidl


ifeq ($(OS), Windows_NT)
    OUTPUT=doom-protoc.exe
else
    OUTPUT=doom-protoc
endif

all: clean
	go build -o ./example/${OUTPUT} ./cmd/main.go

clean:
	rm -f example/${OUTPUT}