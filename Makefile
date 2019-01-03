.PHONY: all clean run

all:
	make -C src/addlib
	make -C src/addbin

clean:
	make -C src/addlib clean
	make -C src/addbin clean

run: all
	make -C src/addbin run

