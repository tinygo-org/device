// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from ATSAMD21G17L.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/Atmel

// Microchip ATSAMD21G17L device: Cortex-M0+ Microcontroller with 128KB Flash, 16KB SRAM, QFN48_LIGHTING-pin package
//
//     Copyright (c) 2018 Microchip Technology Inc.
//
//     SPDX-License-Identifier: Apache-2.0
//
//     Licensed under the Apache License, Version 2.0 (the "License");
//     you may not use this file except in compliance with the License.
//     You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//     Unless required by applicable law or agreed to in writing, software
//     distributed under the License is distributed on an "AS IS" BASIS,
//     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//     See the License for the specific language governing permissions and
//     limitations under the License.

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
    .long PM_IRQHandler
    .long SYSCTRL_IRQHandler
    .long WDT_IRQHandler
    .long RTC_IRQHandler
    .long EIC_IRQHandler
    .long NVMCTRL_IRQHandler
    .long DMAC_IRQHandler
    .long 0
    .long EVSYS_IRQHandler
    .long SERCOM0_IRQHandler
    .long SERCOM1_IRQHandler
    .long SERCOM2_IRQHandler
    .long SERCOM3_IRQHandler
    .long SERCOM4_IRQHandler
    .long SERCOM5_IRQHandler
    .long TCC0_IRQHandler
    .long TCC1_IRQHandler
    .long TCC2_IRQHandler
    .long TC3_IRQHandler
    .long TC4_IRQHandler
    .long TC5_IRQHandler
    .long TC6_IRQHandler
    .long TC7_IRQHandler
    .long ADC_IRQHandler
    .long AC_IRQHandler
    .long DAC_IRQHandler
    .long 0
    .long 0
    .long AC1_IRQHandler
    .long TCC3_IRQHandler

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
    IRQ PM_IRQHandler
    IRQ SYSCTRL_IRQHandler
    IRQ WDT_IRQHandler
    IRQ RTC_IRQHandler
    IRQ EIC_IRQHandler
    IRQ NVMCTRL_IRQHandler
    IRQ DMAC_IRQHandler
    IRQ EVSYS_IRQHandler
    IRQ SERCOM0_IRQHandler
    IRQ SERCOM1_IRQHandler
    IRQ SERCOM2_IRQHandler
    IRQ SERCOM3_IRQHandler
    IRQ SERCOM4_IRQHandler
    IRQ SERCOM5_IRQHandler
    IRQ TCC0_IRQHandler
    IRQ TCC1_IRQHandler
    IRQ TCC2_IRQHandler
    IRQ TC3_IRQHandler
    IRQ TC4_IRQHandler
    IRQ TC5_IRQHandler
    IRQ TC6_IRQHandler
    IRQ TC7_IRQHandler
    IRQ ADC_IRQHandler
    IRQ AC_IRQHandler
    IRQ DAC_IRQHandler
    IRQ AC1_IRQHandler
    IRQ TCC3_IRQHandler
