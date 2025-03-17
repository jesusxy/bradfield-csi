			global sum_to_n
		
			section .text
sum_to_n:	
		xor		rax, rax	
		jmp		loop		
		ret
			

loop:
		add		rax, rdi 	; adds n to rax reg
		dec		rdi		 	; decrement rdi buy 1
		jg		loop	 	; is n > 0, if yes continue loop
		ret


; sum_to_n O(1)

; sum_to_n:
; 								; rax holds return value, rdi holds input
; 			mov		rax, rdi	; mov n (rdi) to rax (to be used later)
; 			imul	rdi, rdi	; multiply n * n
; 			add		rax, rdi	; n + (n * n) -> 5 + (5 * 5) = 25
; 			shr		rax, 1		; shift right once | rax / 2
; 			ret
