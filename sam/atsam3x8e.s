// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from ATSAM3X8E.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/Atmel

// Atmel ATSAM3X8E Microcontroller
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
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long PMC_IRQHandler
    .long EFC0_IRQHandler
    .long EFC1_IRQHandler
    .long UART_IRQHandler
    .long 0
    .long 0
    .long PIOA_IRQHandler
    .long PIOB_IRQHandler
    .long PIOC_IRQHandler
    .long PIOD_IRQHandler
    .long 0
    .long 0
    .long USART0_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long USART3_IRQHandler
    .long HSMCI_IRQHandler
    .long TWI0_IRQHandler
    .long TWI1_IRQHandler
    .long SPI0_IRQHandler
    .long 0
    .long SSC_IRQHandler
    .long TC0_IRQHandler
    .long TC1_IRQHandler
    .long TC2_IRQHandler
    .long TC3_IRQHandler
    .long TC4_IRQHandler
    .long TC5_IRQHandler
    .long TC6_IRQHandler
    .long TC7_IRQHandler
    .long TC8_IRQHandler
    .long PWM_IRQHandler
    .long ADC_IRQHandler
    .long DACC_IRQHandler
    .long DMAC_IRQHandler
    .long UOTGHS_IRQHandler
    .long TRNG_IRQHandler
    .long EMAC_IRQHandler
    .long CAN0_IRQHandler
    .long CAN1_IRQHandler

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
    IRQ PMC_IRQHandler
    IRQ EFC0_IRQHandler
    IRQ EFC1_IRQHandler
    IRQ UART_IRQHandler
    IRQ PIOA_IRQHandler
    IRQ PIOB_IRQHandler
    IRQ PIOC_IRQHandler
    IRQ PIOD_IRQHandler
    IRQ USART0_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ USART3_IRQHandler
    IRQ HSMCI_IRQHandler
    IRQ TWI0_IRQHandler
    IRQ TWI1_IRQHandler
    IRQ SPI0_IRQHandler
    IRQ SSC_IRQHandler
    IRQ TC0_IRQHandler
    IRQ TC1_IRQHandler
    IRQ TC2_IRQHandler
    IRQ TC3_IRQHandler
    IRQ TC4_IRQHandler
    IRQ TC5_IRQHandler
    IRQ TC6_IRQHandler
    IRQ TC7_IRQHandler
    IRQ TC8_IRQHandler
    IRQ PWM_IRQHandler
    IRQ ADC_IRQHandler
    IRQ DACC_IRQHandler
    IRQ DMAC_IRQHandler
    IRQ UOTGHS_IRQHandler
    IRQ TRNG_IRQHandler
    IRQ EMAC_IRQHandler
    IRQ CAN0_IRQHandler
    IRQ CAN1_IRQHandler
