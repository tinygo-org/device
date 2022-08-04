// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from LPC11Axxv0.6.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// LPC11Axx
//


.syntax unified

// This is the default handler for interrupts, if triggered but not defined.
.section .text.Default_Handler
.global  Default_Handler
.type    Default_Handler, %function
Default_Handler:
    wfe
    b    Default_Handler
.size Default_Handler, .-Default_Handler

// Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, Default_Handler
.endm

// Must set the "a" flag on the section:
// https://svnweb.freebsd.org/base/stable/11/sys/arm/arm/locore-v4.S?r1=321049&r2=321048&pathrev=321049
// https://sourceware.org/binutils/docs/as/Section.html#ELF-Version
.section .isr_vector, "a", %progbits
.global  __isr_vector
__isr_vector:
    // Interrupt vector as defined by Cortex-M, starting with the stack top.
    // On reset, SP is initialized with *0x0 and PC is loaded with *0x4, loading
    // _stack_top and Reset_Handler.
    .long _stack_top
    .long Reset_Handler
    .long NMI_Handler
    .long HardFault_Handler
    .long MemoryManagement_Handler
    .long BusFault_Handler
    .long UsageFault_Handler
    .long 0
    .long 0
    .long 0
    .long 0
    .long SVC_Handler
    .long DebugMon_Handler
    .long 0
    .long PendSV_Handler
    .long SysTick_Handler

    // Extra interrupts for peripherals defined by the hardware vendor.
    .long PIN_INT0_IRQHandler
    .long PIN_INT1_IRQHandler
    .long PIN_INT2_IRQHandler
    .long PIN_INT3_IRQHandler
    .long PIN_INT4_IRQHandler
    .long PIN_INT5_IRQHandler
    .long PIN_INT6_IRQHandler
    .long PIN_INT7_IRQHandler
    .long GINT0_IRQHandler
    .long GINT1_IRQHandler
    .long CMP_IRQHandler
    .long DAC_IRQHandler
    .long RESERVED0_IRQHandler
    .long RESERVED1_IRQHandler
    .long SSP1_IRQHandler
    .long I2C_IRQHandler
    .long CT16B0_IRQHandler
    .long CT16B1_IRQHandler
    .long CT32B0_IRQHandler
    .long CT32B1_IRQHandler
    .long SSP0_IRQHandler
    .long USART_IRQHandler
    .long RESERVED2_IRQHandler
    .long RESERVED3_IRQHandler
    .long ADC_IRQHandler
    .long WDT_IRQHandler
    .long BOD_IRQHandler
    .long FMC_IRQHandler
    .long RESERVED4_IRQHandler
    .long RESERVED5_IRQHandler
    .long RESERVED6_IRQHandler
    .long RESERVED7_IRQHandler

    // Define default implementations for interrupts, redirecting to
    // Default_Handler when not implemented.
    IRQ NMI_Handler
    IRQ HardFault_Handler
    IRQ MemoryManagement_Handler
    IRQ BusFault_Handler
    IRQ UsageFault_Handler
    IRQ SVC_Handler
    IRQ DebugMon_Handler
    IRQ PendSV_Handler
    IRQ SysTick_Handler
    IRQ PIN_INT0_IRQHandler
    IRQ PIN_INT1_IRQHandler
    IRQ PIN_INT2_IRQHandler
    IRQ PIN_INT3_IRQHandler
    IRQ PIN_INT4_IRQHandler
    IRQ PIN_INT5_IRQHandler
    IRQ PIN_INT6_IRQHandler
    IRQ PIN_INT7_IRQHandler
    IRQ GINT0_IRQHandler
    IRQ GINT1_IRQHandler
    IRQ CMP_IRQHandler
    IRQ DAC_IRQHandler
    IRQ RESERVED0_IRQHandler
    IRQ RESERVED1_IRQHandler
    IRQ SSP1_IRQHandler
    IRQ I2C_IRQHandler
    IRQ CT16B0_IRQHandler
    IRQ CT16B1_IRQHandler
    IRQ CT32B0_IRQHandler
    IRQ CT32B1_IRQHandler
    IRQ SSP0_IRQHandler
    IRQ USART_IRQHandler
    IRQ RESERVED2_IRQHandler
    IRQ RESERVED3_IRQHandler
    IRQ ADC_IRQHandler
    IRQ WDT_IRQHandler
    IRQ BOD_IRQHandler
    IRQ FMC_IRQHandler
    IRQ RESERVED4_IRQHandler
    IRQ RESERVED5_IRQHandler
    IRQ RESERVED6_IRQHandler
    IRQ RESERVED7_IRQHandler
