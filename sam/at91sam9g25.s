// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from AT91SAM9G25.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/Atmel

// Atmel AT91SAM9G25 device: ARM926EJ Embedded Microprocessor Unit, 400MHz, Ethernet, HS USB, LPDDR/DDR2/MLC NAND support, 217 and 247 Pins (refer to http://www.atmel.com/devices/SAM9G25.aspx for more)
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
    .long FIQ_IRQHandler
    .long DBGU_IRQHandler
    .long PIOA_IRQHandler
    .long PIOC_IRQHandler
    .long SMD_IRQHandler
    .long USART0_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long USART3_IRQHandler
    .long TWI0_IRQHandler
    .long TWI1_IRQHandler
    .long TWI2_IRQHandler
    .long HSMCI0_IRQHandler
    .long SPI0_IRQHandler
    .long SPI1_IRQHandler
    .long UART0_IRQHandler
    .long UART1_IRQHandler
    .long TC0_IRQHandler
    .long PWM_IRQHandler
    .long ADC_IRQHandler
    .long DMAC0_IRQHandler
    .long DMAC1_IRQHandler
    .long 0
    .long UDPHS_IRQHandler
    .long EMAC_IRQHandler
    .long ISI_IRQHandler
    .long HSMCI1_IRQHandler
    .long 0
    .long SSC_IRQHandler
    .long 0
    .long 0
    .long IRQ_IRQHandler

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
    IRQ FIQ_IRQHandler
    IRQ DBGU_IRQHandler
    IRQ PIOA_IRQHandler
    IRQ PIOC_IRQHandler
    IRQ SMD_IRQHandler
    IRQ USART0_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ USART3_IRQHandler
    IRQ TWI0_IRQHandler
    IRQ TWI1_IRQHandler
    IRQ TWI2_IRQHandler
    IRQ HSMCI0_IRQHandler
    IRQ SPI0_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ UART0_IRQHandler
    IRQ UART1_IRQHandler
    IRQ TC0_IRQHandler
    IRQ PWM_IRQHandler
    IRQ ADC_IRQHandler
    IRQ DMAC0_IRQHandler
    IRQ DMAC1_IRQHandler
    IRQ UDPHS_IRQHandler
    IRQ EMAC_IRQHandler
    IRQ ISI_IRQHandler
    IRQ HSMCI1_IRQHandler
    IRQ SSC_IRQHandler
    IRQ IRQ_IRQHandler
