// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from LPC13xx_svd_v1.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// LPC13xx
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
    .long PIO0_0_IRQHandler
    .long PIO0_1_IRQHandler
    .long PIO0_2_IRQHandler
    .long PIO0_3_IRQHandler
    .long PIO0_4_IRQHandler
    .long PIO0_5_IRQHandler
    .long PIO0_6_IRQHandler
    .long PIO0_7_IRQHandler
    .long PIO0_8_IRQHandler
    .long PIO0_9_IRQHandler
    .long PIO0_10_IRQHandler
    .long PIO0_11_IRQHandler
    .long PIO1_0_IRQHandler
    .long PIO1_1_IRQHandler
    .long PIO1_2_IRQHandler
    .long PIO1_3_IRQHandler
    .long PIO1_4_IRQHandler
    .long PIO1_5_IRQHandler
    .long PIO1_6_IRQHandler
    .long PIO1_7_IRQHandler
    .long PIO1_8_IRQHandler
    .long PIO1_9_IRQHandler
    .long PIO1_10_IRQHandler
    .long PIO1_11_IRQHandler
    .long PIO2_0_IRQHandler
    .long PIO2_1_IRQHandler
    .long PIO2_2_IRQHandler
    .long PIO2_3_IRQHandler
    .long PIO2_4_IRQHandler
    .long PIO2_5_IRQHandler
    .long PIO2_6_IRQHandler
    .long PIO2_7_IRQHandler
    .long PIO2_8_IRQHandler
    .long PIO2_9_IRQHandler
    .long PIO2_10_IRQHandler
    .long PIO2_11_IRQHandler
    .long PIO3_0_IRQHandler
    .long PIO3_1_IRQHandler
    .long PIO3_2_IRQHandler
    .long PIO3_3_IRQHandler
    .long I2C0_IRQHandler
    .long CT16B0_IRQHandler
    .long CT16B1_IRQHandler
    .long CT32B0_IRQHandler
    .long CT32B1_IRQHandler
    .long SSP0_IRQHandler
    .long UART_IRQHandler
    .long USBIRQ_IRQHandler
    .long USBFIQ_IRQHandler
    .long ADC_IRQHandler
    .long WDT_IRQHandler
    .long BOD_IRQHandler
    .long 0
    .long PIO_3_IRQHandler
    .long PIO_2_IRQHandler
    .long PIO_1_IRQHandler
    .long PIO_0_IRQHandler
    .long SSP1_IRQHandler

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
    IRQ PIO0_0_IRQHandler
    IRQ PIO0_1_IRQHandler
    IRQ PIO0_2_IRQHandler
    IRQ PIO0_3_IRQHandler
    IRQ PIO0_4_IRQHandler
    IRQ PIO0_5_IRQHandler
    IRQ PIO0_6_IRQHandler
    IRQ PIO0_7_IRQHandler
    IRQ PIO0_8_IRQHandler
    IRQ PIO0_9_IRQHandler
    IRQ PIO0_10_IRQHandler
    IRQ PIO0_11_IRQHandler
    IRQ PIO1_0_IRQHandler
    IRQ PIO1_1_IRQHandler
    IRQ PIO1_2_IRQHandler
    IRQ PIO1_3_IRQHandler
    IRQ PIO1_4_IRQHandler
    IRQ PIO1_5_IRQHandler
    IRQ PIO1_6_IRQHandler
    IRQ PIO1_7_IRQHandler
    IRQ PIO1_8_IRQHandler
    IRQ PIO1_9_IRQHandler
    IRQ PIO1_10_IRQHandler
    IRQ PIO1_11_IRQHandler
    IRQ PIO2_0_IRQHandler
    IRQ PIO2_1_IRQHandler
    IRQ PIO2_2_IRQHandler
    IRQ PIO2_3_IRQHandler
    IRQ PIO2_4_IRQHandler
    IRQ PIO2_5_IRQHandler
    IRQ PIO2_6_IRQHandler
    IRQ PIO2_7_IRQHandler
    IRQ PIO2_8_IRQHandler
    IRQ PIO2_9_IRQHandler
    IRQ PIO2_10_IRQHandler
    IRQ PIO2_11_IRQHandler
    IRQ PIO3_0_IRQHandler
    IRQ PIO3_1_IRQHandler
    IRQ PIO3_2_IRQHandler
    IRQ PIO3_3_IRQHandler
    IRQ I2C0_IRQHandler
    IRQ CT16B0_IRQHandler
    IRQ CT16B1_IRQHandler
    IRQ CT32B0_IRQHandler
    IRQ CT32B1_IRQHandler
    IRQ SSP0_IRQHandler
    IRQ UART_IRQHandler
    IRQ USBIRQ_IRQHandler
    IRQ USBFIQ_IRQHandler
    IRQ ADC_IRQHandler
    IRQ WDT_IRQHandler
    IRQ BOD_IRQHandler
    IRQ PIO_3_IRQHandler
    IRQ PIO_2_IRQHandler
    IRQ PIO_1_IRQHandler
    IRQ PIO_0_IRQHandler
    IRQ SSP1_IRQHandler