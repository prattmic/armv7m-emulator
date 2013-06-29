.syntax divided
.thumb

.arch armv7e-m
.cpu cortex-m3

.global _start
_start:
    mov r0, #0x10
    lsr r0, r0, #0x1
