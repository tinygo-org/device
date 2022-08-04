; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATmega32C1.atdf, see http://packs.download.atmel.com/

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
    jmp __vector_RESET
    jmp __vector_ANACOMP0
    jmp __vector_ANACOMP1
    jmp __vector_ANACOMP2
    jmp __vector_ANACOMP3
    jmp __vector_PSC_FAULT
    jmp __vector_PSC_EC
    jmp __vector_INT0
    jmp __vector_INT1
    jmp __vector_INT2
    jmp __vector_INT3
    jmp __vector_TIMER1_CAPT
    jmp __vector_TIMER1_COMPA
    jmp __vector_TIMER1_COMPB
    jmp __vector_TIMER1_OVF
    jmp __vector_TIMER0_COMPA
    jmp __vector_TIMER0_COMPB
    jmp __vector_TIMER0_OVF
    jmp __vector_CAN_INT
    jmp __vector_CAN_TOVF
    jmp __vector_LIN_TC
    jmp __vector_LIN_ERR
    jmp __vector_PCINT0
    jmp __vector_PCINT1
    jmp __vector_PCINT2
    jmp __vector_PCINT3
    jmp __vector_SPI_STC
    jmp __vector_ADC
    jmp __vector_WDT
    jmp __vector_EE_READY
    jmp __vector_SPM_READY

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_ANACOMP0
    IRQ __vector_ANACOMP1
    IRQ __vector_ANACOMP2
    IRQ __vector_ANACOMP3
    IRQ __vector_PSC_FAULT
    IRQ __vector_PSC_EC
    IRQ __vector_INT0
    IRQ __vector_INT1
    IRQ __vector_INT2
    IRQ __vector_INT3
    IRQ __vector_TIMER1_CAPT
    IRQ __vector_TIMER1_COMPA
    IRQ __vector_TIMER1_COMPB
    IRQ __vector_TIMER1_OVF
    IRQ __vector_TIMER0_COMPA
    IRQ __vector_TIMER0_COMPB
    IRQ __vector_TIMER0_OVF
    IRQ __vector_CAN_INT
    IRQ __vector_CAN_TOVF
    IRQ __vector_LIN_TC
    IRQ __vector_LIN_ERR
    IRQ __vector_PCINT0
    IRQ __vector_PCINT1
    IRQ __vector_PCINT2
    IRQ __vector_PCINT3
    IRQ __vector_SPI_STC
    IRQ __vector_ADC
    IRQ __vector_WDT
    IRQ __vector_EE_READY
    IRQ __vector_SPM_READY
