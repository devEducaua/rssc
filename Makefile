TARGET ?= rssc
SRCS := $(wildcard */*.go)

all: $(TARGET)

$(TARGET): $(SRCS)
	go build -o $(TARGET) .

clean:
	rm $(TARGET)
