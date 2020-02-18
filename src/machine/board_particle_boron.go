// +build particle_boron

package machine

const HasLowFrequencyCrystal = true

// More info: https://docs.particle.io/datasheets/cellular/boron-datasheet/
// Board diagram: https://docs.particle.io/assets/images/boron/boron-block-diagram.png

// LEDs
const (
	LED       Pin = 44
	LED_GREEN Pin = 14
	LED_RED   Pin = 13
	LED_BLUE  Pin = 15
)

// GPIOs
const (
	A0  Pin = 3
	A1  Pin = 4
	A2  Pin = 28
	A3  Pin = 29
	A4  Pin = 30
	A5  Pin = 31
	D0  Pin = 26 // Also SDA
	D1  Pin = 27 // Also SCL
	D2  Pin = 33
	D3  Pin = 34
	D4  Pin = 40
	D5  Pin = 42
	D6  Pin = 43
	D7  Pin = 44 // Also LED
	D8  Pin = 35
	D9  Pin = 6  // Also TX
	D10 Pin = 8  // Also RX
	D11 Pin = 46 // Also MISO
	D12 Pin = 45 // Also MOSI
	D13 Pin = 47 // Also SCK
)

// UART
const (
	Serial          = USB
	UART0           = NRF_UART0
	UART_TX_PIN Pin = 6
	UART_RX_PIN Pin = 8
)

// I2C pins
const (
	SDA_PIN Pin = 26
	SCL_PIN Pin = 27

	// Internal I2C with MAX17043 (Fuel gauge) and BQ24195 (Power management) chips on it
	SDA1_PIN Pin = 24
	SCL1_PIN Pin = 41
	INT1_PIN Pin = 5
)

// SPI pins
const (
	SPI0_SCK_PIN  Pin = 47
	SPI0_MOSI_PIN Pin = 45
	SPI0_MISO_PIN Pin = 46
)

// Internal 4MB SPI Flash
const (
	SPI1_SCK_PIN  Pin = 19
	SPI1_MOSI_PIN Pin = 20
	SPI1_MISO_PIN Pin = 21
	SPI1_CS_PIN   Pin = 17
	SPI1_WP_PIN   Pin = 22
	SPI1_HOLD_PIN Pin = 23
)

// u-blox coprocessor
const (
	UBLOX_TXD_PIN      Pin = 37
	UBLOX_RXD_PIN      Pin = 36
	UBLOX_CTS_PIN      Pin = 38
	UBLOX_RTS_PIN      Pin = 39
	UBLOX_RESET_PIN    Pin = 12
	UBLOX_POWER_ON_PIN Pin = 16
	UBLOX_BUFF_EN_PIN  Pin = 25
	UBLOX_VINT_PIN     Pin = 2
)

// Other periferals
const (
	MODE_BUTTON_PIN Pin = 11
	ANTENNA_SEL_PIN Pin = 7 // Low: chip antenna, High: External uFL
	NFC1_PIN        Pin = 9
	NFC2_PIN        Pin = 10
)