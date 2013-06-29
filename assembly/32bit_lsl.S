.syntax unified
.thumb

.arch armv7e-m
.cpu cortex-m3

.global _start
.thumb_func
_start:
    mov r0, #0x10
    lsr r0, r0, #0x1
    lsl r1, r0, r0
