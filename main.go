package main

import (
	m "machine"
	"time"
)

func main() {
	m.D0.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D1.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D2.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D3.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D4.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D5.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D6.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D7.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D12.Configure(m.PinConfig{Mode: m.PinOutput})
	m.D13.Configure(m.PinConfig{Mode: m.PinOutput})

	/* --- */

	time.Sleep(100 * time.Millisecond)
	m.D13.High() // Start command.

	// Sets to 8-bit operation and selects 2-line display and 5 × 8 dot character font.
	m.D7.Low()
	m.D6.Low()
	m.D5.High()
	m.D4.High()
	m.D3.High()
	m.D2.Low()
	m.D1.Low()
	m.D0.Low()

	m.D13.Low() // Finish command.

	/* --- */

	time.Sleep(100 * time.Millisecond)
	m.D13.High() // Start command.

	// Turns on display and cursor.
	m.D7.Low()
	m.D6.Low()
	m.D5.Low()
	m.D4.Low()
	m.D3.High()
	m.D2.High()
	m.D1.High()
	m.D0.Low()

	m.D13.Low() // Finish command.

	/* --- */

	time.Sleep(100 * time.Millisecond)
	m.D13.High() // Start command.

	// Write clear display command. 00000001
	m.D7.Low()
	m.D6.Low()
	m.D5.Low()
	m.D4.Low()
	m.D3.Low()
	m.D2.Low()
	m.D1.Low()
	m.D0.High()

	m.D13.Low() // Finish command.

	/* --- */

	time.Sleep(100 * time.Millisecond)
	m.D13.High() // Start command.

	// Write set cursor position command. 11000001
	m.D7.High() // Indicate set address command.
	m.D6.High() // Indicate column 2.
	m.D5.Low()
	m.D4.Low()
	m.D3.Low()
	m.D2.Low()
	m.D1.Low()
	m.D0.High() // Indicate row 1.

	m.D13.Low() // Finish command.
}
