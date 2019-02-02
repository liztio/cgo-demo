.PHONY: all clean

all: addbin

CFLAGS=-Wall -Werror
LDFLAGS=-L"../addlib"

addbin: addbin.c
  # -L"../addlib"
	$(CC) $(CFLAGS) -I"../addlib"  -o $@ $^ -laddlib $(LDFLAGS)

run: addbin
	LD_LIBRARY_PATH="../addlib:$LD_LIBRARY_PATH" ./addbin 10

clean:
	rm -f addbin
