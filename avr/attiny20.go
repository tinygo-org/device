// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny20.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny20
// +build avr,attiny20

// Device information for the ATtiny20.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny20"
	ARCH   = "AVR8L"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET      = 0  // External Reset, Power-on Reset and Watchdog Reset
	IRQ_INT0       = 1  // External Interrupt Request 0
	IRQ_PCINT0     = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1     = 3  // Pin Change Interrupt Request 1
	IRQ_WDT        = 4  // Watchdog Time-out
	IRQ_TIM1_CAPT  = 5  // Timer/Counter1 Input Capture
	IRQ_TIM1_COMPA = 6  //  Timer/Counter1 Compare Match A
	IRQ_TIM1_COMPB = 7  //  Timer/Counter1 Compare Match B
	IRQ_TIM1_OVF   = 8  // Timer/Counter1 Overflow
	IRQ_TIM0_COMPA = 9  // Timer/Counter0 Compare Match A
	IRQ_TIM0_COMPB = 10 // Timer/Counter0 Compare Match B
	IRQ_TIM0_OVF   = 11 // Timer/Counter0 Overflow
	IRQ_ANA_COMP   = 12 // Analog Comparator
	IRQ_ADC_ADC    = 13 // Conversion Complete
	IRQ_TWI_SLAVE  = 14 // Two-Wire Interface
	IRQ_SPI        = 15 // Serial Peripheral Interface
	IRQ_QTRIP      = 16 // Touch Sensing
	IRQ_max        = 16 // Highest interrupt number on this device.
)

// Pseudo function call that is replaced by the compiler with the actual
// functions registered through interrupt.New.
//
//go:linkname callHandlers runtime/interrupt.callHandlers
func callHandlers(num int)

//export __vector_RESET
//go:interrupt
func interruptRESET() {
	callHandlers(IRQ_RESET)
}

//export __vector_INT0
//go:interrupt
func interruptINT0() {
	callHandlers(IRQ_INT0)
}

//export __vector_PCINT0
//go:interrupt
func interruptPCINT0() {
	callHandlers(IRQ_PCINT0)
}

//export __vector_PCINT1
//go:interrupt
func interruptPCINT1() {
	callHandlers(IRQ_PCINT1)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_TIM1_CAPT
//go:interrupt
func interruptTIM1_CAPT() {
	callHandlers(IRQ_TIM1_CAPT)
}

//export __vector_TIM1_COMPA
//go:interrupt
func interruptTIM1_COMPA() {
	callHandlers(IRQ_TIM1_COMPA)
}

//export __vector_TIM1_COMPB
//go:interrupt
func interruptTIM1_COMPB() {
	callHandlers(IRQ_TIM1_COMPB)
}

//export __vector_TIM1_OVF
//go:interrupt
func interruptTIM1_OVF() {
	callHandlers(IRQ_TIM1_OVF)
}

//export __vector_TIM0_COMPA
//go:interrupt
func interruptTIM0_COMPA() {
	callHandlers(IRQ_TIM0_COMPA)
}

//export __vector_TIM0_COMPB
//go:interrupt
func interruptTIM0_COMPB() {
	callHandlers(IRQ_TIM0_COMPB)
}

//export __vector_TIM0_OVF
//go:interrupt
func interruptTIM0_OVF() {
	callHandlers(IRQ_TIM0_OVF)
}

//export __vector_ANA_COMP
//go:interrupt
func interruptANA_COMP() {
	callHandlers(IRQ_ANA_COMP)
}

//export __vector_ADC_ADC
//go:interrupt
func interruptADC_ADC() {
	callHandlers(IRQ_ADC_ADC)
}

//export __vector_TWI_SLAVE
//go:interrupt
func interruptTWI_SLAVE() {
	callHandlers(IRQ_TWI_SLAVE)
}

//export __vector_SPI
//go:interrupt
func interruptSPI() {
	callHandlers(IRQ_SPI)
}

//export __vector_QTRIP
//go:interrupt
func interruptQTRIP() {
	callHandlers(IRQ_QTRIP)
}

// Peripherals.
var (
	// Fuses
	BYTE0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// I/O Port
	PORTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8)))
	PUEB   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7)))
	DDRB   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5)))
	PINB   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4)))
	PORTB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6)))
	PUEA   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3)))
	PORTA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	DDRA   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	PINA   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x10)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x12)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x11)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd)))

	// Analog Comparator
	ACSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x13)))
	ACSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x14)))

	// CPU Registers
	CCP    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	CLKMSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	CLKPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	RSTFLR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	NVMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	NVMCMD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))

	// External Interrupts
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9)))
	GIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb)))
	GIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc)))

	// Timer/Counter, 8-bit
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x19)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x18)))
	TIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	TIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	GTCCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x17)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x16)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x15)))

	// Two Wire Serial Interface
	TWSCRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	TWSCRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	TWSSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	TWSA   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	TWSD   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	TWSAM  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))

	// Timer/Counter, 16-bit
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1e)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1f)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1c)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1d)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1a)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1b)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
)

