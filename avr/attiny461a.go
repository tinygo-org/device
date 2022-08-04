// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny461A.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny461a
// +build avr,attiny461a

// Device information for the ATtiny461A.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny461A"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET            = 0  // External Reset, Power-on Reset and Watchdog Reset
	IRQ_INT0             = 1  // External Interrupt 0
	IRQ_PCINT            = 2  // Pin Change Interrupt
	IRQ_TIMER1_COMPA     = 3  // Timer/Counter1 Compare Match 1A
	IRQ_TIMER1_COMPB     = 4  // Timer/Counter1 Compare Match 1B
	IRQ_TIMER1_OVF       = 5  // Timer/Counter1 Overflow
	IRQ_TIMER0_OVF       = 6  // Timer/Counter0 Overflow
	IRQ_USI_START        = 7  // USI Start
	IRQ_USI_OVF          = 8  // USI Overflow
	IRQ_EE_RDY           = 9  // EEPROM Ready
	IRQ_ANA_COMP         = 10 // Analog Comparator
	IRQ_ADC              = 11 // ADC Conversion Complete
	IRQ_WDT              = 12 // Watchdog Time-Out
	IRQ_INT1             = 13 // External Interrupt 1
	IRQ_TIMER0_COMPA     = 14 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB     = 15 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_CAPT      = 16 // ADC Conversion Complete
	IRQ_TIMER1_COMPD     = 17 // Timer/Counter1 Compare Match D
	IRQ_FAULT_PROTECTION = 18 // Timer/Counter1 Fault Protection
	IRQ_max              = 18 // Highest interrupt number on this device.
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

