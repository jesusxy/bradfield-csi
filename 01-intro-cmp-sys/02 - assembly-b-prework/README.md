## NASM

Nasm programs are divided into `sections`

- Data section
- Text section

### Data

---

This sections is for declaring `constants`

### Text

---

The text section is where the actual `code` lives \
`global _start` is where program execution **begins** \

### .bss section

---

This section is for _writeable_ data

### NASM Instruction set

---

Nasm instructions consist of two parts

- Name of the instruction
- Operands of command

`MOV COUNT, 48 ; Put value 48 in the COUNT variable`

> je : jump equal \
> jne : jump if not equal

Instructions can have on of the following operands

- register

  - internal variables
  - ineger registers (r8-r15)
  - general purpose (rax, rbx, rcx, rdx)
  - index registers (rsi, rbi, rsp, rbp)
    - rbp: base pointer
    - rsp: stack pointer

- memory \
  `effective addressing` \
  Using [ ] means we are reading | writing from this address / memory location

```
mov [ my_var ] , ecx
mov ecx, [ my_var ]
```

If we wish to get the address of a variable, we would do the following

```
mov eax , my_var ; sets eax to the address of my_var
lea eax, [ my_var ] ; read and moves the address of my_var. (L) oad (E) ffective (A) ddress
```

- immediate values
