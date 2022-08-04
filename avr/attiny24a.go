// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny24A.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny24a
// +build avr,attiny24a

// Device information for the ATtiny24A.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny24A"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET      = 0  // External Pin, Power-on Reset, Brown-out Reset,Watchdog Reset
	IRQ_EXT_INT0   = 1  // External Interrupt Request 0
	IRQ_PCINT0     = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1     = 3  // Pin Change Interrupt Request 1
	IRQ_WDT        = 4  // Watchdog Time-out
	IRQ_TIM1_CAPT  = 5  // Timer/Counter1 Capture Event
	IRQ_TIM1_COMPA = 6  // Timer/Counter1 Compare Match A
	IRQ_TIM1_COMPB = 7  // Timer/Counter1 Compare Match B
	IRQ_TIM1_OVF   = 8  // Timer/Counter1 Overflow
	IRQ_TIM0_COMPA = 9  // Timer/Counter0 Compare Match A
	IRQ_TIM0_COMPB = 10 // Timer/Counter0 Compare Match B
	IRQ_TIM0_OVF   = 11 // Timer/Counter0 Overflow
	IRQ_ANA_COMP   = 12 // Analog Comparator
	IRQ_ADC        = 13 // ADC Conversion Complete
	IRQ_EE_RDY     = 14 // EEPROM Ready
	IRQ_USI_STR    = 15 // USI START
	IRQ_USI_OVF    = 16 // USI Overflow
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

//export __vector_EXT_INT0
//go:interrupt
func interruptEXT_INT0() {
	callHandlers(IRQ_EXT_INT0)
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

//export __vector_USI_STR
//go:interrupt
func interruptUSI_STR() {
	callHandlers(IRQ_USI_STR)
}

//export __vector_USI_OVF
//go:interrupt
func interruptUSI_OVF() {
	callHandlers(IRQ_USI_OVF)
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

	// Analog Comparator
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	ACSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))

	// Universal Serial Interface
	USIBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	USIDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	USISR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	USICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// External Interrupts
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	GIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Timer/Counter, 8-bit
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5c)))
	GTCCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))

	// Timer/Counter, 16-bit
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))

	// CPU Registers
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_SELFPRGEN     = 0x1 // Self Programming enable
	EXTENDED_SELFPRGEN_Msk = 0x1 // Self Programming enable

	// HIGH
	HIGH_RSTDISBL     = 0x80 // Reset Disabled (Enable PB3 as i/o pin)
	HIGH_RSTDISBL_Msk = 0x80 // Reset Disabled (Enable PB3 as i/o pin)
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
	LOW_CKOUT         = 0x40 // Clock output on PORTB2
	LOW_CKOUT_Msk     = 0x40 // Clock output on PORTB2
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

