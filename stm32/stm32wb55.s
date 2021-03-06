// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32wb55.svd, see https://github.com/tinygo-org/stm32-svd

// STM32WBxx_CM4
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
    .long WWDG_IRQHandler
    .long PVD_IRQHandler
    .long RTC_TAMP_IRQHandler
    .long RTC_WKUP_IRQHandler
    .long FLASH_IRQHandler
    .long RCC_IRQHandler
    .long EXTI0_IRQHandler
    .long EXTI1_IRQHandler
    .long EXTI2_IRQHandler
    .long EXTI3_IRQHandler
    .long EXTI4_IRQHandler
    .long DMA1_Channel1_IRQHandler
    .long DMA1_Channel2_IRQHandler
    .long DMA1_Channel3_IRQHandler
    .long DMA1_Channel4_IRQHandler
    .long DMA1_Channel5_IRQHandler
    .long DMA1_Channel6_IRQHandler
    .long DMA1_Channel7_IRQHandler
    .long ADC1_IRQHandler
    .long USB_HP_IRQHandler
    .long USB_LP_IRQHandler
    .long C2SEV_IRQHandler
    .long COMP_IRQHandler
    .long EXTI9_5_IRQHandler
    .long TIM1_BRK_IRQHandler
    .long TIM1_UP_IRQHandler
    .long TIM1_TRG_COM_TIM17_IRQHandler
    .long TIM1_CC_IRQHandler
    .long TIM2_IRQHandler
    .long PKA_IRQHandler
    .long I2C1_EV_IRQHandler
    .long I2C1_ER_IRQHandler
    .long I2C3_EV_IRQHandler
    .long I2C3_ER_IRQHandler
    .long SPI1_IRQHandler
    .long SPI2_IRQHandler
    .long USART1_IRQHandler
    .long LPUART1_IRQHandler
    .long SAI1_IRQHandler
    .long TSC_IRQHandler
    .long EXTI15_10_IRQHandler
    .long RTC_ALARM_IRQHandler
    .long CRS_IT_IRQHandler
    .long PWR_SOTF_IRQHandler
    .long IPCC_C1_RX_IT_IRQHandler
    .long IPCC_C1_TX_IT_IRQHandler
    .long HSEM_IRQHandler
    .long LPTIM1_IRQHandler
    .long LPTIM2_IRQHandler
    .long LCD_IRQHandler
    .long QUADSPI_IRQHandler
    .long AES1_IRQHandler
    .long AES2_IRQHandler
    .long True_RNG_IRQHandler
    .long FPU_IRQHandler
    .long DMA2_CH1_IRQHandler
    .long DMA2_CH2_IRQHandler
    .long DMA2_CH3_IRQHandler
    .long DMA2_CH4_IRQHandler
    .long DMA2_CH5_IRQHandler
    .long DMA2_CH6_IRQHandler
    .long DMA2_CH7_IRQHandler
    .long DMAMUX_OVR_IRQHandler

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
    IRQ WWDG_IRQHandler
    IRQ PVD_IRQHandler
    IRQ RTC_TAMP_IRQHandler
    IRQ RTC_WKUP_IRQHandler
    IRQ FLASH_IRQHandler
    IRQ RCC_IRQHandler
    IRQ EXTI0_IRQHandler
    IRQ EXTI1_IRQHandler
    IRQ EXTI2_IRQHandler
    IRQ EXTI3_IRQHandler
    IRQ EXTI4_IRQHandler
    IRQ DMA1_Channel1_IRQHandler
    IRQ DMA1_Channel2_IRQHandler
    IRQ DMA1_Channel3_IRQHandler
    IRQ DMA1_Channel4_IRQHandler
    IRQ DMA1_Channel5_IRQHandler
    IRQ DMA1_Channel6_IRQHandler
    IRQ DMA1_Channel7_IRQHandler
    IRQ ADC1_IRQHandler
    IRQ USB_HP_IRQHandler
    IRQ USB_LP_IRQHandler
    IRQ C2SEV_IRQHandler
    IRQ COMP_IRQHandler
    IRQ EXTI9_5_IRQHandler
    IRQ TIM1_BRK_IRQHandler
    IRQ TIM1_UP_IRQHandler
    IRQ TIM1_TRG_COM_TIM17_IRQHandler
    IRQ TIM1_CC_IRQHandler
    IRQ TIM2_IRQHandler
    IRQ PKA_IRQHandler
    IRQ I2C1_EV_IRQHandler
    IRQ I2C1_ER_IRQHandler
    IRQ I2C3_EV_IRQHandler
    IRQ I2C3_ER_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ SPI2_IRQHandler
    IRQ USART1_IRQHandler
    IRQ LPUART1_IRQHandler
    IRQ SAI1_IRQHandler
    IRQ TSC_IRQHandler
    IRQ EXTI15_10_IRQHandler
    IRQ RTC_ALARM_IRQHandler
    IRQ CRS_IT_IRQHandler
    IRQ PWR_SOTF_IRQHandler
    IRQ IPCC_C1_RX_IT_IRQHandler
    IRQ IPCC_C1_TX_IT_IRQHandler
    IRQ HSEM_IRQHandler
    IRQ LPTIM1_IRQHandler
    IRQ LPTIM2_IRQHandler
    IRQ LCD_IRQHandler
    IRQ QUADSPI_IRQHandler
    IRQ AES1_IRQHandler
    IRQ AES2_IRQHandler
    IRQ True_RNG_IRQHandler
    IRQ FPU_IRQHandler
    IRQ DMA2_CH1_IRQHandler
    IRQ DMA2_CH2_IRQHandler
    IRQ DMA2_CH3_IRQHandler
    IRQ DMA2_CH4_IRQHandler
    IRQ DMA2_CH5_IRQHandler
    IRQ DMA2_CH6_IRQHandler
    IRQ DMA2_CH7_IRQHandler
    IRQ DMAMUX_OVR_IRQHandler
