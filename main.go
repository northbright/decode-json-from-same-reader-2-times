package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Laptop defines the laptop interface.
type Laptop interface {
	Brand() string
	OS() string
}

// Base contains the common information of the laptop.
type Base struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}

// MacBook represents the Apple MacBook.
type MacBook struct {
	Base
	WithAppleSilicon bool `json:"with_apple_silicon"`
}

func (m *MacBook) Brand() string {
	return "Apple"
}

func (m *MacBook) OS() string {
	return "macOS"
}

// WindowsLaptop defines the Windows Laptop.
type WindowsLaptop struct {
	Base
	OfficePreInstalled bool `json:"office_pre_installed"`
}

func (w *WindowsLaptop) Brand() string {
	return w.Manufacturer
}

func (w *WindowsLaptop) OS() string {
	return "Windows"
}

// LoadFromJSON reads JSON data and returns a laptop.
func LoadFromJSON(data string) Laptop {
	var (
		buf  bytes.Buffer
		base Base
	)

	// Create a string reader from data.
	r := strings.NewReader(data)

	// Create a io.TeeReader to duplicate the stream.
	tr := io.TeeReader(r, &buf)

	// 1st Decode: decode data to a Base variable and get the manufacturer.
	dec := json.NewDecoder(tr)
	dec.Decode(&base)

	// 2nd Decode: decode data to MacBook or WindowsLaptop according to the manufacturer.
	switch base.Manufacturer {
	case "Apple":
		newDec := json.NewDecoder(&buf)
		mac := MacBook{}
		newDec.Decode(&mac)
		return &mac
	default:
		newDec := json.NewDecoder(&buf)
		laptop := WindowsLaptop{}
		newDec.Decode(&laptop)
		return &laptop
	}
}

func main() {
	// arr contains laptop JSON data.
	arr := []string{
		mbp,
		lenovoY7000,
	}

	for i, data := range arr {
		// Create a laptop by loading JSON data.
		laptop := LoadFromJSON(data)

		fmt.Printf("========== Laptop: %d ==========\n", i)

		switch vv := laptop.(type) {
		case *MacBook:
			fmt.Printf("%v\n", vv.Model)
			fmt.Printf("with Apple Silicon: %v\n", vv.WithAppleSilicon)
		case *WindowsLaptop:
			fmt.Printf("%v %v\n", vv.Manufacturer, vv.Model)
			fmt.Printf("Office Pre-Installed: %v\n", vv.OfficePreInstalled)
		default:
			fmt.Printf("Unknown Laptop\n")
		}

		fmt.Printf("===============================\n\n")
	}
}

var (
	mbp = `
{
  "manufacturer": "Apple",
  "model": "MacBook Pro 160in.(M3 Max)",
  "with_apple_silicon": true
}
`

	lenovoY7000 = `
{
  "manufacturer": "Lenovo",
  "model": "Y7000",
  "office_pre_installed": true
}
`
)