//export __vector_PCINT
//go:interrupt
func interruptPCINT() {
	callHandlers(IRQ_PCINT)
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

//export __vector_TIMER0_OVF
//go:interrupt
func interruptTIMER0_OVF() {
	callHandlers(IRQ_TIMER0_OVF)
}

//export __vector_USI_START
//go:interrupt
func interruptUSI_START() {
	callHandlers(IRQ_USI_START)
}

//export __vector_USI_OVF
//go:interrupt
func interruptUSI_OVF() {
	callHandlers(IRQ_USI_OVF)
}

//export __vector_EE_RDY
//go:interrupt
func interruptEE_RDY() {
	callHandlers(IRQ_EE_RDY)
}

//export __vector_ANA_COMP
//go:interrupt
func interruptANA_COMP() {
	callHandlers(IRQ_ANA_COMP)
}

//export __vector_ADC
//go:interrupt
func interruptADC() {
	callHandlers(IRQ_ADC)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_INT1
//go:interrupt
func interruptINT1() {
	callHandlers(IRQ_INT1)
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

//export __vector_TIMER0_CAPT
//go:interrupt
func interruptTIMER0_CAPT() {
	callHandlers(IRQ_TIMER0_CAPT)
}

//export __vector_TIMER1_COMPD
//go:interrupt
func interruptTIMER1_COMPD() {
	callHandlers(IRQ_TIMER1_COMPD)
}

//export __vector_FAULT_PROTECTION
//go:interrupt
func interruptFAULT_PROTECTION() {
	callHandlers(IRQ_FAULT_PROTECTION)
}

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// I/O Port
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	DIDR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))

	// Analog Comparator
	ACSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	ACSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))

	// Universal Serial Interface
	USIPP = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	USIBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	USIDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	USISR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	USICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// Watchdog Timer
	WDTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Timer/Counter, 8-bit
	TIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	TIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	TCNT0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TCCR1D = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	TCCR1E = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	TCNT1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TC1H   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	OCR1A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	OCR1C  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1D  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	DT1    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))

	// CPU Registers
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	PLLCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	DWDR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))

	// External Interrupts
	GIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_SELFPRGEN     = 0x1 // Self Programming enable
	EXTENDED_SELFPRGEN_Msk = 0x1 // Self Programming enable

	// HIGH
	HIGH_RSTDISBL     = 0x80 // Reset Disabled (Enable PB7 as i/o pin)
	HIGH_RSTDISBL_Msk = 0x80 // Reset Disabled (Enable PB7 as i/o pin)
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
	LOW_CKOUT         = 0x40 // Clock output on PORTB5
	LOW_CKOUT_Msk     = 0x40 // Clock output on PORTB5
	LOW_SUT_CKSEL0    = 0x1  // Select Clock source
	LOW_SUT_CKSEL1    = 0x2  // Select Clock source
	LOW_SUT_CKSEL2    = 0x4  // Select Clock source
	LOW_SUT_CKSEL3    = 0x8  // Select Clock source
	LOW_SUT_CKSEL4    = 0x10 // Select Clock source
	LOW_SUT_CKSEL5    = 0x20 // Select Clock source
	LOW_SUT_CKSEL_Msk = 0x3f // Select Clock source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0    = 0x1 // Memory Lock
	LOCKBIT_LB1    = 0x2 // Memory Lock
	LOCKBIT_LB_Msk = 0x3 // Memory Lock
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0     = 0x40 // Reference Selection Bits
	ADMUX_REFS1     = 0x80 // Reference Selection Bits
	ADMUX_REFS_Msk  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR     = 0x20 // Left Adjust Result
	ADMUX_ADLAR_Msk = 0x20 // Left Adjust Result
	ADMUX_MUX0      = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1      = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2      = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3      = 0x8  // Analog Channel and Gain Selection Bits
	ADMUX_MUX4      = 0x10 // Analog Channel and Gain Selection Bits
	ADMUX_MUX_Msk   = 0x1f // Analog Channel and Gain Selection Bits

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
	ADCSRB_BIN       = 0x80 // Bipolar Input Mode
	ADCSRB_BIN_Msk   = 0x80 // Bipolar Input Mode
	ADCSRB_GSEL      = 0x40 // Gain Select
	ADCSRB_GSEL_Msk  = 0x40 // Gain Select
	ADCSRB_IPR       = 0x20 // Input Polarity Mode
	ADCSRB_IPR_Msk   = 0x20 // Input Polarity Mode
	ADCSRB_REFS2     = 0x10
	ADCSRB_REFS2_Msk = 0x10
	ADCSRB_MUX5      = 0x8
	ADCSRB_MUX5_Msk  = 0x8
	ADCSRB_ADTS0     = 0x1 // ADC Auto Trigger Sources
	ADCSRB_ADTS1     = 0x2 // ADC Auto Trigger Sources
	ADCSRB_ADTS2     = 0x4 // ADC Auto Trigger Sources
	ADCSRB_ADTS_Msk  = 0x7 // ADC Auto Trigger Sources

	// DIDR1: Digital Input Disable Register 1
	DIDR1_ADC10D     = 0x80 // ADC10 Digital input Disable
	DIDR1_ADC10D_Msk = 0x80 // ADC10 Digital input Disable
	DIDR1_ADC9D      = 0x40 // ADC9 Digital input Disable
	DIDR1_ADC9D_Msk  = 0x40 // ADC9 Digital input Disable
	DIDR1_ADC8D      = 0x20 // ADC8 Digital input Disable
	DIDR1_ADC8D_Msk  = 0x20 // ADC8 Digital input Disable
	DIDR1_ADC7D      = 0x10 // ADC7 Digital input Disable
	DIDR1_ADC7D_Msk  = 0x10 // ADC7 Digital input Disable

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC6D     = 0x80 // ADC6 Digital input Disable
	DIDR0_ADC6D_Msk = 0x80 // ADC6 Digital input Disable
	DIDR0_ADC5D     = 0x40 // ADC5 Digital input Disable
	DIDR0_ADC5D_Msk = 0x40 // ADC5 Digital input Disable
	DIDR0_ADC4D     = 0x20 // ADC4 Digital input Disable
	DIDR0_ADC4D_Msk = 0x20 // ADC4 Digital input Disable
	DIDR0_ADC3D     = 0x10 // ADC3 Digital input Disable
	DIDR0_ADC3D_Msk = 0x10 // ADC3 Digital input Disable
	DIDR0_AREFD     = 0x8  // AREF Digital Input Disable
	DIDR0_AREFD_Msk = 0x8  // AREF Digital Input Disable
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
	ACSRB_ACM0     = 0x1  // Analog Comparator Multiplexer
	ACSRB_ACM1     = 0x2  // Analog Comparator Multiplexer
	ACSRB_ACM2     = 0x4  // Analog Comparator Multiplexer
	ACSRB_ACM_Msk  = 0x7  // Analog Comparator Multiplexer

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
	ACSRA_ACME     = 0x4  // Analog Comparator Multiplexer Enable
	ACSRA_ACME_Msk = 0x4  // Analog Comparator Multiplexer Enable
	ACSRA_ACIS0    = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSRA_ACIS1    = 0x2  // Analog Comparator Interrupt Mode Select bits
	ACSRA_ACIS_Msk = 0x3  // Analog Comparator Interrupt Mode Select bits
)

