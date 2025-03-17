section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex
	imul	rdx, rcx 		
	lea		rax, [rdx * 4]		; gives us where the starting point of target row is
	lea		rax, [rax + r8 * 4]	; give us starting point of column
	mov		rax, [rdi + rax]	; moves the value at given address
	ret

; first index and load the address of the correct row into a register
; index and load the address of the correct column
; move the value of that column into return register rax

; third test
; int index(int *matrix, int rows, int cols, int rindex, int cindex);
; TEST_ASSERT_EQUAL(6, index((int *)matrix, 2, 3, 1, 2))
; int matrix[2][3] = {{1, 2, 3}, {4, 5, 6}};
;
; an array is stored in memory contigously
; an array of arrays (2D) should be conitgous as well
; 
; memory layout of the matrix in test2

;  [ 1 2 3   4 5 6  ]	-> 24 bytes in size
;   <-m[0]-><-m[1]->

; cols * rindex : 3 * 1 = 3
; add 3 + r8 = 3 + 2 = 5
; 4 * 5 = 20

; cols * rindex : 3 * 1 = 3
; 3 * 4 = 12, starting point of row : 12th byte
; 12 + 2 * 4 = 20 starting point of value: 20th byte