// Bitfields for FUSE: Fuses
const (
	// BYTE0
	BYTE0_BODLEVEL0    = 0x10 // Brown-out Detector trigger level
	BYTE0_BODLEVEL1    = 0x20 // Brown-out Detector trigger level
	BYTE0_BODLEVEL2    = 0x40 // Brown-out Detector trigger level
	BYTE0_BODLEVEL_Msk = 0x70 // Brown-out Detector trigger level
	BYTE0_CKOUT        = 0x4  // Output external clock
	BYTE0_CKOUT_Msk    = 0x4  // Output external clock
	BYTE0_WDTON        = 0x2  // Watch dog timer always on
	BYTE0_WDTON_Msk    = 0x2  // Watch dog timer always on
	BYTE0_RSTDISBL     = 0x1  // Disable external reset
	BYTE0_RSTDISBL_Msk = 0x1  // Disable external reset
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0    = 0x1 // Memory Lock
	LOCKBIT_LB1    = 0x2 // Memory Lock
	LOCKBIT_LB_Msk = 0x3 // Memory Lock
)

// Bitfields for PORT: I/O Port
const (
	// PORTCR: Port Control Register
	PORTCR_BBMB     = 0x2 // Break-Before-Make Mode Enable
	PORTCR_BBMB_Msk = 0x2 // Break-Before-Make Mode Enable
	PORTCR_BBMA     = 0x1 // Break-Before-Make Mode Enable
	PORTCR_BBMA_Msk = 0x1 // Break-Before-Make Mode Enable
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control and Status Register
	WDTCSR_WDIF     = 0x80 // Watchdog Timer Interrupt Flag
	WDTCSR_WDIF_Msk = 0x80 // Watchdog Timer Interrupt Flag
	WDTCSR_WDIE     = 0x40 // Watchdog Timer Interrupt Enable
	WDTCSR_WDIE_Msk = 0x40 // Watchdog Timer Interrupt Enable
	WDTCSR_WDP0     = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1     = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2     = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3     = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDP_Msk  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDE      = 0x8  // Watch Dog Enable
	WDTCSR_WDE_Msk  = 0x8  // Watch Dog Enable
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS     = 0x40 // Reference Selection Bit
	ADMUX_REFS_Msk = 0x40 // Reference Selection Bit
	ADMUX_MUX0     = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1     = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2     = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3     = 0x8  // Analog Channel and Gain Selection Bits
	ADMUX_MUX_Msk  = 0xf  // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN      = 0x80 // ADC Enable
	ADCSRA_ADEN_Msk  = 0x80 // ADC Enable
	ADCSRA_ADSC      = 0x40 // ADC Start Conversion
	ADCSRA_ADSC_Msk  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE     = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADATE_Msk = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF      = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIF_Msk  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE      = 0x8  // ADC Interrupt Enable
	ADCSRA_ADIE_Msk  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0     = 0x1  // ADC Prescaler Select Bits
	ADCSRA_ADPS1     = 0x2  // ADC Prescaler Select Bits
	ADCSRA_ADPS2     = 0x4  // ADC Prescaler Select Bits
	ADCSRA_ADPS_Msk  = 0x7  // ADC Prescaler Select Bits

	// ADCSRB: ADC Control and Status Register B
	ADCSRB_ADLAR     = 0x8
	ADCSRB_ADLAR_Msk = 0x8
	ADCSRB_ADTS0     = 0x1 // ADC Auto Trigger Sources
	ADCSRB_ADTS1     = 0x2 // ADC Auto Trigger Sources
	ADCSRB_ADTS2     = 0x4 // ADC Auto Trigger Sources
	ADCSRB_ADTS_Msk  = 0x7 // ADC Auto Trigger Sources

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D     = 0x80 // ADC6 Digital input Disable
	DIDR0_ADC7D_Msk = 0x80 // ADC6 Digital input Disable
	DIDR0_ADC6D     = 0x40 // ADC5 Digital input Disable
	DIDR0_ADC6D_Msk = 0x40 // ADC5 Digital input Disable
	DIDR0_ADC5D     = 0x20 // ADC4 Digital input Disable
	DIDR0_ADC5D_Msk = 0x20 // ADC4 Digital input Disable
	DIDR0_ADC4D     = 0x10 // ADC3 Digital input Disable
	DIDR0_ADC4D_Msk = 0x10 // ADC3 Digital input Disable
	DIDR0_ADC3D     = 0x8  // AREF Digital Input Disable
	DIDR0_ADC3D_Msk = 0x8  // AREF Digital Input Disable
	DIDR0_ADC2D     = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC2D_Msk = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D     = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC1D_Msk = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D     = 0x1  // ADC0 Digital input Disable
	DIDR0_ADC0D_Msk = 0x1  // ADC0 Digital input Disable
)

