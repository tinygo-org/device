// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from nrf5340_application.svd, see https://github.com/NordicSemiconductor/nrfx/tree/master/mdk

// nRF53 reference description for system-on-chip with dual ARM 32-bit Cortex-M33 microcontrollers
//
//     Copyright (c) 2010 - 2020, Nordic Semiconductor ASA All rights reserved. Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met: 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer. 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution. 3. Neither the name of Nordic Semiconductor ASA nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY, AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL NORDIC SEMICONDUCTOR ASA OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
    .long FPU_IRQHandler
    .long CACHE_IRQHandler
    .long 0
    .long SPU_IRQHandler
    .long 0
    .long CLOCK_POWER_IRQHandler
    .long 0
    .long 0
    .long SPIM0_SPIS0_TWIM0_TWIS0_UARTE0_IRQHandler
    .long SPIM1_SPIS1_TWIM1_TWIS1_UARTE1_IRQHandler
    .long SPIM4_IRQHandler
    .long SPIM2_SPIS2_TWIM2_TWIS2_UARTE2_IRQHandler
    .long SPIM3_SPIS3_TWIM3_TWIS3_UARTE3_IRQHandler
    .long GPIOTE0_IRQHandler
    .long SAADC_IRQHandler
    .long TIMER0_IRQHandler
    .long TIMER1_IRQHandler
    .long TIMER2_IRQHandler
    .long 0
    .long 0
    .long RTC0_IRQHandler
    .long RTC1_IRQHandler
    .long 0
    .long 0
    .long WDT0_IRQHandler
    .long WDT1_IRQHandler
    .long COMP_LPCOMP_IRQHandler
    .long EGU0_IRQHandler
    .long EGU1_IRQHandler
    .long EGU2_IRQHandler
    .long EGU3_IRQHandler
    .long EGU4_IRQHandler
    .long EGU5_IRQHandler
    .long PWM0_IRQHandler
    .long PWM1_IRQHandler
    .long PWM2_IRQHandler
    .long PWM3_IRQHandler
    .long 0
    .long PDM0_IRQHandler
    .long 0
    .long I2S0_IRQHandler
    .long 0
    .long IPC_IRQHandler
    .long QSPI_IRQHandler
    .long 0
    .long NFCT_IRQHandler
    .long 0
    .long GPIOTE1_IRQHandler
    .long 0
    .long 0
    .long 0
    .long QDEC0_IRQHandler
    .long QDEC1_IRQHandler
    .long 0
    .long USBD_IRQHandler
    .long USBREGULATOR_IRQHandler
    .long 0
    .long KMU_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long CRYPTOCELL_IRQHandler

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
    IRQ FPU_IRQHandler
    IRQ FPU_NS_IRQHandler
    IRQ FPU_S_IRQHandler
    IRQ CACHE_IRQHandler
    IRQ CACHE_S_IRQHandler
    IRQ SPU_IRQHandler
    IRQ SPU_S_IRQHandler
    IRQ CLOCK_POWER_IRQHandler
    IRQ CLOCK_NS_IRQHandler
    IRQ POWER_NS_IRQHandler
    IRQ CLOCK_S_IRQHandler
    IRQ POWER_S_IRQHandler
    IRQ SPIM0_SPIS0_TWIM0_TWIS0_UARTE0_IRQHandler
    IRQ SPIM0_NS_IRQHandler
    IRQ SPIS0_NS_IRQHandler
    IRQ TWIM0_NS_IRQHandler
    IRQ TWIS0_NS_IRQHandler
    IRQ UARTE0_NS_IRQHandler
    IRQ SPIM0_S_IRQHandler
    IRQ SPIS0_S_IRQHandler
    IRQ TWIM0_S_IRQHandler
    IRQ TWIS0_S_IRQHandler
    IRQ UARTE0_S_IRQHandler
    IRQ SPIM1_SPIS1_TWIM1_TWIS1_UARTE1_IRQHandler
    IRQ SPIM1_NS_IRQHandler
    IRQ SPIS1_NS_IRQHandler
    IRQ TWIM1_NS_IRQHandler
    IRQ TWIS1_NS_IRQHandler
    IRQ UARTE1_NS_IRQHandler
    IRQ SPIM1_S_IRQHandler
    IRQ SPIS1_S_IRQHandler
    IRQ TWIM1_S_IRQHandler
    IRQ TWIS1_S_IRQHandler
    IRQ UARTE1_S_IRQHandler
    IRQ SPIM4_IRQHandler
    IRQ SPIM4_NS_IRQHandler
    IRQ SPIM4_S_IRQHandler
    IRQ SPIM2_SPIS2_TWIM2_TWIS2_UARTE2_IRQHandler
    IRQ SPIM2_NS_IRQHandler
    IRQ SPIS2_NS_IRQHandler
    IRQ TWIM2_NS_IRQHandler
    IRQ TWIS2_NS_IRQHandler
    IRQ UARTE2_NS_IRQHandler
    IRQ SPIM2_S_IRQHandler
    IRQ SPIS2_S_IRQHandler
    IRQ TWIM2_S_IRQHandler
    IRQ TWIS2_S_IRQHandler
    IRQ UARTE2_S_IRQHandler
    IRQ SPIM3_SPIS3_TWIM3_TWIS3_UARTE3_IRQHandler
    IRQ SPIM3_NS_IRQHandler
    IRQ SPIS3_NS_IRQHandler
    IRQ TWIM3_NS_IRQHandler
    IRQ TWIS3_NS_IRQHandler
    IRQ UARTE3_NS_IRQHandler
    IRQ SPIM3_S_IRQHandler
    IRQ SPIS3_S_IRQHandler
    IRQ TWIM3_S_IRQHandler
    IRQ TWIS3_S_IRQHandler
    IRQ UARTE3_S_IRQHandler
    IRQ GPIOTE0_IRQHandler
    IRQ GPIOTE0_S_IRQHandler
    IRQ SAADC_IRQHandler
    IRQ SAADC_NS_IRQHandler
    IRQ SAADC_S_IRQHandler
    IRQ TIMER0_IRQHandler
    IRQ TIMER0_NS_IRQHandler
    IRQ TIMER0_S_IRQHandler
    IRQ TIMER1_IRQHandler
    IRQ TIMER1_NS_IRQHandler
    IRQ TIMER1_S_IRQHandler
    IRQ TIMER2_IRQHandler
    IRQ TIMER2_NS_IRQHandler
    IRQ TIMER2_S_IRQHandler
    IRQ RTC0_IRQHandler
    IRQ RTC0_NS_IRQHandler
    IRQ RTC0_S_IRQHandler
    IRQ RTC1_IRQHandler
    IRQ RTC1_NS_IRQHandler
    IRQ RTC1_S_IRQHandler
    IRQ WDT0_IRQHandler
    IRQ WDT0_NS_IRQHandler
    IRQ WDT0_S_IRQHandler
    IRQ WDT1_IRQHandler
    IRQ WDT1_NS_IRQHandler
    IRQ WDT1_S_IRQHandler
    IRQ COMP_LPCOMP_IRQHandler
    IRQ COMP_NS_IRQHandler
    IRQ LPCOMP_NS_IRQHandler
    IRQ COMP_S_IRQHandler
    IRQ LPCOMP_S_IRQHandler
    IRQ EGU0_IRQHandler
    IRQ EGU0_NS_IRQHandler
    IRQ EGU0_S_IRQHandler
    IRQ EGU1_IRQHandler
    IRQ EGU1_NS_IRQHandler
    IRQ EGU1_S_IRQHandler
    IRQ EGU2_IRQHandler
    IRQ EGU2_NS_IRQHandler
    IRQ EGU2_S_IRQHandler
    IRQ EGU3_IRQHandler
    IRQ EGU3_NS_IRQHandler
    IRQ EGU3_S_IRQHandler
    IRQ EGU4_IRQHandler
    IRQ EGU4_NS_IRQHandler
    IRQ EGU4_S_IRQHandler
    IRQ EGU5_IRQHandler
    IRQ EGU5_NS_IRQHandler
    IRQ EGU5_S_IRQHandler
    IRQ PWM0_IRQHandler
    IRQ PWM0_NS_IRQHandler
    IRQ PWM0_S_IRQHandler
    IRQ PWM1_IRQHandler
    IRQ PWM1_NS_IRQHandler
    IRQ PWM1_S_IRQHandler
    IRQ PWM2_IRQHandler
    IRQ PWM2_NS_IRQHandler
    IRQ PWM2_S_IRQHandler
    IRQ PWM3_IRQHandler
    IRQ PWM3_NS_IRQHandler
    IRQ PWM3_S_IRQHandler
    IRQ PDM0_IRQHandler
    IRQ PDM0_NS_IRQHandler
    IRQ PDM0_S_IRQHandler
    IRQ I2S0_IRQHandler
    IRQ I2S0_NS_IRQHandler
    IRQ I2S0_S_IRQHandler
    IRQ IPC_IRQHandler
    IRQ IPC_NS_IRQHandler
    IRQ IPC_S_IRQHandler
    IRQ QSPI_IRQHandler
    IRQ QSPI_NS_IRQHandler
    IRQ QSPI_S_IRQHandler
    IRQ NFCT_IRQHandler
    IRQ NFCT_NS_IRQHandler
    IRQ NFCT_S_IRQHandler
    IRQ GPIOTE1_IRQHandler
    IRQ GPIOTE1_NS_IRQHandler
    IRQ QDEC0_IRQHandler
    IRQ QDEC0_NS_IRQHandler
    IRQ QDEC0_S_IRQHandler
    IRQ QDEC1_IRQHandler
    IRQ QDEC1_NS_IRQHandler
    IRQ QDEC1_S_IRQHandler
    IRQ USBD_IRQHandler
    IRQ USBD_NS_IRQHandler
    IRQ USBD_S_IRQHandler
    IRQ USBREGULATOR_IRQHandler
    IRQ USBREGULATOR_NS_IRQHandler
    IRQ USBREGULATOR_S_IRQHandler
    IRQ KMU_IRQHandler
    IRQ KMU_NS_IRQHandler
    IRQ KMU_S_IRQHandler
    IRQ CRYPTOCELL_IRQHandler
    IRQ CRYPTOCELL_S_IRQHandler
