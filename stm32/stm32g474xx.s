// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32g474.svd, see https://github.com/tinygo-org/stm32-svd

// STM32G474xx
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
    .long PVD_PVM_IRQHandler
    .long RTC_TAMP_CSS_LSE_IRQHandler
    .long RTC_WKUP_IRQHandler
    .long FLASH_IRQHandler
    .long RCC_IRQHandler
    .long EXTI0_IRQHandler
    .long EXTI1_IRQHandler
    .long EXTI2_IRQHandler
    .long EXTI3_IRQHandler
    .long EXTI4_IRQHandler
    .long DMA1_CH1_IRQHandler
    .long DMA1_CH2_IRQHandler
    .long DMA1_CH3_IRQHandler
    .long DMA1_CH4_IRQHandler
    .long DMA1_CH5_IRQHandler
    .long DMA1_CH6_IRQHandler
    .long DMA1_CH7_IRQHandler
    .long ADC1_2_IRQHandler
    .long USB_HP_IRQHandler
    .long USB_LP_IRQHandler
    .long fdcan1_intr1_it_IRQHandler
    .long fdcan1_intr0_it_IRQHandler
    .long EXTI9_5_IRQHandler
    .long TIM1_BRK_TIM15_IRQHandler
    .long TIM1_UP_TIM16_IRQHandler
    .long TIM1_TRG_COM_IRQHandler
    .long TIM1_CC_IRQHandler
    .long TIM2_IRQHandler
    .long TIM3_IRQHandler
    .long TIM4_IRQHandler
    .long I2C1_EV_IRQHandler
    .long I2C1_ER_IRQHandler
    .long I2C2_EV_IRQHandler
    .long I2C2_ER_IRQHandler
    .long SPI1_IRQHandler
    .long SPI2_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long USART3_IRQHandler
    .long EXTI15_10_IRQHandler
    .long RTC_ALARM_IRQHandler
    .long USBWakeUP_IRQHandler
    .long TIM8_BRK_IRQHandler
    .long TIM8_UP_IRQHandler
    .long TIM8_TRG_COM_IRQHandler
    .long TIM8_CC_IRQHandler
    .long ADC3_IRQHandler
    .long FMC_IRQHandler
    .long LPTIM1_IRQHandler
    .long TIM5_IRQHandler
    .long SPI3_IRQHandler
    .long UART4_IRQHandler
    .long UART5_IRQHandler
    .long TIM6_DACUNDER_IRQHandler
    .long TIM7_IRQHandler
    .long DMA2_CH1_IRQHandler
    .long DMA2_CH2_IRQHandler
    .long DMA2_CH3_IRQHandler
    .long DMA2_CH4_IRQHandler
    .long DMA2_CH5_IRQHandler
    .long ADC4_IRQHandler
    .long ADC5_IRQHandler
    .long UCPD1_IRQHandler
    .long COMP1_2_3_IRQHandler
    .long COMP4_5_6_IRQHandler
    .long COMP7_IRQHandler
    .long HRTIM_Master_IRQn_IRQHandler
    .long HRTIM_TIMA_IRQn_IRQHandler
    .long HRTIM_TIMB_IRQn_IRQHandler
    .long HRTIM_TIMC_IRQn_IRQHandler
    .long HRTIM_TIMD_IRQn_IRQHandler
    .long HRTIM_TIME_IRQn_IRQHandler
    .long HRTIM_TIM_FLT_IRQn_IRQHandler
    .long HRTIM_TIMF_IRQn_IRQHandler
    .long CRS_IRQHandler
    .long SAI_IRQHandler
    .long TIM20_BRK_IRQHandler
    .long TIM20_UP_IRQHandler
    .long TIM20_TRG_COM_IRQHandler
    .long TIM20_CC_IRQHandler
    .long FPU_IRQHandler
    .long I2C4_EV_IRQHandler
    .long I2C4_ER_IRQHandler
    .long SPI4_IRQHandler
    .long 0
    .long FDCAN2_intr0_IRQHandler
    .long FDCAN2_intr1_IRQHandler
    .long FDCAN3_intr0_IRQHandler
    .long FDCAN3_intr1_IRQHandler
    .long RNG_IRQHandler
    .long LPUART_IRQHandler
    .long I2C3_EV_IRQHandler
    .long I2C3_ER_IRQHandler
    .long DMAMUX_OVR_IRQHandler
    .long QUADSPI_IRQHandler
    .long DMA1_CH8_IRQHandler
    .long DMA2_CH6_IRQHandler
    .long DMA2_CH7_IRQHandler
    .long DMA2_CH8_IRQHandler
    .long Cordic_IRQHandler
    .long FMAC_IRQHandler

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
    IRQ PVD_PVM_IRQHandler
    IRQ RTC_TAMP_CSS_LSE_IRQHandler
    IRQ RTC_WKUP_IRQHandler
    IRQ FLASH_IRQHandler
    IRQ RCC_IRQHandler
    IRQ EXTI0_IRQHandler
    IRQ EXTI1_IRQHandler
    IRQ EXTI2_IRQHandler
    IRQ EXTI3_IRQHandler
    IRQ EXTI4_IRQHandler
    IRQ DMA1_CH1_IRQHandler
    IRQ DMA1_CH2_IRQHandler
    IRQ DMA1_CH3_IRQHandler
    IRQ DMA1_CH4_IRQHandler
    IRQ DMA1_CH5_IRQHandler
    IRQ DMA1_CH6_IRQHandler
    IRQ DMA1_CH7_IRQHandler
    IRQ ADC1_2_IRQHandler
    IRQ USB_HP_IRQHandler
    IRQ USB_LP_IRQHandler
    IRQ fdcan1_intr1_it_IRQHandler
    IRQ fdcan1_intr0_it_IRQHandler
    IRQ EXTI9_5_IRQHandler
    IRQ TIM1_BRK_TIM15_IRQHandler
    IRQ TIM1_UP_TIM16_IRQHandler
    IRQ TIM1_TRG_COM_IRQHandler
    IRQ TIM1_CC_IRQHandler
    IRQ TIM2_IRQHandler
    IRQ TIM3_IRQHandler
    IRQ TIM4_IRQHandler
    IRQ I2C1_EV_IRQHandler
    IRQ I2C1_ER_IRQHandler
    IRQ I2C2_EV_IRQHandler
    IRQ I2C2_ER_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ SPI2_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ USART3_IRQHandler
    IRQ EXTI15_10_IRQHandler
    IRQ RTC_ALARM_IRQHandler
    IRQ USBWakeUP_IRQHandler
    IRQ TIM8_BRK_IRQHandler
    IRQ TIM8_UP_IRQHandler
    IRQ TIM8_TRG_COM_IRQHandler
    IRQ TIM8_CC_IRQHandler
    IRQ ADC3_IRQHandler
    IRQ FMC_IRQHandler
    IRQ LPTIM1_IRQHandler
    IRQ TIM5_IRQHandler
    IRQ SPI3_IRQHandler
    IRQ UART4_IRQHandler
    IRQ UART5_IRQHandler
    IRQ TIM6_DACUNDER_IRQHandler
    IRQ TIM7_IRQHandler
    IRQ DMA2_CH1_IRQHandler
    IRQ DMA2_CH2_IRQHandler
    IRQ DMA2_CH3_IRQHandler
    IRQ DMA2_CH4_IRQHandler
    IRQ DMA2_CH5_IRQHandler
    IRQ ADC4_IRQHandler
    IRQ ADC5_IRQHandler
    IRQ UCPD1_IRQHandler
    IRQ COMP1_2_3_IRQHandler
    IRQ COMP4_5_6_IRQHandler
    IRQ COMP7_IRQHandler
    IRQ HRTIM_Master_IRQn_IRQHandler
    IRQ HRTIM_TIMA_IRQn_IRQHandler
    IRQ HRTIM_TIMB_IRQn_IRQHandler
    IRQ HRTIM_TIMC_IRQn_IRQHandler
    IRQ HRTIM_TIMD_IRQn_IRQHandler
    IRQ HRTIM_TIME_IRQn_IRQHandler
    IRQ HRTIM_TIM_FLT_IRQn_IRQHandler
    IRQ HRTIM_TIMF_IRQn_IRQHandler
    IRQ CRS_IRQHandler
    IRQ SAI_IRQHandler
    IRQ TIM20_BRK_IRQHandler
    IRQ TIM20_UP_IRQHandler
    IRQ TIM20_TRG_COM_IRQHandler
    IRQ TIM20_CC_IRQHandler
    IRQ FPU_IRQHandler
    IRQ I2C4_EV_IRQHandler
    IRQ I2C4_ER_IRQHandler
    IRQ SPI4_IRQHandler
    IRQ FDCAN2_intr0_IRQHandler
    IRQ FDCAN2_intr1_IRQHandler
    IRQ FDCAN3_intr0_IRQHandler
    IRQ FDCAN3_intr1_IRQHandler
    IRQ RNG_IRQHandler
    IRQ LPUART_IRQHandler
    IRQ I2C3_EV_IRQHandler
    IRQ I2C3_ER_IRQHandler
    IRQ DMAMUX_OVR_IRQHandler
    IRQ QUADSPI_IRQHandler
    IRQ DMA1_CH8_IRQHandler
    IRQ DMA2_CH6_IRQHandler
    IRQ DMA2_CH7_IRQHandler
    IRQ DMA2_CH8_IRQHandler
    IRQ Cordic_IRQHandler
    IRQ FMAC_IRQHandler
