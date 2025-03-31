#gpio in golang 

setup(): Initializes memory mapping.

setupGPIO(gpio, direction, pud): Configures a GPIO pin.

gpioFunction(gpio): Gets the function of a GPIO pin.

outputGPIO(gpio, value): Sets GPIO output.

inputGPIO(gpio): Reads GPIO input.

setRisingEvent(gpio, enable), setFallingEvent(gpio, enable), setHighEvent(gpio, enable), setLowEvent(gpio, enable): Configures edge detection.

eventDetected(gpio): Checks if an event occurred.

cleanup(): Unmaps memory.
