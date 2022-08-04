// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny88.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny88
// +build avr,attiny88

// Device information for the ATtiny88.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny88"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_PCINT0       = 3  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 4  // Pin Change Interrupt Request 1
	IRQ_PCINT2       = 5  // Pin Change Interrupt Request 2
	IRQ_PCINT3       = 6  // Pin Change Interrupt Request 3
	IRQ_WDT          = 7  // Watchdog Time-out Interrupt
	IRQ_TIMER1_CAPT  = 8  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 9  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 10 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 11 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 12 // TimerCounter0 Compare Match A
	IRQ_TIMER0_COMPB = 13 // TimerCounter0 Compare Match B
	IRQ_TIMER0_OVF   = 14 // Timer/Couner0 Overflow
	IRQ_SPI_STC      = 15 // SPI Serial Transfer Complete
	IRQ_ADC          = 16 // ADC Conversion Complete
	IRQ_EE_RDY       = 17 // EEPROM Ready
	IRQ_ANALOG_COMP  = 18 // Analog Comparator
	IRQ_TWI          = 19 // Two-wire Serial Interface
	IRQ_max          = 19 // Highest interrupt number on this device.
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

//export __vector_INT1
//go:interrupt
func interruptINT1() {
	callHandlers(IRQ_INT1)
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

//export __vector_PCINT2
//go:interrupt
func interruptPCINT2() {
	callHandlers(IRQ_PCINT2)
}

//export __vector_PCINT3
//go:interrupt
func interruptPCINT3() {
	callHandlers(IRQ_PCINT3)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_TIMER1_CAPT
//go:interrupt
func interruptTIMER1_CAPT() {
	callHandlers(IRQ_TIMER1_CAPT)
}

//export __vector_TIMER1_COMPA
//go:interrupt
func interruptTIMER1_COMPA() {
	callHandlers(IRQ_TIMER1_COMPA)
}

//export __vector_TIMER1_COMPB
//go:interrupt
func interruptTIMER1_COMPB() {
	callHandlers(IRQ_TIMER1_COMPB)
}

//export __vector_TIMER1_OVF
//go:interrupt
func interruptTIMER1_OVF() {
	callHandlers(IRQ_TIMER1_OVF)
}

//export __vector_TIMER0_COMPA
//go:interrupt
func interruptTIMER0_COMPA() {
	callHandlers(IRQ_TIMER0_COMPA)
}

//export __vector_TIMER0_COMPB
//go:interrupt
func interruptTIMER0_COMPB() {
	callHandlers(IRQ_TIMER0_COMPB)
}

//export __vector_TIMER0_OVF
//go:interrupt
func interruptTIMER0_OVF() {
	callHandlers(IRQ_TIMER0_OVF)
}

//export __vector_SPI_STC
//go:interrupt
func interruptSPI_STC() {
	callHandlers(IRQ_SPI_STC)
}

//export __vector_ADC
//go:interrupt
func interruptADC() {
	callHandlers(IRQ_ADC)
}

//export __vector_EE_RDY
//go:interrupt
func interruptEE_RDY() {
	callHandlers(IRQ_EE_RDY)
}

//export __vector_ANALOG_COMP
//go:interrupt
func interruptANALOG_COMP() {
	callHandlers(IRQ_ANALOG_COMP)
}

//export __vector_TWI
//go:interrupt
func interruptTWI() {
	callHandlers(IRQ_TWI)
}

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Timer/Counter, 16-bit
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8b)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))
	GTCCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))

	// Analog Comparator
	ACSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	DIDR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7f)))

	// I/O Port
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))

	// Serial Peripheral Interface
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))

	// Two Wire Serial Interface
	TWHSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbe)))
	TWAMR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbd)))
	TWBR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb8)))
	TWCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbc)))
	TWSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb9)))
	TWDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbb)))
	TWAR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xba)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7c)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7b)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7e)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	PCICR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))
	PCMSK3 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6a)))
	PCMSK2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6d)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6b)))
	PCIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))

	// Timer/Counter, 8-bit
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))

	// CPU Registers
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	SMCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	PORTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_SELFPRGEN     = 0x1 // Self Programming enable
	EXTENDED_SELFPRGEN_Msk = 0x1 // Self Programming enable

	// HIGH
	HIGH_RSTDISBL     = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_RSTDISBL_Msk = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_DWEN         = 0x40 // Debug Wire enable
	HIGH_DWEN_Msk     = 0x40 // Debug Wire enable
	HIGH_SPIEN        = 0x20 // Serial program downloading (SPI) enabled
	HIGH_SPIEN_Msk    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON        = 0x10 // Watch-dog Timer always on
	HIGH_WDTON_Msk    = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE       = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_EESAVE_Msk   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BODLEVEL0    = 0x1  // Brown-out Detector trigger level
	HIGH_BODLEVEL1    = 0x2  // Brown-out Detector trigger level
	HIGH_BODLEVEL2    = 0x4  // Brown-out Detector trigger level
	HIGH_BODLEVEL_Msk = 0x7  // Brown-out Detector trigger level

	// LOW
	LOW_CKDIV8        = 0x80 // Divide clock by 8 internally
	LOW_CKDIV8_Msk    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT         = 0x40 // Clock output on PORTB0
	LOW_CKOUT_Msk     = 0x40 // Clock output on PORTB0
	LOW_SUT_CKSEL0    = 0x1  // Select Clock Source
	LOW_SUT_CKSEL1    = 0x2  // Select Clock Source
	LOW_SUT_CKSEL2    = 0x4  // Select Clock Source
	LOW_SUT_CKSEL3    = 0x8  // Select Clock Source
	LOW_SUT_CKSEL4    = 0x10 // Select Clock Source
	LOW_SUT_CKSEL5    = 0x20 // Select Clock Source
	LOW_SUT_CKSEL_Msk = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0    = 0x1 // Memory Lock
	LOCKBIT_LB1    = 0x2 // Memory Lock
	LOCKBIT_LB_Msk = 0x3 // Memory Lock
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1      = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_ICIE1_Msk  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B     = 0x4  // Timer/Counter1 Output CompareB Match Interrupt Enable
	TIMSK1_OCIE1B_Msk = 0x4  // Timer/Counter1 Output CompareB Match Interrupt Enable
	TIMSK1_OCIE1A     = 0x2  // Timer/Counter1 Output CompareA Match Interrupt Enable
	TIMSK1_OCIE1A_Msk = 0x2  // Timer/Counter1 Output CompareA Match Interrupt Enable
	TIMSK1_TOIE1      = 0x1  // Timer/Counter1 Overflow Interrupt Enable
	TIMSK1_TOIE1_Msk  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1      = 0x20 // Input Capture Flag 1
	TIFR1_ICF1_Msk  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B     = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1B_Msk = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A     = 0x2  // Output Compare Flag 1A
	TIFR1_OCF1A_Msk = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1      = 0x1  // Timer/Counter1 Overflow Flag
	TIFR1_TOV1_Msk  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0    = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1    = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1A_Msk = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0    = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1    = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_COM1B_Msk = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM10     = 0x1  // Waveform Generation Mode
	TCCR1A_WGM11     = 0x2  // Waveform Generation Mode
	TCCR1A_WGM1_Msk  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1     = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICNC1_Msk = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1     = 0x40 // Input Capture 1 Edge Select
	TCCR1B_ICES1_Msk = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10     = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11     = 0x10 // Waveform Generation Mode
	TCCR1B_WGM1_Msk  = 0x18 // Waveform Generation Mode
	TCCR1B_CS10      = 0x1  // Prescaler source of Timer/Counter 1
	TCCR1B_CS11      = 0x2  // Prescaler source of Timer/Counter 1
	TCCR1B_CS12      = 0x4  // Prescaler source of Timer/Counter 1
	TCCR1B_CS1_Msk   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A     = 0x80
	TCCR1C_FOC1A_Msk = 0x80
	TCCR1C_FOC1B     = 0x40
	TCCR1C_FOC1B_Msk = 0x40

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM         = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_TSM_Msk     = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_PSRSYNC     = 0x1  // Prescaler Reset Timer/Counter1 and Timer/Counter0
	GTCCR_PSRSYNC_Msk = 0x1  // Prescaler Reset Timer/Counter1 and Timer/Counter0
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD      = 0x80 // Analog Comparator Disable
	ACSR_ACD_Msk  = 0x80 // Analog Comparator Disable
	ACSR_ACBG     = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACBG_Msk = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO      = 0x20 // Analog Compare Output
	ACSR_ACO_Msk  = 0x20 // Analog Compare Output
	ACSR_ACI      = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACI_Msk  = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE     = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIE_Msk = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC     = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIC_Msk = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS0    = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS1    = 0x2  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS_Msk = 0x3  // Analog Comparator Interrupt Mode Select bits

	// DIDR1: Digital Input Disable Register 1
	DIDR1_AIN1D     = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN1D_Msk = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D     = 0x1 // AIN0 Digital Input Disable
	DIDR1_AIN0D_Msk = 0x1 // AIN0 Digital Input Disable
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPSR: SPI Status Register
	SPSR_SPIF      = 0x80 // SPI Interrupt Flag
	SPSR_SPIF_Msk  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL      = 0x40 // Write Collision Flag
	SPSR_WCOL_Msk  = 0x40 // Write Collision Flag
	SPSR_SPI2X     = 0x1  // Double SPI Speed Bit
	SPSR_SPI2X_Msk = 0x1  // Double SPI Speed Bit

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
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF     = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIF_Msk = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE     = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDIE_Msk = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP0     = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1     = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2     = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3     = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDP_Msk  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE     = 0x10 // Watchdog Change Enable
	WDTCSR_WDCE_Msk = 0x10 // Watchdog Change Enable
	WDTCSR_WDE      = 0x8  // Watch Dog Enable
	WDTCSR_WDE_Msk  = 0x8  // Watch Dog Enable
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWHSR: TWHSR
	TWHSR_TWHS     = 0x1
	TWHSR_TWHS_Msk = 0x1

	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM0    = 0x2
	TWAMR_TWAM1    = 0x4
	TWAMR_TWAM2    = 0x8
	TWAMR_TWAM3    = 0x10
	TWAMR_TWAM4    = 0x20
	TWAMR_TWAM5    = 0x40
	TWAMR_TWAM6    = 0x80
	TWAMR_TWAM_Msk = 0xfe

	// TWCR: TWI Control Register
	TWCR_TWINT     = 0x80 // TWI Interrupt Flag
	TWCR_TWINT_Msk = 0x80 // TWI Interrupt Flag
	TWCR_TWEA      = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWEA_Msk  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA     = 0x20 // TWI Start Condition Bit
	TWCR_TWSTA_Msk = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO     = 0x10 // TWI Stop Condition Bit
	TWCR_TWSTO_Msk = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC      = 0x8  // TWI Write Collition Flag
	TWCR_TWWC_Msk  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN      = 0x4  // TWI Enable Bit
	TWCR_TWEN_Msk  = 0x4  // TWI Enable Bit
	TWCR_TWIE      = 0x1  // TWI Interrupt Enable
	TWCR_TWIE_Msk  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS0     = 0x8  // TWI Status
	TWSR_TWS1     = 0x10 // TWI Status
	TWSR_TWS2     = 0x20 // TWI Status
	TWSR_TWS3     = 0x40 // TWI Status
	TWSR_TWS4     = 0x80 // TWI Status
	TWSR_TWS_Msk  = 0xf8 // TWI Status
	TWSR_TWPS0    = 0x1  // TWI Prescaler
	TWSR_TWPS1    = 0x2  // TWI Prescaler
	TWSR_TWPS_Msk = 0x3  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA0      = 0x2  // TWI (Slave) Address register Bits
	TWAR_TWA1      = 0x4  // TWI (Slave) Address register Bits
	TWAR_TWA2      = 0x8  // TWI (Slave) Address register Bits
	TWAR_TWA3      = 0x10 // TWI (Slave) Address register Bits
	TWAR_TWA4      = 0x20 // TWI (Slave) Address register Bits
	TWAR_TWA5      = 0x40 // TWI (Slave) Address register Bits
	TWAR_TWA6      = 0x80 // TWI (Slave) Address register Bits
	TWAR_TWA_Msk   = 0xfe // TWI (Slave) Address register Bits
	TWAR_TWGCE     = 0x1  // TWI General Call Recognition Enable Bit
	TWAR_TWGCE_Msk = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0     = 0x40 // Reference Selection Bit 0
	ADMUX_REFS0_Msk = 0x40 // Reference Selection Bit 0
	ADMUX_ADLAR     = 0x20 // Left Adjust Result
	ADMUX_ADLAR_Msk = 0x20 // Left Adjust Result
	ADMUX_MUX0      = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1      = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2      = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3      = 0x8  // Analog Channel and Gain Selection Bits
	ADMUX_MUX_Msk   = 0xf  // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register A
	ADCSRA_ADEN      = 0x80 // ADC Enable
	ADCSRA_ADEN_Msk  = 0x80 // ADC Enable
	ADCSRA_ADSC      = 0x40 // ADC Start Conversion
	ADCSRA_ADSC_Msk  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE     = 0x20 // ADC  Auto Trigger Enable
	ADCSRA_ADATE_Msk = 0x20 // ADC  Auto Trigger Enable
	ADCSRA_ADIF      = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIF_Msk  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE      = 0x8  // ADC Interrupt Enable
	ADCSRA_ADIE_Msk  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0     = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1     = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2     = 0x4  // ADC  Prescaler Select Bits
	ADCSRA_ADPS_Msk  = 0x7  // ADC  Prescaler Select Bits

	// ADCSRB: The ADC Control and Status register B
	ADCSRB_ACME     = 0x40
	ADCSRB_ACME_Msk = 0x40
	ADCSRB_ADTS0    = 0x1 // ADC Auto Trigger Source bits
	ADCSRB_ADTS1    = 0x2 // ADC Auto Trigger Source bits
	ADCSRB_ADTS2    = 0x4 // ADC Auto Trigger Source bits
	ADCSRB_ADTS_Msk = 0x7 // ADC Auto Trigger Source bits

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D     = 0x80
	DIDR0_ADC7D_Msk = 0x80
	DIDR0_ADC6D     = 0x40
	DIDR0_ADC6D_Msk = 0x40
	DIDR0_ADC5D     = 0x20
	DIDR0_ADC5D_Msk = 0x20
	DIDR0_ADC4D     = 0x10
	DIDR0_ADC4D_Msk = 0x10
	DIDR0_ADC3D     = 0x8
	DIDR0_ADC3D_Msk = 0x8
	DIDR0_ADC2D     = 0x4
	DIDR0_ADC2D_Msk = 0x4
	DIDR0_ADC1D     = 0x2
	DIDR0_ADC1D_Msk = 0x2
	DIDR0_ADC0D     = 0x1
	DIDR0_ADC0D_Msk = 0x1
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC10    = 0x4 // External Interrupt Sense Control 1 Bits
	EICRA_ISC11    = 0x8 // External Interrupt Sense Control 1 Bits
	EICRA_ISC1_Msk = 0xc // External Interrupt Sense Control 1 Bits
	EICRA_ISC00    = 0x1 // External Interrupt Sense Control 0 Bits
	EICRA_ISC01    = 0x2 // External Interrupt Sense Control 0 Bits
	EICRA_ISC0_Msk = 0x3 // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT0    = 0x1 // External Interrupt Request 1 Enable
	EIMSK_INT1    = 0x2 // External Interrupt Request 1 Enable
	EIMSK_INT_Msk = 0x3 // External Interrupt Request 1 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF0    = 0x1 // External Interrupt Flags
	EIFR_INTF1    = 0x2 // External Interrupt Flags
	EIFR_INTF_Msk = 0x3 // External Interrupt Flags

	// PCICR
	PCICR_PCIE0    = 0x1
	PCICR_PCIE1    = 0x2
	PCICR_PCIE2    = 0x4
	PCICR_PCIE3    = 0x8
	PCICR_PCIE_Msk = 0xf

	// PCMSK3: Pin Change Mask Register 3
	PCMSK3_PCINT0    = 0x1 // Pin Change Enable Masks
	PCMSK3_PCINT1    = 0x2 // Pin Change Enable Masks
	PCMSK3_PCINT2    = 0x4 // Pin Change Enable Masks
	PCMSK3_PCINT3    = 0x8 // Pin Change Enable Masks
	PCMSK3_PCINT_Msk = 0xf // Pin Change Enable Masks

	// PCMSK2: Pin Change Mask Register 2
	PCMSK2_PCINT0    = 0x1  // Pin Change Enable Masks
	PCMSK2_PCINT1    = 0x2  // Pin Change Enable Masks
	PCMSK2_PCINT2    = 0x4  // Pin Change Enable Masks
	PCMSK2_PCINT3    = 0x8  // Pin Change Enable Masks
	PCMSK2_PCINT4    = 0x10 // Pin Change Enable Masks
	PCMSK2_PCINT5    = 0x20 // Pin Change Enable Masks
	PCMSK2_PCINT6    = 0x40 // Pin Change Enable Masks
	PCMSK2_PCINT7    = 0x80 // Pin Change Enable Masks
	PCMSK2_PCINT_Msk = 0xff // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT0    = 0x1  // Pin Change Enable Masks
	PCMSK1_PCINT1    = 0x2  // Pin Change Enable Masks
	PCMSK1_PCINT2    = 0x4  // Pin Change Enable Masks
	PCMSK1_PCINT3    = 0x8  // Pin Change Enable Masks
	PCMSK1_PCINT4    = 0x10 // Pin Change Enable Masks
	PCMSK1_PCINT5    = 0x20 // Pin Change Enable Masks
	PCMSK1_PCINT6    = 0x40 // Pin Change Enable Masks
	PCMSK1_PCINT7    = 0x80 // Pin Change Enable Masks
	PCMSK1_PCINT_Msk = 0xff // Pin Change Enable Masks

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

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0    = 0x1 // Pin Change Interrupt Flags
	PCIFR_PCIF1    = 0x2 // Pin Change Interrupt Flags
	PCIFR_PCIF2    = 0x4 // Pin Change Interrupt Flags
	PCIFR_PCIF3    = 0x8 // Pin Change Interrupt Flags
	PCIFR_PCIF_Msk = 0xf // Pin Change Interrupt Flags
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_CTC0     = 0x8 // Clear Timer on Compare Match
	TCCR0A_CTC0_Msk = 0x8 // Clear Timer on Compare Match
	TCCR0A_CS00     = 0x1 // Clock Select
	TCCR0A_CS01     = 0x2 // Clock Select
	TCCR0A_CS02     = 0x4 // Clock Select
	TCCR0A_CS0_Msk  = 0x7 // Clock Select

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B     = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0B_Msk = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A     = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_OCIE0A_Msk = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0      = 0x1 // Timer/Counter0 Overflow Interrupt Enable
	TIMSK0_TOIE0_Msk  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B     = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0B_Msk = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A     = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_OCF0A_Msk = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0      = 0x1 // Timer/Counter0 Overflow Flag
	TIFR0_TOV0_Msk  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTWI      = 0x80 // Power Reduction TWI
	PRR_PRTWI_Msk  = 0x80 // Power Reduction TWI
	PRR_PRTIM0     = 0x20 // Power Reduction Timer/Counter0
	PRR_PRTIM0_Msk = 0x20 // Power Reduction Timer/Counter0
	PRR_PRTIM1     = 0x8  // Power Reduction Timer/Counter1
	PRR_PRTIM1_Msk = 0x8  // Power Reduction Timer/Counter1
	PRR_PRSPI      = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRSPI_Msk  = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRADC      = 0x1  // Power Reduction ADC
	PRR_PRADC_Msk  = 0x1  // Power Reduction ADC

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

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE     = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPCE_Msk = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS0     = 0x1  // Clock Prescaler Select Bits
	CLKPR_CLKPS1     = 0x2  // Clock Prescaler Select Bits
	CLKPR_CLKPS2     = 0x4  // Clock Prescaler Select Bits
	CLKPR_CLKPS3     = 0x8  // Clock Prescaler Select Bits
	CLKPR_CLKPS_Msk  = 0xf  // Clock Prescaler Select Bits

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

	// SPMCSR: Store Program Memory Control Register
	SPMCSR_RWWSB         = 0x40 // Read-While-Write Section Busy
	SPMCSR_RWWSB_Msk     = 0x40 // Read-While-Write Section Busy
	SPMCSR_CTPB          = 0x10 // Clear Temporary Page Buffer
	SPMCSR_CTPB_Msk      = 0x10 // Clear Temporary Page Buffer
	SPMCSR_RFLB          = 0x8  // Read Fuse and Lock Bits
	SPMCSR_RFLB_Msk      = 0x8  // Read Fuse and Lock Bits
	SPMCSR_PGWRT         = 0x4  // Page Write
	SPMCSR_PGWRT_Msk     = 0x4  // Page Write
	SPMCSR_PGERS         = 0x2  // Page Erase
	SPMCSR_PGERS_Msk     = 0x2  // Page Erase
	SPMCSR_SELFPRGEN     = 0x1  // Self Programming Enable
	SPMCSR_SELFPRGEN_Msk = 0x1  // Self Programming Enable

	// MCUCR: MCU Control Register
	MCUCR_BODS      = 0x40 // BOD Sleep
	MCUCR_BODS_Msk  = 0x40 // BOD Sleep
	MCUCR_BODSE     = 0x20 // BOD Sleep Enable
	MCUCR_BODSE_Msk = 0x20 // BOD Sleep Enable
	MCUCR_PUD       = 0x10 // Pull-up Disable
	MCUCR_PUD_Msk   = 0x10 // Pull-up Disable

	// MCUSR: MCU Status Register
	MCUSR_WDRF      = 0x8 // Watchdog Reset Flag
	MCUSR_WDRF_Msk  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF      = 0x4 // Brown-out Reset Flag
	MCUSR_BORF_Msk  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF     = 0x2 // External Reset Flag
	MCUSR_EXTRF_Msk = 0x2 // External Reset Flag
	MCUSR_PORF      = 0x1 // Power-on reset flag
	MCUSR_PORF_Msk  = 0x1 // Power-on reset flag

	// SMCR: Sleep Mode Control Register
	SMCR_SM0    = 0x2 // Sleep Mode
	SMCR_SM1    = 0x4 // Sleep Mode
	SMCR_SM_Msk = 0x6 // Sleep Mode
	SMCR_SE     = 0x1 // Sleep Enable
	SMCR_SE_Msk = 0x1 // Sleep Enable

	// PORTCR: Port Configuration Register
	PORTCR_BBMD     = 0x80
	PORTCR_BBMD_Msk = 0x80
	PORTCR_BBMC     = 0x40
	PORTCR_BBMC_Msk = 0x40
	PORTCR_BBMB     = 0x20
	PORTCR_BBMB_Msk = 0x20
	PORTCR_BBMA     = 0x10
	PORTCR_BBMA_Msk = 0x10
	PORTCR_PUDD     = 0x8
	PORTCR_PUDD_Msk = 0x8
	PORTCR_PUDC     = 0x4
	PORTCR_PUDC_Msk = 0x4
	PORTCR_PUDB     = 0x2
	PORTCR_PUDB_Msk = 0x2
	PORTCR_PUDA     = 0x1
	PORTCR_PUDA_Msk = 0x1
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM0     = 0x10 // EEPROM Programming Mode Bits
	EECR_EEPM1     = 0x20 // EEPROM Programming Mode Bits
	EECR_EEPM_Msk  = 0x30 // EEPROM Programming Mode Bits
	EECR_EERIE     = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EERIE_Msk = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE     = 0x4  // EEPROM Master Write Enable
	EECR_EEMPE_Msk = 0x4  // EEPROM Master Write Enable
	EECR_EEPE      = 0x2  // EEPROM Write Enable
	EECR_EEPE_Msk  = 0x2  // EEPROM Write Enable
	EECR_EERE      = 0x1  // EEPROM Read Enable
	EECR_EERE_Msk  = 0x1  // EEPROM Read Enable
)