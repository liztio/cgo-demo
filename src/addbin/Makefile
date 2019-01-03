all: addbin

CFLAGS=-Wall -Werror

addbin: addbin.c
	$(CC) $(CFLAGS) -L"../addlib" -I"../addlib"  -o $@ $^ -laddlib

run: addbin
	LD_LIBRARY_PATH="../addlib:$LD_LIBRARY_PATH" ./addbin 10

clean:
	rm addbin