// Bitfields for USI: Universal Serial Interface
const (
	// USISR: USI Status Register
	USISR_USISIF     = 0x80 // Start Condition Interrupt Flag
	USISR_USISIF_Msk = 0x80 // Start Condition Interrupt Flag
	USISR_USIOIF     = 0x40 // Counter Overflow Interrupt Flag
	USISR_USIOIF_Msk = 0x40 // Counter Overflow Interrupt Flag
	USISR_USIPF      = 0x20 // Stop Condition Flag
	USISR_USIPF_Msk  = 0x20 // Stop Condition Flag
	USISR_USIDC      = 0x10 // Data Output Collision
	USISR_USIDC_Msk  = 0x10 // Data Output Collision
	USISR_USICNT0    = 0x1  // USI Counter Value Bits
	USISR_USICNT1    = 0x2  // USI Counter Value Bits
	USISR_USICNT2    = 0x4  // USI Counter Value Bits
	USISR_USICNT3    = 0x8  // USI Counter Value Bits
	USISR_USICNT_Msk = 0xf  // USI Counter Value Bits

	// USICR: USI Control Register
	USICR_USISIE     = 0x80 // Start Condition Interrupt Enable
	USICR_USISIE_Msk = 0x80 // Start Condition Interrupt Enable
	USICR_USIOIE     = 0x40 // Counter Overflow Interrupt Enable
	USICR_USIOIE_Msk = 0x40 // Counter Overflow Interrupt Enable
	USICR_USIWM0     = 0x10 // USI Wire Mode Bits
	USICR_USIWM1     = 0x20 // USI Wire Mode Bits
	USICR_USIWM_Msk  = 0x30 // USI Wire Mode Bits
	USICR_USICS0     = 0x4  // USI Clock Source Select Bits
	USICR_USICS1     = 0x8  // USI Clock Source Select Bits
	USICR_USICS_Msk  = 0xc  // USI Clock Source Select Bits
	USICR_USICLK     = 0x2  // Clock Strobe
	USICR_USICLK_Msk = 0x2  // Clock Strobe
	USICR_USITC      = 0x1  // Toggle Clock Port Pin
	USICR_USITC_Msk  = 0x1  // Toggle Clock Port Pin
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

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDIF     = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCR_WDIF_Msk = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCR_WDIE     = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCR_WDIE_Msk = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCR_WDP0     = 0x1  // Watchdog Timer Prescaler Bits
	WDTCR_WDP1     = 0x2  // Watchdog Timer Prescaler Bits
	WDTCR_WDP2     = 0x4  // Watchdog Timer Prescaler Bits
	WDTCR_WDP3     = 0x20 // Watchdog Timer Prescaler Bits
	WDTCR_WDP_Msk  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCR_WDCE     = 0x10 // Watchdog Change Enable
	WDTCR_WDCE_Msk = 0x10 // Watchdog Change Enable
	WDTCR_WDE      = 0x8  // Watch Dog Enable
	WDTCR_WDE_Msk  = 0x8  // Watch Dog Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TIMSK: Timer/Counter Interrupt Mask Register
	TIMSK_OCIE0A     = 0x10 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK_OCIE0A_Msk = 0x10 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK_OCIE0B     = 0x8  // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK_OCIE0B_Msk = 0x8  // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK_TOIE0      = 0x2  // Timer/Counter0 Overflow Interrupt Enable
	TIMSK_TOIE0_Msk  = 0x2  // Timer/Counter0 Overflow Interrupt Enable
	TIMSK_TICIE0     = 0x1  // Timer/Counter0 Input Capture Interrupt Enable
	TIMSK_TICIE0_Msk = 0x1  // Timer/Counter0 Input Capture Interrupt Enable
	TIMSK_OCIE1D     = 0x80 // OCIE1D: Timer/Counter1 Output Compare Interrupt Enable
	TIMSK_OCIE1D_Msk = 0x80 // OCIE1D: Timer/Counter1 Output Compare Interrupt Enable
	TIMSK_OCIE1A     = 0x40 // OCIE1A: Timer/Counter1 Output Compare Interrupt Enable
	TIMSK_OCIE1A_Msk = 0x40 // OCIE1A: Timer/Counter1 Output Compare Interrupt Enable
	TIMSK_OCIE1B     = 0x20 // OCIE1A: Timer/Counter1 Output Compare B Interrupt Enable
	TIMSK_OCIE1B_Msk = 0x20 // OCIE1A: Timer/Counter1 Output Compare B Interrupt Enable
	TIMSK_TOIE1      = 0x4  // Timer/Counter1 Overflow Interrupt Enable
	TIMSK_TOIE1_Msk  = 0x4  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR: Timer/Counter0 Interrupt Flag register
	TIFR_OCF0A     = 0x10 // Timer/Counter0 Output Compare Flag 0A
	TIFR_OCF0A_Msk = 0x10 // Timer/Counter0 Output Compare Flag 0A
	TIFR_OCF0B     = 0x8  // Timer/Counter0 Output Compare Flag 0B
	TIFR_OCF0B_Msk = 0x8  // Timer/Counter0 Output Compare Flag 0B
	TIFR_TOV0      = 0x2  // Timer/Counter0 Overflow Flag
	TIFR_TOV0_Msk  = 0x2  // Timer/Counter0 Overflow Flag
	TIFR_ICF0      = 0x1  // Timer/Counter0 Input Capture Flag
	TIFR_ICF0_Msk  = 0x1  // Timer/Counter0 Input Capture Flag
	TIFR_OCF1D     = 0x80 // Timer/Counter1 Output Compare Flag 1D
	TIFR_OCF1D_Msk = 0x80 // Timer/Counter1 Output Compare Flag 1D
	TIFR_OCF1A     = 0x40 // Timer/Counter1 Output Compare Flag 1A
	TIFR_OCF1A_Msk = 0x40 // Timer/Counter1 Output Compare Flag 1A
	TIFR_OCF1B     = 0x20 // Timer/Counter1 Output Compare Flag 1B
	TIFR_OCF1B_Msk = 0x20 // Timer/Counter1 Output Compare Flag 1B
	TIFR_TOV1      = 0x4  // Timer/Counter1 Overflow Flag
	TIFR_TOV1_Msk  = 0x4  // Timer/Counter1 Overflow Flag

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_TCW0      = 0x80 // Timer/Counter 0 Width
	TCCR0A_TCW0_Msk  = 0x80 // Timer/Counter 0 Width
	TCCR0A_ICEN0     = 0x40 // Input Capture Mode Enable
	TCCR0A_ICEN0_Msk = 0x40 // Input Capture Mode Enable
	TCCR0A_ICNC0     = 0x20 // Input Capture Noice Canceler
	TCCR0A_ICNC0_Msk = 0x20 // Input Capture Noice Canceler
	TCCR0A_ICES0     = 0x10 // Input Capture Edge Select
	TCCR0A_ICES0_Msk = 0x10 // Input Capture Edge Select
	TCCR0A_ACIC0     = 0x8  // Analog Comparator Input Capture Enable
	TCCR0A_ACIC0_Msk = 0x8  // Analog Comparator Input Capture Enable
	TCCR0A_WGM00     = 0x1  // Waveform Generation Mode
	TCCR0A_WGM00_Msk = 0x1  // Waveform Generation Mode

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_TSM      = 0x10 // Timer/Counter Synchronization Mode
	TCCR0B_TSM_Msk  = 0x10 // Timer/Counter Synchronization Mode
	TCCR0B_PSR0     = 0x8  // Timer/Counter 0 Prescaler Reset
	TCCR0B_PSR0_Msk = 0x8  // Timer/Counter 0 Prescaler Reset
	TCCR0B_CS00     = 0x1  // Clock Select
	TCCR0B_CS01     = 0x2  // Clock Select
	TCCR0B_CS02     = 0x4  // Clock Select
	TCCR0B_CS0_Msk  = 0x7  // Clock Select

	// TCCR1A: Timer/Counter Control Register A
	TCCR1A_COM1A0    = 0x40 // Compare Output Mode, Bits
	TCCR1A_COM1A1    = 0x80 // Compare Output Mode, Bits
	TCCR1A_COM1A_Msk = 0xc0 // Compare Output Mode, Bits
	TCCR1A_COM1B0    = 0x10 // Compare Output Mode, Bits
	TCCR1A_COM1B1    = 0x20 // Compare Output Mode, Bits
	TCCR1A_COM1B_Msk = 0x30 // Compare Output Mode, Bits
	TCCR1A_FOC1A     = 0x8  // Force Output Compare Match 1A
	TCCR1A_FOC1A_Msk = 0x8  // Force Output Compare Match 1A
	TCCR1A_FOC1B     = 0x4  // Force Output Compare Match 1B
	TCCR1A_FOC1B_Msk = 0x4  // Force Output Compare Match 1B
	TCCR1A_PWM1A     = 0x2  // Pulse Width Modulator Enable
	TCCR1A_PWM1A_Msk = 0x2  // Pulse Width Modulator Enable
	TCCR1A_PWM1B     = 0x1  // Pulse Width Modulator Enable
	TCCR1A_PWM1B_Msk = 0x1  // Pulse Width Modulator Enable

	// TCCR1B: Timer/Counter Control Register B
	TCCR1B_PWM1X     = 0x80 // PWM Inversion Mode
	TCCR1B_PWM1X_Msk = 0x80 // PWM Inversion Mode
	TCCR1B_PSR1      = 0x40 // Timer/Counter 1 Prescaler reset
	TCCR1B_PSR1_Msk  = 0x40 // Timer/Counter 1 Prescaler reset
	TCCR1B_DTPS10    = 0x10 // Dead Time Prescaler
	TCCR1B_DTPS11    = 0x20 // Dead Time Prescaler
	TCCR1B_DTPS1_Msk = 0x30 // Dead Time Prescaler
	TCCR1B_CS10      = 0x1  // Clock Select Bits
	TCCR1B_CS11      = 0x2  // Clock Select Bits
	TCCR1B_CS12      = 0x4  // Clock Select Bits
	TCCR1B_CS13      = 0x8  // Clock Select Bits
	TCCR1B_CS1_Msk   = 0xf  // Clock Select Bits

	// TCCR1C: Timer/Counter Control Register C
	TCCR1C_COM1A1S     = 0x80 // COM1A1 Shadow Bit
	TCCR1C_COM1A1S_Msk = 0x80 // COM1A1 Shadow Bit
	TCCR1C_COM1A0S     = 0x40 // COM1A0 Shadow Bit
	TCCR1C_COM1A0S_Msk = 0x40 // COM1A0 Shadow Bit
	TCCR1C_COM1B1S     = 0x20 // COM1B1 Shadow Bit
	TCCR1C_COM1B1S_Msk = 0x20 // COM1B1 Shadow Bit
	TCCR1C_COM1B0S     = 0x10 // COM1B0 Shadow Bit
	TCCR1C_COM1B0S_Msk = 0x10 // COM1B0 Shadow Bit
	TCCR1C_COM1D0      = 0x4  // Comparator D output mode
	TCCR1C_COM1D1      = 0x8  // Comparator D output mode
	TCCR1C_COM1D_Msk   = 0xc  // Comparator D output mode
	TCCR1C_FOC1D       = 0x2  // Force Output Compare Match 1D
	TCCR1C_FOC1D_Msk   = 0x2  // Force Output Compare Match 1D
	TCCR1C_PWM1D       = 0x1  // Pulse Width Modulator D Enable
	TCCR1C_PWM1D_Msk   = 0x1  // Pulse Width Modulator D Enable

	// TCCR1D: Timer/Counter Control Register D
	TCCR1D_FPIE1     = 0x80 // Fault Protection Interrupt Enable
	TCCR1D_FPIE1_Msk = 0x80 // Fault Protection Interrupt Enable
	TCCR1D_FPEN1     = 0x40 // Fault Protection Mode Enable
	TCCR1D_FPEN1_Msk = 0x40 // Fault Protection Mode Enable
	TCCR1D_FPNC1     = 0x20 // Fault Protection Noise Canceler
	TCCR1D_FPNC1_Msk = 0x20 // Fault Protection Noise Canceler
	TCCR1D_FPES1     = 0x10 // Fault Protection Edge Select
	TCCR1D_FPES1_Msk = 0x10 // Fault Protection Edge Select
	TCCR1D_FPAC1     = 0x8  // Fault Protection Analog Comparator Enable
	TCCR1D_FPAC1_Msk = 0x8  // Fault Protection Analog Comparator Enable
	TCCR1D_FPF1      = 0x4  // Fault Protection Interrupt Flag
	TCCR1D_FPF1_Msk  = 0x4  // Fault Protection Interrupt Flag
	TCCR1D_WGM10     = 0x1  // Waveform Generation Mode Bit
	TCCR1D_WGM11     = 0x2  // Waveform Generation Mode Bit
	TCCR1D_WGM1_Msk  = 0x3  // Waveform Generation Mode Bit

	// TCCR1E: Timer/Counter1 Control Register E
	TCCR1E_OC1OE0    = 0x1  // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE1    = 0x2  // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE2    = 0x4  // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE3    = 0x8  // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE4    = 0x10 // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE5    = 0x20 // Ouput Compare Override Enable Bits
	TCCR1E_OC1OE_Msk = 0x3f // Ouput Compare Override Enable Bits

	// DT1: Timer/Counter 1 Dead Time Value
	DT1_DT1H0    = 0x10
	DT1_DT1H1    = 0x20
	DT1_DT1H2    = 0x40
	DT1_DT1H3    = 0x80
	DT1_DT1H_Msk = 0xf0
	DT1_DT1L0    = 0x1
	DT1_DT1L1    = 0x2
	DT1_DT1L2    = 0x4
	DT1_DT1L3    = 0x8
	DT1_DT1L_Msk = 0xf
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_CTPB      = 0x10 // Clear temporary page buffer
	SPMCSR_CTPB_Msk  = 0x10 // Clear temporary page buffer
	SPMCSR_RFLB      = 0x8  // Read fuse and lock bits
	SPMCSR_RFLB_Msk  = 0x8  // Read fuse and lock bits
	SPMCSR_PGWRT     = 0x4  // Page Write
	SPMCSR_PGWRT_Msk = 0x4  // Page Write
	SPMCSR_PGERS     = 0x2  // Page Erase
	SPMCSR_PGERS_Msk = 0x2  // Page Erase
	SPMCSR_SPMEN     = 0x1  // Store Program Memory Enable
	SPMCSR_SPMEN_Msk = 0x1  // Store Program Memory Enable
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

	// PRR: Power Reduction Register
	PRR_PRTIM1     = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM1_Msk = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM0     = 0x4 // Power Reduction Timer/Counter0
	PRR_PRTIM0_Msk = 0x4 // Power Reduction Timer/Counter0
	PRR_PRUSI      = 0x2 // Power Reduction USI
	PRR_PRUSI_Msk  = 0x2 // Power Reduction USI
	PRR_PRADC      = 0x1 // Power Reduction ADC
	PRR_PRADC_Msk  = 0x1 // Power Reduction ADC

	// MCUCR: MCU Control Register
	MCUCR_BODS      = 0x80 // BOD Sleep
	MCUCR_BODS_Msk  = 0x80 // BOD Sleep
	MCUCR_PUD       = 0x40 // Pull-up Disable
	MCUCR_PUD_Msk   = 0x40 // Pull-up Disable
	MCUCR_SE        = 0x20 // Sleep Enable
	MCUCR_SE_Msk    = 0x20 // Sleep Enable
	MCUCR_SM0       = 0x8  // Sleep Mode Select Bits
	MCUCR_SM1       = 0x10 // Sleep Mode Select Bits
	MCUCR_SM_Msk    = 0x18 // Sleep Mode Select Bits
	MCUCR_BODSE     = 0x4  // BOD Sleep Enable
	MCUCR_BODSE_Msk = 0x4  // BOD Sleep Enable
	MCUCR_ISC00     = 0x1  // Interrupt Sense Control 0 bits
	MCUCR_ISC01     = 0x2  // Interrupt Sense Control 0 bits
	MCUCR_ISC0_Msk  = 0x3  // Interrupt Sense Control 0 bits
	MCUCR_ISC01     = 0x2  // Interrupt Sense Control 0 Bit 1
	MCUCR_ISC01_Msk = 0x2  // Interrupt Sense Control 0 Bit 1
	MCUCR_ISC00     = 0x1  // Interrupt Sense Control 0 Bit 0
	MCUCR_ISC00_Msk = 0x1  // Interrupt Sense Control 0 Bit 0

	// MCUSR: MCU Status register
	MCUSR_WDRF      = 0x8 // Watchdog Reset Flag
	MCUSR_WDRF_Msk  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF      = 0x4 // Brown-out Reset Flag
	MCUSR_BORF_Msk  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF     = 0x2 // External Reset Flag
	MCUSR_EXTRF_Msk = 0x2 // External Reset Flag
	MCUSR_PORF      = 0x1 // Power-On Reset Flag
	MCUSR_PORF_Msk  = 0x1 // Power-On Reset Flag

	// OSCCAL: Oscillator Calibration Register
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

	// PLLCSR: PLL Control and status register
	PLLCSR_LSM       = 0x80 // Low speed mode
	PLLCSR_LSM_Msk   = 0x80 // Low speed mode
	PLLCSR_PCKE      = 0x4  // PCK Enable
	PLLCSR_PCKE_Msk  = 0x4  // PCK Enable
	PLLCSR_PLLE      = 0x2  // PLL Enable
	PLLCSR_PLLE_Msk  = 0x2  // PLL Enable
	PLLCSR_PLOCK     = 0x1  // PLL Lock detector
	PLLCSR_PLOCK_Msk = 0x1  // PLL Lock detector
)

// Bitfields for EXINT: External Interrupts
const (
	// GIMSK: General Interrupt Mask Register
	GIMSK_INT0     = 0x40 // External Interrupt Request 1 Enable
	GIMSK_INT1     = 0x80 // External Interrupt Request 1 Enable
	GIMSK_INT_Msk  = 0xc0 // External Interrupt Request 1 Enable
	GIMSK_PCIE0    = 0x10 // Pin Change Interrupt Enables
	GIMSK_PCIE1    = 0x20 // Pin Change Interrupt Enables
	GIMSK_PCIE_Msk = 0x30 // Pin Change Interrupt Enables

	// GIFR: General Interrupt Flag register
	GIFR_INTF0    = 0x40 // External Interrupt Flags
	GIFR_INTF1    = 0x80 // External Interrupt Flags
	GIFR_INTF_Msk = 0xc0 // External Interrupt Flags
	GIFR_PCIF     = 0x20 // Pin Change Interrupt Flag
	GIFR_PCIF_Msk = 0x20 // Pin Change Interrupt Flag
)
