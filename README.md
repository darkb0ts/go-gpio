# GPIO in Go

A lightweight Go library for controlling General Purpose Input/Output (GPIO) pins on embedded systems, such as the Raspberry Pi. This library provides a simple and efficient way to configure, read, write, and monitor GPIO pins using memory-mapped I/O.

## Features
- Initialize and configure GPIO pins with ease.
- Set pin direction (input/output) and pull-up/pull-down resistors.
- Read and write digital values to GPIO pins.
- Configure edge detection for rising, falling, high, and low events.
- Check for detected events on GPIO pins.
- Clean up resources properly when done.

## Installation
To use this library in your Go project, run:

```bash
go get github.com/darkb0ts/go-gpio
```

## Usage
Here’s a quick example of how to use the library:

```go
package main

import (
	"fmt"
	"github.com/darkb0ts/go-gpio/gpio"
)

func main() {
	// Initialize GPIO memory mapping
	err := gpio.Setup()
	if err != nil {
		fmt.Println("Failed to initialize GPIO:", err)
		return
	}
	defer gpio.Cleanup() // Ensure cleanup on exit

	// Configure GPIO pin 18 as an output
	gpioPin := 18
	err = gpio.SetupGPIO(gpioPin, gpio.Output, gpio.PullOff)
	if err != nil {
		fmt.Println("Failed to configure GPIO:", err)
		return
	}

	// Set the pin high
	gpio.OutputGPIO(gpioPin, true)
	fmt.Println("GPIO", gpioPin, "set to HIGH")

	// Read the pin's function
	func := gpio.GPIOFunction(gpioPin)
	fmt.Println("GPIO", gpioPin, "function:", func)
}
```

## API Reference

### `Setup() error`
Initializes memory mapping for GPIO access. Must be called before using any GPIO functions.

### `SetupGPIO(gpio int, direction Direction, pud PullUpDown) error`
Configures a GPIO pin with the specified direction (`Input` or `Output`) and pull-up/pull-down resistor setting (`PullUp`, `PullDown`, or `PullOff`).

### `GPIOFunction(gpio int) string`
Returns the current function of the specified GPIO pin (e.g., "input", "output").

### `OutputGPIO(gpio int, value bool)`
Sets the output state of a GPIO pin (`true` for HIGH, `false` for LOW).

### `InputGPIO(gpio int) (bool, error)`
Reads the current input state of a GPIO pin (`true` for HIGH, `false` for LOW).

### `SetRisingEvent(gpio int, enable bool)`
Enables or disables rising edge detection for the specified GPIO pin.

### `SetFallingEvent(gpio int, enable bool)`
Enables or disables falling edge detection for the specified GPIO pin.

### `SetHighEvent(gpio int, enable bool)`
Enables or disables high-level event detection for the specified GPIO pin.

### `SetLowEvent(gpio int, enable bool)`
Enables or disables low-level event detection for the specified GPIO pin.

### `EventDetected(gpio int) bool`
Checks if an event (configured via `Set*Event`) has occurred on the specified GPIO pin.

### `Cleanup()`
Unmaps the memory used for GPIO access. Call this when you're done using the library to free resources.

## Prerequisites
- Go 1.x or higher
- Access to GPIO hardware (e.g., Raspberry Pi with appropriate permissions)
- Run as root or with sufficient privileges (e.g., using `sudo`) for memory mapping

## Notes
- This library uses direct memory mapping, so it’s intended for use on systems with accessible GPIO registers (e.g., Raspberry Pi).
- Ensure proper cleanup with `Cleanup()` to avoid memory leaks or locked resources.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes. For major updates, open an issue first to discuss your ideas.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