// Bitfields for AC: Analog Comparator
const (
	// ACSRB: Analog Comparator Control And Status Register B
	ACSRB_HSEL     = 0x80 // Hysteresis Select
	ACSRB_HSEL_Msk = 0x80 // Hysteresis Select
	ACSRB_HLEV     = 0x40 // Hysteresis Level
	ACSRB_HLEV_Msk = 0x40 // Hysteresis Level
	ACSRB_ACME     = 0x4  // Analog Comparator Multiplexer Enable
	ACSRB_ACME_Msk = 0x4  // Analog Comparator Multiplexer Enable

	// ACSRA: Analog Comparator Control And Status Register A
	ACSRA_ACD      = 0x80 // Analog Comparator Disable
	ACSRA_ACD_Msk  = 0x80 // Analog Comparator Disable
	ACSRA_ACBG     = 0x40 // Analog Comparator Bandgap Select
	ACSRA_ACBG_Msk = 0x40 // Analog Comparator Bandgap Select
	ACSRA_ACO      = 0x20 // Analog Compare Output
	ACSRA_ACO_Msk  = 0x20 // Analog Compare Output
	ACSRA_ACI      = 0x10 // Analog Comparator Interrupt Flag
	ACSRA_ACI_Msk  = 0x10 // Analog Comparator Interrupt Flag
	ACSRA_ACIE     = 0x8  // Analog Comparator Interrupt Enable
	ACSRA_ACIE_Msk = 0x8  // Analog Comparator Interrupt Enable
	ACSRA_ACIC     = 0x4  // Analog Comparator Input Capture Enable
	ACSRA_ACIC_Msk = 0x4  // Analog Comparator Input Capture Enable
	ACSRA_ACIS0    = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSRA_ACIS1    = 0x2  // Analog Comparator Interrupt Mode Select bits
	ACSRA_ACIS_Msk = 0x3  // Analog Comparator Interrupt Mode Select bits
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I     = 0x80 // Global Interrupt Enable
	SREG_I_Msk = 0x80 // Global Interrupt Enable
	SREG_T     = 0x40 // Bit Copy Storage
	SREG_T_Msk = 0x40 // Bit Copy Storage
	SREG_H     = 0x20 // Half Carry Flag
	SREG_H_Msk = 0x20 // Half Carry Flag
	SREG_S     = 0x10 // Sign Bit
	SREG_S_Msk = 0x10 // Sign Bit
	SREG_V     = 0x8  // Two's Complement Overflow Flag
	SREG_V_Msk = 0x8  // Two's Complement Overflow Flag
	SREG_N     = 0x4  // Negative Flag
	SREG_N_Msk = 0x4  // Negative Flag
	SREG_Z     = 0x2  // Zero Flag
	SREG_Z_Msk = 0x2  // Zero Flag
	SREG_C     = 0x1  // Carry Flag
	SREG_C_Msk = 0x1  // Carry Flag

	// CLKMSR: Clock Main Settings Register
	CLKMSR_CLKMS0    = 0x1 // Clock Main Select Bits
	CLKMSR_CLKMS1    = 0x2 // Clock Main Select Bits
	CLKMSR_CLKMS_Msk = 0x3 // Clock Main Select Bits

	// CLKPSR: Clock Prescale Register
	CLKPSR_CLKPS0    = 0x1 // Clock Prescaler Select Bits
	CLKPSR_CLKPS1    = 0x2 // Clock Prescaler Select Bits
	CLKPSR_CLKPS2    = 0x4 // Clock Prescaler Select Bits
	CLKPSR_CLKPS3    = 0x8 // Clock Prescaler Select Bits
	CLKPSR_CLKPS_Msk = 0xf // Clock Prescaler Select Bits

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0    = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1    = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2    = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3    = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4    = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5    = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6    = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7    = 0x80 // Oscillator Calibration
	OSCCAL_OSCCAL_Msk = 0xff // Oscillator Calibration

	// PRR: Power Reduction Register
	PRR_PRTWI      = 0x10 // Power Reduction TWI
	PRR_PRTWI_Msk  = 0x10 // Power Reduction TWI
	PRR_PRSPI      = 0x8  // Power Reduction Serial Peripheral Interface
	PRR_PRSPI_Msk  = 0x8  // Power Reduction Serial Peripheral Interface
	PRR_PRTIM1     = 0x4  // Power Reduction Timer/Counter1
	PRR_PRTIM1_Msk = 0x4  // Power Reduction Timer/Counter1
	PRR_PRTIM0     = 0x2  // Power Reduction Timer/Counter0
	PRR_PRTIM0_Msk = 0x2  // Power Reduction Timer/Counter0
	PRR_PRADC      = 0x1  // Power Reduction ADC
	PRR_PRADC_Msk  = 0x1  // Power Reduction ADC

	// RSTFLR: Reset Flag Register
	RSTFLR_WDRF      = 0x8 // Watchdog Reset Flag
	RSTFLR_WDRF_Msk  = 0x8 // Watchdog Reset Flag
	RSTFLR_EXTRF     = 0x2 // External Reset Flag
	RSTFLR_EXTRF_Msk = 0x2 // External Reset Flag
	RSTFLR_PORF      = 0x1 // Power-on Reset Flag
	RSTFLR_PORF_Msk  = 0x1 // Power-on Reset Flag

	// NVMCSR: Non-Volatile Memory Control and Status Register
	NVMCSR_NVMBSY     = 0x80 // Non-Volatile Memory Busy
	NVMCSR_NVMBSY_Msk = 0x80 // Non-Volatile Memory Busy

	// MCUCR: MCU Control Register
	MCUCR_BODS      = 0x10 // BOD Sleep
	MCUCR_BODS_Msk  = 0x10 // BOD Sleep
	MCUCR_SM0       = 0x2  // Sleep Mode
	MCUCR_SM1       = 0x4  // Sleep Mode
	MCUCR_SM2       = 0x8  // Sleep Mode
	MCUCR_SM_Msk    = 0xe  // Sleep Mode
	MCUCR_SE        = 0x1  // Sleep Enable
	MCUCR_SE_Msk    = 0x1  // Sleep Enable
	MCUCR_ISC01     = 0x80 // Interrupt Sense Control 0 Bit 1
	MCUCR_ISC01_Msk = 0x80 // Interrupt Sense Control 0 Bit 1
	MCUCR_ISC00     = 0x40 // Interrupt Sense Control 0 Bit 0
	MCUCR_ISC00_Msk = 0x40 // Interrupt Sense Control 0 Bit 0
)

