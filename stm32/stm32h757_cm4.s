// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32h747cm4.svd, see https://github.com/tinygo-org/stm32-svd

// STM32H757_CM4
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
    .long WWDG2_IRQHandler
    .long PVD_PVM_IRQHandler
    .long RTC_TAMP_STAMP_CSS_LSE_IRQHandler
    .long RTC_WKUP_IRQHandler
    .long 0
    .long RCC_IRQHandler
    .long EXTI0_IRQHandler
    .long EXTI1_IRQHandler
    .long EXTI2_IRQHandler
    .long EXTI3_IRQHandler
    .long EXTI4_IRQHandler
    .long DMA_STR0_IRQHandler
    .long DMA_STR1_IRQHandler
    .long DMA_STR2_IRQHandler
    .long DMA_STR3_IRQHandler
    .long DMA_STR4_IRQHandler
    .long DMA_STR5_IRQHandler
    .long DMA_STR6_IRQHandler
    .long ADC1_2_IRQHandler
    .long FDCAN1_IT0_IRQHandler
    .long FDCAN2_IT0_IRQHandler
    .long FDCAN1_IT1_IRQHandler
    .long FDCAN2_IT1_IRQHandler
    .long EXTI9_5_IRQHandler
    .long TIM1_BRK_IRQHandler
    .long TIM1_UP_IRQHandler
    .long TIM1_TRG_COM_IRQHandler
    .long TIM_CC_IRQHandler
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
    .long 0
    .long TIM8_BRK_TIM12_IRQHandler
    .long TIM8_UP_TIM13_IRQHandler
    .long TIM8_TRG_COM_TIM14_IRQHandler
    .long TIM8_CC_IRQHandler
    .long DMA1_STR7_IRQHandler
    .long FMC_IRQHandler
    .long SDMMC1_IRQHandler
    .long TIM5_IRQHandler
    .long SPI3_IRQHandler
    .long UART4_IRQHandler
    .long UART5_IRQHandler
    .long TIM6_DAC_IRQHandler
    .long TIM7_IRQHandler
    .long DMA2_STR0_IRQHandler
    .long DMA2_STR1_IRQHandler
    .long DMA2_STR2_IRQHandler
    .long DMA2_STR3_IRQHandler
    .long DMA2_STR4_IRQHandler
    .long ETH_IRQHandler
    .long ETH_WKUP_IRQHandler
    .long FDCAN_CAL_IRQHandler
    .long cm7_sev_it_IRQHandler
    .long 0
    .long 0
    .long 0
    .long DMA2_STR5_IRQHandler
    .long DMA2_STR6_IRQHandler
    .long DMA2_STR7_IRQHandler
    .long USART6_IRQHandler
    .long I2C3_EV_IRQHandler
    .long I2C3_ER_IRQHandler
    .long OTG_HS_EP1_OUT_IRQHandler
    .long OTG_HS_EP1_IN_IRQHandler
    .long OTG_HS_WKUP_IRQHandler
    .long OTG_HS_IRQHandler
    .long DCMI_IRQHandler
    .long CRYP_IRQHandler
    .long HASH_RNG_IRQHandler
    .long FPU_IRQHandler
    .long UART7_IRQHandler
    .long UART8_IRQHandler
    .long SPI4_IRQHandler
    .long SPI5_IRQHandler
    .long SPI6_IRQHandler
    .long SAI1_IRQHandler
    .long LTDC_IRQHandler
    .long LTDC_ER_IRQHandler
    .long DMA2D_IRQHandler
    .long SAI2_IRQHandler
    .long QUADSPI_IRQHandler
    .long LPTIM1_IRQHandler
    .long CEC_IRQHandler
    .long I2C4_EV_IRQHandler
    .long I2C4_ER_IRQHandler
    .long SPDIF_IRQHandler
    .long OTG_FS_EP1_OUT_IRQHandler
    .long OTG_FS_EP1_IN_IRQHandler
    .long OTG_FS_WKUP_IRQHandler
    .long OTG_FS_IRQHandler
    .long DMAMUX1_OV_IRQHandler
    .long HRTIM1_MST_IRQHandler
    .long HRTIM1_TIMA_IRQHandler
    .long HRTIM_TIMB_IRQHandler
    .long HRTIM1_TIMC_IRQHandler
    .long HRTIM1_TIMD_IRQHandler
    .long HRTIM_TIME_IRQHandler
    .long HRTIM1_FLT_IRQHandler
    .long DFSDM1_FLT0_IRQHandler
    .long DFSDM1_FLT1_IRQHandler
    .long DFSDM1_FLT2_IRQHandler
    .long DFSDM1_FLT3_IRQHandler
    .long SAI3_IRQHandler
    .long SWPMI1_IRQHandler
    .long TIM15_IRQHandler
    .long TIM16_IRQHandler
    .long TIM17_IRQHandler
    .long MDIOS_WKUP_IRQHandler
    .long MDIOS_IRQHandler
    .long JPEG_IRQHandler
    .long MDMA_IRQHandler
    .long DSI_IRQHandler
    .long SDMMC_IRQHandler
    .long HSEM0_IRQHandler
    .long 0
    .long ADC3_IRQHandler
    .long DMAMUX2_OVR_IRQHandler
    .long BDMA_CH1_IRQHandler
    .long BDMA_CH2_IRQHandler
    .long BDMA_CH3_IRQHandler
    .long BDMA_CH4_IRQHandler
    .long BDMA_CH5_IRQHandler
    .long BDMA_CH6_IRQHandler
    .long BDMA_CH7_IRQHandler
    .long BDMA_CH8_IRQHandler
    .long COMP_IRQHandler
    .long LPTIM2_IRQHandler
    .long LPTIM3_IRQHandler
    .long LPTIM4_IRQHandler
    .long LPTIM5_IRQHandler
    .long LPUART_IRQHandler
    .long WWDG1_RST_IRQHandler
    .long CRS_IRQHandler
    .long RAMECC_IRQHandler
    .long SAI4_IRQHandler
    .long 0
    .long HOLD_CORE_IRQHandler
    .long WKUP_IRQHandler

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
    IRQ WWDG2_IRQHandler
    IRQ PVD_PVM_IRQHandler
    IRQ RTC_TAMP_STAMP_CSS_LSE_IRQHandler
    IRQ RTC_WKUP_IRQHandler
    IRQ RCC_IRQHandler
    IRQ EXTI0_IRQHandler
    IRQ EXTI1_IRQHandler
    IRQ EXTI2_IRQHandler
    IRQ EXTI3_IRQHandler
    IRQ EXTI4_IRQHandler
    IRQ DMA_STR0_IRQHandler
    IRQ DMA1_STR0_IRQHandler
    IRQ DMA_STR1_IRQHandler
    IRQ DMA1_STR1_IRQHandler
    IRQ DMA_STR2_IRQHandler
    IRQ DMA1_STR2_IRQHandler
    IRQ DMA_STR3_IRQHandler
    IRQ DMA1_STR3_IRQHandler
    IRQ DMA_STR4_IRQHandler
    IRQ DMA1_STR4_IRQHandler
    IRQ DMA_STR5_IRQHandler
    IRQ DMA1_STR5_IRQHandler
    IRQ DMA_STR6_IRQHandler
    IRQ DMA1_STR6_IRQHandler
    IRQ ADC1_2_IRQHandler
    IRQ FDCAN1_IT0_IRQHandler
    IRQ FDCAN2_IT0_IRQHandler
    IRQ FDCAN1_IT1_IRQHandler
    IRQ FDCAN2_IT1_IRQHandler
    IRQ EXTI9_5_IRQHandler
    IRQ TIM1_BRK_IRQHandler
    IRQ TIM1_UP_IRQHandler
    IRQ TIM1_TRG_COM_IRQHandler
    IRQ TIM_CC_IRQHandler
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
    IRQ TIM8_BRK_TIM12_IRQHandler
    IRQ TIM8_UP_TIM13_IRQHandler
    IRQ TIM8_TRG_COM_TIM14_IRQHandler
    IRQ TIM8_CC_IRQHandler
    IRQ DMA1_STR7_IRQHandler
    IRQ FMC_IRQHandler
    IRQ SDMMC1_IRQHandler
    IRQ TIM5_IRQHandler
    IRQ SPI3_IRQHandler
    IRQ UART4_IRQHandler
    IRQ UART5_IRQHandler
    IRQ TIM6_DAC_IRQHandler
    IRQ TIM7_IRQHandler
    IRQ DMA2_STR0_IRQHandler
    IRQ DMA2_STR1_IRQHandler
    IRQ DMA2_STR2_IRQHandler
    IRQ DMA2_STR3_IRQHandler
    IRQ DMA2_STR4_IRQHandler
    IRQ ETH_IRQHandler
    IRQ ETH_WKUP_IRQHandler
    IRQ FDCAN_CAL_IRQHandler
    IRQ cm7_sev_it_IRQHandler
    IRQ DMA2_STR5_IRQHandler
    IRQ DMA2_STR6_IRQHandler
    IRQ DMA2_STR7_IRQHandler
    IRQ USART6_IRQHandler
    IRQ I2C3_EV_IRQHandler
    IRQ I2C3_ER_IRQHandler
    IRQ OTG_HS_EP1_OUT_IRQHandler
    IRQ OTG_HS_EP1_IN_IRQHandler
    IRQ OTG_HS_WKUP_IRQHandler
    IRQ OTG_HS_IRQHandler
    IRQ DCMI_IRQHandler
    IRQ CRYP_IRQHandler
    IRQ HASH_RNG_IRQHandler
    IRQ FPU_IRQHandler
    IRQ UART7_IRQHandler
    IRQ UART8_IRQHandler
    IRQ SPI4_IRQHandler
    IRQ SPI5_IRQHandler
    IRQ SPI6_IRQHandler
    IRQ SAI1_IRQHandler
    IRQ LTDC_IRQHandler
    IRQ LTDC_ER_IRQHandler
    IRQ DMA2D_IRQHandler
    IRQ SAI2_IRQHandler
    IRQ QUADSPI_IRQHandler
    IRQ LPTIM1_IRQHandler
    IRQ CEC_IRQHandler
    IRQ I2C4_EV_IRQHandler
    IRQ I2C4_ER_IRQHandler
    IRQ SPDIF_IRQHandler
    IRQ OTG_FS_EP1_OUT_IRQHandler
    IRQ OTG_FS_EP1_IN_IRQHandler
    IRQ OTG_FS_WKUP_IRQHandler
    IRQ OTG_FS_IRQHandler
    IRQ DMAMUX1_OV_IRQHandler
    IRQ HRTIM1_MST_IRQHandler
    IRQ HRTIM1_TIMA_IRQHandler
    IRQ HRTIM_TIMB_IRQHandler
    IRQ HRTIM1_TIMC_IRQHandler
    IRQ HRTIM1_TIMD_IRQHandler
    IRQ HRTIM_TIME_IRQHandler
    IRQ HRTIM1_FLT_IRQHandler
    IRQ DFSDM1_FLT0_IRQHandler
    IRQ DFSDM1_FLT1_IRQHandler
    IRQ DFSDM1_FLT2_IRQHandler
    IRQ DFSDM1_FLT3_IRQHandler
    IRQ SAI3_IRQHandler
    IRQ SWPMI1_IRQHandler
    IRQ TIM15_IRQHandler
    IRQ TIM16_IRQHandler
    IRQ TIM17_IRQHandler
    IRQ MDIOS_WKUP_IRQHandler
    IRQ MDIOS_IRQHandler
    IRQ JPEG_IRQHandler
    IRQ MDMA_IRQHandler
    IRQ DSI_IRQHandler
    IRQ SDMMC_IRQHandler
    IRQ HSEM0_IRQHandler
    IRQ ADC3_IRQHandler
    IRQ DMAMUX2_OVR_IRQHandler
    IRQ BDMA_CH1_IRQHandler
    IRQ BDMA_CH2_IRQHandler
    IRQ BDMA_CH3_IRQHandler
    IRQ BDMA_CH4_IRQHandler
    IRQ BDMA_CH5_IRQHandler
    IRQ BDMA_CH6_IRQHandler
    IRQ BDMA_CH7_IRQHandler
    IRQ BDMA_CH8_IRQHandler
    IRQ COMP_IRQHandler
    IRQ LPTIM2_IRQHandler
    IRQ LPTIM3_IRQHandler
    IRQ LPTIM4_IRQHandler
    IRQ LPTIM5_IRQHandler
    IRQ LPUART_IRQHandler
    IRQ WWDG1_RST_IRQHandler
    IRQ CRS_IRQHandler
    IRQ RAMECC_IRQHandler
    IRQ SAI4_IRQHandler
    IRQ HOLD_CORE_IRQHandler
    IRQ WKUP_IRQHandler
