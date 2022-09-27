build:
	nasm -f elf64 hello.s
	ld hello.o -o hello

run: build
	./hello 2> std_err_output.txt
	rm hello hello.o
