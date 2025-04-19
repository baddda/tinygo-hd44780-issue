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

	time.Sleep(100 * time.Millisecond)
	m.D13.High()
	m.D7.Low()
	m.D6.Low()
	m.D5.Low()
	m.D4.Low()
	m.D3.Low()
	m.D2.Low()
	m.D1.Low()
	m.D0.High()
	m.D13.Low()

	time.Sleep(100 * time.Millisecond)
	m.D13.High()
	m.D7.High()
	m.D6.High()
	m.D5.Low()
	m.D4.Low()
	m.D3.Low()
	m.D2.Low()
	m.D1.Low()
	m.D0.High()
	m.D13.Low()
}
