.PHONY: all clean

CFLAGS=-Wall -Werror -fpic
LDFLAGS=-shared

all: libaddlib.so

addlib.o: addlib.c addlib.h
	$(CC) $(CFLAGS) -c $^

libaddlib.so: addlib.o
	$(CC) -o $@ $^ $(LDFLAGS)

clean:
	rm -f libaddlib.so addlib.o
