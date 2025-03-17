section .text
global binary_convert
binary_convert:
	xor			eax, eax			; eax will hold our return value

.loop:
	movzx 		ebx, byte [rdi]		; get first byte from input
	test		ebx, ebx			; compare byte to 0 does a bitwise &
	jz			fi
	shl			eax, 1				; equivalent to return * 2
	sub			ebx, 48				; change ascii to int, 48 ascii = 0
	add			eax, ebx
	inc			rdi
	jmp			.loop					

fi:
	ret