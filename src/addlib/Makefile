CFLAGS=-Wall -Werror -fpic

all: libaddlib.so

addlib.o: addlib.c addlib.h
	$(CC) $(CFLAGS) -c $^

libaddlib.so: addlib.o
	$(CC) -shared -o $@ $^

clean:
	rm libaddlib.so addlib.o