// Bitfields for AC: Analog Comparator
const (
	// ADCSRB: ADC Control and Status Register B
	ADCSRB_ACME      = 0x40 // Analog Comparator Multiplexer Enable
	ADCSRB_ACME_Msk  = 0x40 // Analog Comparator Multiplexer Enable
	ADCSRB_BIN       = 0x80 // Bipolar Input Mode
	ADCSRB_BIN_Msk   = 0x80 // Bipolar Input Mode
	ADCSRB_ADLAR     = 0x10 // ADC Left Adjust Result
	ADCSRB_ADLAR_Msk = 0x10 // ADC Left Adjust Result
	ADCSRB_ADTS0     = 0x1  // ADC Auto Trigger Source bits
	ADCSRB_ADTS1     = 0x2  // ADC Auto Trigger Source bits
	ADCSRB_ADTS2     = 0x4  // ADC Auto Trigger Source bits
	ADCSRB_ADTS_Msk  = 0x7  // ADC Auto Trigger Source bits

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

	// DIDR0
	DIDR0_ADC1D     = 0x2 // ADC 1 Digital input buffer disable
	DIDR0_ADC1D_Msk = 0x2 // ADC 1 Digital input buffer disable
	DIDR0_ADC0D     = 0x1 // ADC 0 Digital input buffer disable
	DIDR0_ADC0D_Msk = 0x1 // ADC 0 Digital input buffer disable
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADCSRA: ADC Control and Status Register A
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
	ADCSRA_ADPS0     = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1     = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2     = 0x4  // ADC  Prescaler Select Bits
	ADCSRA_ADPS_Msk  = 0x7  // ADC  Prescaler Select Bits
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

// Bitfields for EXINT: External Interrupts
const (
	// MCUCR: MCU Control Register
	MCUCR_ISC00    = 0x1  // Interrupt Sense Control 0
	MCUCR_ISC01    = 0x2  // Interrupt Sense Control 0
	MCUCR_ISC0_Msk = 0x3  // Interrupt Sense Control 0
	MCUCR_PUD      = 0x40 // Pull-up Disable
	MCUCR_PUD_Msk  = 0x40 // Pull-up Disable
	MCUCR_SE       = 0x20 // Sleep Enable
	MCUCR_SE_Msk   = 0x20 // Sleep Enable
	MCUCR_SM0      = 0x8  // Sleep Mode Select Bits
	MCUCR_SM1      = 0x10 // Sleep Mode Select Bits
	MCUCR_SM_Msk   = 0x18 // Sleep Mode Select Bits

	// GIMSK: General Interrupt Mask Register
	GIMSK_INT0     = 0x40 // External Interrupt Request 0 Enable
	GIMSK_INT0_Msk = 0x40 // External Interrupt Request 0 Enable
	GIMSK_PCIE0    = 0x10 // Pin Change Interrupt Enables
	GIMSK_PCIE1    = 0x20 // Pin Change Interrupt Enables
	GIMSK_PCIE_Msk = 0x30 // Pin Change Interrupt Enables

	// GIFR: General Interrupt Flag register
	GIFR_INTF0     = 0x40 // External Interrupt Flag 0
	GIFR_INTF0_Msk = 0x40 // External Interrupt Flag 0
	GIFR_PCIF0     = 0x10 // Pin Change Interrupt Flags
	GIFR_PCIF1     = 0x20 // Pin Change Interrupt Flags
	GIFR_PCIF_Msk  = 0x30 // Pin Change Interrupt Flags
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

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TIMSK0: Timer/Counter Interrupt Mask Register
	TIMSK0_OCIE0B     = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0B_Msk = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A     = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_OCIE0A_Msk = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0      = 0x1 // Timer/Counter0 Overflow Interrupt Enable
	TIMSK0_TOIE0_Msk  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag Register
	TIFR0_OCF0B     = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0B_Msk = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0A     = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_OCF0A_Msk = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_TOV0      = 0x1 // Timer/Counter0 Overflow Flag
	TIFR0_TOV0_Msk  = 0x1 // Timer/Counter0 Overflow Flag

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0    = 0x40 // Compare Match Output A Mode bits
	TCCR0A_COM0A1    = 0x80 // Compare Match Output A Mode bits
	TCCR0A_COM0A_Msk = 0xc0 // Compare Match Output A Mode bits
	TCCR0A_COM0B0    = 0x10 // Compare Match Output B Mode bits
	TCCR0A_COM0B1    = 0x20 // Compare Match Output B Mode bits
	TCCR0A_COM0B_Msk = 0x30 // Compare Match Output B Mode bits
	TCCR0A_WGM00     = 0x1  // Waveform Generation Mode bits
	TCCR0A_WGM01     = 0x2  // Waveform Generation Mode bits
	TCCR0A_WGM0_Msk  = 0x3  // Waveform Generation Mode bits

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A     = 0x80 // Force Output Compare A
	TCCR0B_FOC0A_Msk = 0x80 // Force Output Compare A
	TCCR0B_FOC0B     = 0x40 // Force Output Compare B
	TCCR0B_FOC0B_Msk = 0x40 // Force Output Compare B
	TCCR0B_WGM02     = 0x8  // Waveform Generation Mode bit 2
	TCCR0B_WGM02_Msk = 0x8  // Waveform Generation Mode bit 2
	TCCR0B_CS00      = 0x1  // Clock Select bits
	TCCR0B_CS01      = 0x2  // Clock Select bits
	TCCR0B_CS02      = 0x4  // Clock Select bits
	TCCR0B_CS0_Msk   = 0x7  // Clock Select bits

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM       = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_TSM_Msk   = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_PSR10     = 0x1  // Prescaler Reset Timer/CounterN
	GTCCR_PSR10_Msk = 0x1  // Prescaler Reset Timer/CounterN
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1      = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_ICIE1_Msk  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B     = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1B_Msk = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A     = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_OCIE1A_Msk = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1      = 0x1  // Timer/Counter1 Overflow Interrupt Enable
	TIMSK1_TOIE1_Msk  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1      = 0x20 // Timer/Counter1 Input Capture Flag
	TIFR1_ICF1_Msk  = 0x20 // Timer/Counter1 Input Capture Flag
	TIFR1_OCF1B     = 0x4  // Timer/Counter1 Output Compare B Match Flag
	TIFR1_OCF1B_Msk = 0x4  // Timer/Counter1 Output Compare B Match Flag
	TIFR1_OCF1A     = 0x2  // Timer/Counter1 Output Compare A Match Flag
	TIFR1_OCF1A_Msk = 0x2  // Timer/Counter1 Output Compare A Match Flag
	TIFR1_TOV1      = 0x1  // Timer/Counter1 Overflow Flag
	TIFR1_TOV1_Msk  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0    = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1    = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1A_Msk = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0    = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1    = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_COM1B_Msk = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM10     = 0x1  // Pulse Width Modulator Select Bits
	TCCR1A_WGM11     = 0x2  // Pulse Width Modulator Select Bits
	TCCR1A_WGM1_Msk  = 0x3  // Pulse Width Modulator Select Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1     = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICNC1_Msk = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1     = 0x40 // Input Capture 1 Edge Select
	TCCR1B_ICES1_Msk = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10     = 0x8  // Waveform Generation Mode Bits
	TCCR1B_WGM11     = 0x10 // Waveform Generation Mode Bits
	TCCR1B_WGM1_Msk  = 0x18 // Waveform Generation Mode Bits
	TCCR1B_CS10      = 0x1  // Clock Select1 bits
	TCCR1B_CS11      = 0x2  // Clock Select1 bits
	TCCR1B_CS12      = 0x4  // Clock Select1 bits
	TCCR1B_CS1_Msk   = 0x7  // Clock Select1 bits

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A     = 0x80 // Force Output Compare for Channel A
	TCCR1C_FOC1A_Msk = 0x80 // Force Output Compare for Channel A
	TCCR1C_FOC1B     = 0x40 // Force Output Compare for Channel B
	TCCR1C_FOC1B_Msk = 0x40 // Force Output Compare for Channel B
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTIM1     = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM1_Msk = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM0     = 0x4 // Power Reduction Timer/Counter0
	PRR_PRTIM0_Msk = 0x4 // Power Reduction Timer/Counter0
	PRR_PRUSI      = 0x2 // Power Reduction USI
	PRR_PRUSI_Msk  = 0x2 // Power Reduction USI
	PRR_PRADC      = 0x1 // Power Reduction ADC
	PRR_PRADC_Msk  = 0x1 // Power Reduction ADC

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

	// MCUSR: MCU Status Register
	MCUSR_WDRF      = 0x8 // Watchdog Reset Flag
	MCUSR_WDRF_Msk  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF      = 0x4 // Brown-out Reset Flag
	MCUSR_BORF_Msk  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF     = 0x2 // External Reset Flag
	MCUSR_EXTRF_Msk = 0x2 // External Reset Flag
	MCUSR_PORF      = 0x1 // Power-on reset flag
	MCUSR_PORF_Msk  = 0x1 // Power-on reset flag
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
