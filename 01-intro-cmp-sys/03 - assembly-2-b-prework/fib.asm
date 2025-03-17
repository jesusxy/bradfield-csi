section .text
global fib
fib:
	mov		rax, rdi
	cmp 	rdi, 1
	jle		.base

	push	rdi
	dec		rdi
	call	fib

	pop		rdi
	push	rax
	sub		rdi, 2
	call	fib
	pop		rdi
	add		rax, rdi
	ret

.base:
	ret