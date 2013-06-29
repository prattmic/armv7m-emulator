.arch armv7e-m
.cpu cortex-m4
.thumb

.global _start
_start:
    mov r0, #0x10
    lsr r0, r0, #0x1
