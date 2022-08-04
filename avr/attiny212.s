; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATtiny212.atdf, see http://packs.download.atmel.com/

; This is the default handler for interrupts, if triggered but not defined.
; Sleep inside so that an accidentally triggered interrupt won't drain the
; battery of a battery-powered device.
.section .text.__vector_default
.global  __vector_default
__vector_default:
    sleep
    rjmp __vector_default

; Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, __vector_default
.endm

; The interrupt vector of this device. Must be placed at address 0 by the linker.
.section .vectors
.global  __vectors
    rjmp __vector_default
    rjmp __vector_NMI
    rjmp __vector_VLM
    rjmp __vector_PORT
    rjmp __vector_default
    rjmp __vector_default
    rjmp __vector_CNT
    rjmp __vector_PIT
    rjmp __vector_LUNF
    rjmp __vector_HUNF
    rjmp __vector_LCMP0
    rjmp __vector_CMP1
    rjmp __vector_CMP2
    rjmp __vector_INT
    rjmp __vector_OVF
    rjmp __vector_TRIG
    rjmp __vector_AC
    rjmp __vector_RESRDY
    rjmp __vector_WCOMP
    rjmp __vector_TWIS
    rjmp __vector_TWIM
    rjmp __vector_INT
    rjmp __vector_RXC
    rjmp __vector_DRE
    rjmp __vector_TXC
    rjmp __vector_EE

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_NMI
    IRQ __vector_VLM
    IRQ __vector_PORT
    IRQ __vector_CNT
    IRQ __vector_PIT
    IRQ __vector_LUNF
    IRQ __vector_OVF
    IRQ __vector_HUNF
    IRQ __vector_LCMP0
    IRQ __vector_CMP0
    IRQ __vector_CMP1
    IRQ __vector_LCMP1
    IRQ __vector_CMP2
    IRQ __vector_LCMP2
    IRQ __vector_INT
    IRQ __vector_OVF
    IRQ __vector_TRIG
    IRQ __vector_AC
    IRQ __vector_RESRDY
    IRQ __vector_WCOMP
    IRQ __vector_TWIS
    IRQ __vector_TWIM
    IRQ __vector_INT
    IRQ __vector_RXC
    IRQ __vector_DRE
    IRQ __vector_TXC
    IRQ __vector_EE
