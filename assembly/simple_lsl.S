.syntax divided
.thumb

.arch armv7e-m
.cpu cortex-m3

.global _start
_start:
    mov r0, #0x10
    mov r1, #1
    lsr r0, r0, #0x1
    lsl r1, r1, r0
    asr r1, r0, #32