// Bitfields for EXINT: External Interrupts
const (
	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT0    = 0x1 // Pin Change Enable Masks
	PCMSK1_PCINT1    = 0x2 // Pin Change Enable Masks
	PCMSK1_PCINT2    = 0x4 // Pin Change Enable Masks
	PCMSK1_PCINT3    = 0x8 // Pin Change Enable Masks
	PCMSK1_PCINT_Msk = 0xf // Pin Change Enable Masks

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT0    = 0x1  // Pin Change Enable Masks
	PCMSK0_PCINT1    = 0x2  // Pin Change Enable Masks
	PCMSK0_PCINT2    = 0x4  // Pin Change Enable Masks
	PCMSK0_PCINT3    = 0x8  // Pin Change Enable Masks
	PCMSK0_PCINT4    = 0x10 // Pin Change Enable Masks
	PCMSK0_PCINT5    = 0x20 // Pin Change Enable Masks
	PCMSK0_PCINT6    = 0x40 // Pin Change Enable Masks
	PCMSK0_PCINT7    = 0x80 // Pin Change Enable Masks
	PCMSK0_PCINT_Msk = 0xff // Pin Change Enable Masks

	// GIFR: General Interrupt Flag Register
	GIFR_PCIF0     = 0x10 // Pin Change Interrupt Flags
	GIFR_PCIF1     = 0x20 // Pin Change Interrupt Flags
	GIFR_PCIF_Msk  = 0x30 // Pin Change Interrupt Flags
	GIFR_INTF0     = 0x1  // External Interrupt Flag 0
	GIFR_INTF0_Msk = 0x1  // External Interrupt Flag 0

	// GIMSK: General Interrupt Mask Register
	GIMSK_PCIE0    = 0x10 // Pin Change Interrupt Enables
	GIMSK_PCIE1    = 0x20 // Pin Change Interrupt Enables
	GIMSK_PCIE_Msk = 0x30 // Pin Change Interrupt Enables
	GIMSK_INT0     = 0x1  // External Interrupt Request 0 Enable
	GIMSK_INT0_Msk = 0x1  // External Interrupt Request 0 Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0A: Timer/Counter 0 Control Register A
	TCCR0A_COM0A0    = 0x40 // Compare Output Mode for Channel A bits
	TCCR0A_COM0A1    = 0x80 // Compare Output Mode for Channel A bits
	TCCR0A_COM0A_Msk = 0xc0 // Compare Output Mode for Channel A bits
	TCCR0A_COM0B0    = 0x10 // Compare Output Mode for Channel B bits
	TCCR0A_COM0B1    = 0x20 // Compare Output Mode for Channel B bits
	TCCR0A_COM0B_Msk = 0x30 // Compare Output Mode for Channel B bits
	TCCR0A_WGM00     = 0x1  // Waveform Generation Mode
	TCCR0A_WGM01     = 0x2  // Waveform Generation Mode
	TCCR0A_WGM0_Msk  = 0x3  // Waveform Generation Mode

	// TCCR0B: Timer/Counter 0 Control Register B
	TCCR0B_FOC0A     = 0x80 // Force Output Compare A
	TCCR0B_FOC0A_Msk = 0x80 // Force Output Compare A
	TCCR0B_FOC0B     = 0x40 // Force Output Compare B
	TCCR0B_FOC0B_Msk = 0x40 // Force Output Compare B
	TCCR0B_WGM02     = 0x8  // Waveform Generation Mode
	TCCR0B_WGM02_Msk = 0x8  // Waveform Generation Mode
	TCCR0B_CS00      = 0x1  // Clock Select
	TCCR0B_CS01      = 0x2  // Clock Select
	TCCR0B_CS02      = 0x4  // Clock Select
	TCCR0B_CS0_Msk   = 0x7  // Clock Select

	// TIMSK: Timer Interrupt Mask Register
	TIMSK_ICIE1      = 0x80 // Input Capture Interrupt Enable
	TIMSK_ICIE1_Msk  = 0x80 // Input Capture Interrupt Enable
	TIMSK_OCIE1B     = 0x20 // Output Compare B Match Interrupt Enable
	TIMSK_OCIE1B_Msk = 0x20 // Output Compare B Match Interrupt Enable
	TIMSK_OCIE1A     = 0x10 // Output Compare A Match Interrupt Enable
	TIMSK_OCIE1A_Msk = 0x10 // Output Compare A Match Interrupt Enable
	TIMSK_TOIE0      = 0x1  // Overflow Interrupt Enable
	TIMSK_TOIE1      = 0x8  // Overflow Interrupt Enable
	TIMSK_TOIE_Msk   = 0x9  // Overflow Interrupt Enable
	TIMSK_OCIE0B     = 0x4  // Timer/Counter Output Compare Match B Interrupt Enable
	TIMSK_OCIE0B_Msk = 0x4  // Timer/Counter Output Compare Match B Interrupt Enable
	TIMSK_OCIE0A     = 0x2  // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK_OCIE0A_Msk = 0x2  // Timer/Counter0 Output Compare Match A Interrupt Enable

	// TIFR: Overflow Interrupt Enable
	TIFR_ICF1      = 0x80 // Input Capture Flag
	TIFR_ICF1_Msk  = 0x80 // Input Capture Flag
	TIFR_OCF1B     = 0x20 // Timer Output Compare Flag 1B
	TIFR_OCF1B_Msk = 0x20 // Timer Output Compare Flag 1B
	TIFR_OCF1A     = 0x10 // Timer Output Compare Flag 1A
	TIFR_OCF1A_Msk = 0x10 // Timer Output Compare Flag 1A
	TIFR_TOV0      = 0x1  // Timer Overflow Flag
	TIFR_TOV1      = 0x8  // Timer Overflow Flag
	TIFR_TOV_Msk   = 0x9  // Timer Overflow Flag
	TIFR_OCF0B     = 0x4  // Output Compare Flag 0 B
	TIFR_OCF0B_Msk = 0x4  // Output Compare Flag 0 B
	TIFR_OCF0A     = 0x2  // Output Compare Flag 0 A
	TIFR_OCF0A_Msk = 0x2  // Output Compare Flag 0 A

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM     = 0x80 // Timer Synchronization Mode
	GTCCR_TSM_Msk = 0x80 // Timer Synchronization Mode
	GTCCR_PSR     = 0x1  // Prescaler Reset
	GTCCR_PSR_Msk = 0x1  // Prescaler Reset
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWSCRA: TWI Slave Control Register A
	TWSCRA_TWSHE      = 0x80 // TWI SDA Hold Time Enable
	TWSCRA_TWSHE_Msk  = 0x80 // TWI SDA Hold Time Enable
	TWSCRA_TWDIE      = 0x20 // TWI Data Interrupt Enable
	TWSCRA_TWDIE_Msk  = 0x20 // TWI Data Interrupt Enable
	TWSCRA_TWASIE     = 0x10 // TWI Address/Stop Interrupt Enable
	TWSCRA_TWASIE_Msk = 0x10 // TWI Address/Stop Interrupt Enable
	TWSCRA_TWEN       = 0x8  // Two-Wire Interface Enable
	TWSCRA_TWEN_Msk   = 0x8  // Two-Wire Interface Enable
	TWSCRA_TWSIE      = 0x4  // TWI Stop Interrupt Enable
	TWSCRA_TWSIE_Msk  = 0x4  // TWI Stop Interrupt Enable
	TWSCRA_TWPME      = 0x2  // TWI Promiscuous Mode Enable
	TWSCRA_TWPME_Msk  = 0x2  // TWI Promiscuous Mode Enable
	TWSCRA_TWSME      = 0x1  // TWI Smart Mode Enable
	TWSCRA_TWSME_Msk  = 0x1  // TWI Smart Mode Enable

	// TWSCRB: TWI Slave Control Register B
	TWSCRB_TWAA      = 0x4 // TWI Acknowledge Action
	TWSCRB_TWAA_Msk  = 0x4 // TWI Acknowledge Action
	TWSCRB_TWCMD0    = 0x1
	TWSCRB_TWCMD1    = 0x2
	TWSCRB_TWCMD_Msk = 0x3

	// TWSA: TWI Slave Address Register
	TWSA_TWSA0    = 0x1  // TWI slave address bit
	TWSA_TWSA1    = 0x2  // TWI slave address bit
	TWSA_TWSA2    = 0x4  // TWI slave address bit
	TWSA_TWSA3    = 0x8  // TWI slave address bit
	TWSA_TWSA4    = 0x10 // TWI slave address bit
	TWSA_TWSA5    = 0x20 // TWI slave address bit
	TWSA_TWSA6    = 0x40 // TWI slave address bit
	TWSA_TWSA7    = 0x80 // TWI slave address bit
	TWSA_TWSA_Msk = 0xff // TWI slave address bit

	// TWSD: TWI Slave Data Register
	TWSD_TWSD0    = 0x1  // TWI slave data bit
	TWSD_TWSD1    = 0x2  // TWI slave data bit
	TWSD_TWSD2    = 0x4  // TWI slave data bit
	TWSD_TWSD3    = 0x8  // TWI slave data bit
	TWSD_TWSD4    = 0x10 // TWI slave data bit
	TWSD_TWSD5    = 0x20 // TWI slave data bit
	TWSD_TWSD6    = 0x40 // TWI slave data bit
	TWSD_TWSD7    = 0x80 // TWI slave data bit
	TWSD_TWSD_Msk = 0xff // TWI slave data bit
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0    = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1    = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1A_Msk = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0    = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1    = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_COM1B_Msk = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM10     = 0x1  // Waveform Generation Mode Bits
	TCCR1A_WGM11     = 0x2  // Waveform Generation Mode Bits
	TCCR1A_WGM1_Msk  = 0x3  // Waveform Generation Mode Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1     = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICNC1_Msk = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1     = 0x40 // Input Capture 1 Edge Select
	TCCR1B_ICES1_Msk = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10     = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11     = 0x10 // Waveform Generation Mode
	TCCR1B_WGM1_Msk  = 0x18 // Waveform Generation Mode
	TCCR1B_CS10      = 0x1  // Clock Select1 bits
	TCCR1B_CS11      = 0x2  // Clock Select1 bits
	TCCR1B_CS12      = 0x4  // Clock Select1 bits
	TCCR1B_CS1_Msk   = 0x7  // Clock Select1 bits

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A     = 0x80 // Force Output Compare for channel A
	TCCR1C_FOC1A_Msk = 0x80 // Force Output Compare for channel A
	TCCR1C_FOC1B     = 0x40 // Force Output Compare for channel B
	TCCR1C_FOC1B_Msk = 0x40 // Force Output Compare for channel B
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE     = 0x80 // SPI Interrupt Enable
	SPCR_SPIE_Msk = 0x80 // SPI Interrupt Enable
	SPCR_SPE      = 0x40 // SPI Enable
	SPCR_SPE_Msk  = 0x40 // SPI Enable
	SPCR_DORD     = 0x20 // Data Order
	SPCR_DORD_Msk = 0x20 // Data Order
	SPCR_MSTR     = 0x10 // Master/Slave Select
	SPCR_MSTR_Msk = 0x10 // Master/Slave Select
	SPCR_CPOL     = 0x8  // Clock polarity
	SPCR_CPOL_Msk = 0x8  // Clock polarity
	SPCR_CPHA     = 0x4  // Clock Phase
	SPCR_CPHA_Msk = 0x4  // Clock Phase
	SPCR_SPR0     = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1     = 0x2  // SPI Clock Rate Selects
	SPCR_SPR_Msk  = 0x3  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF      = 0x80 // SPI Interrupt Flag
	SPSR_SPIF_Msk  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL      = 0x40 // Write Collision Flag
	SPSR_WCOL_Msk  = 0x40 // Write Collision Flag
	SPSR_SPI2X     = 0x1  // Double SPI Speed Bit
	SPSR_SPI2X_Msk = 0x1  // Double SPI Speed Bit
)
