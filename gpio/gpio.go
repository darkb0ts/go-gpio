package gpio

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	BCM2711_PERI_BASE   = 0xfe000000
	GPIO_BASE_OFFSET    = 0x200000
	BLOCK_SIZE          = 4096
	INPUT               = 0
	OUTPUT              = 1
	HIGH                = 1
	LOW                 = 0
	FSEL_OFFSET         = 0x00
	SET_OFFSET          = 0x1C / 4
	CLR_OFFSET          = 0x28 / 4
	PINLEVEL_OFFSET     = 0x34 / 4
	EVENT_DETECT_OFFSET = 0x40 / 4
	RISING_ED_OFFSET    = 0x4C / 4
	FALLING_ED_OFFSET   = 0x58 / 4
	HIGH_DETECT_OFFSET  = 0x64 / 4
	LOW_DETECT_OFFSET   = 0x70 / 4
)

var gpioMem []uint32
var logger bool = false

func Setup() error {
	mem, err := os.OpenFile("/dev/gpiomem", os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
			return fmt.Errorf("failed to open /dev/gpiomem: %v", err)
	}
	defer mem.Close()

	gpioBase := uintptr(BCM2711_PERI_BASE + GPIO_BASE_OFFSET)
	memMap, err := syscall.Mmap(int(mem.Fd()), int64(gpioBase), BLOCK_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
			return log.ERROR("failed to mmap: %v", err)
			log.Printf("ERROR: failed to mmap: %v", err)
					return err
			}

	// Ensure that the mmap is unmapped when the function returns.
	defer syscall.Munmap(memMap)

	gpioMem = (*(*[BLOCK_SIZE / 4]uint32)(unsafe.Pointer(&memMap[0])))[:]

	return nil
}

func SetupGPIO(gpio, direction int) {
	offset := FSEL_OFFSET + (gpio / 10)
	shift := (gpio % 10) * 3
	gpioMem[offset] &^= (7 << shift)
	if direction == OUTPUT {
		gpioMem[offset] |= (1 << shift)
	}
}

func SetupGPIOArray(gpios []int, direction int) {
	for _, gpio := range gpios {
		SetupGPIO(gpio, direction)
	}
}

func OutputGPIOArray(gpios []int, value int) {
	for _, gpio := range gpios {
		OutputGPIO(gpio, value)
	}
}

func OutputGPIO(gpio, value int) {
	offset := CLR_OFFSET
	action := "LOW"

	if value == HIGH {
			offset = SET_OFFSET
			action = "HIGH"
	}

	if logger {
			log.Printf("INFO: Setting GPIO %d to %s", gpio, action)
	}

	gpioMem[offset+gpio/32] = 1 << (gpio % 32)
}

func InputGPIO(gpio int) int {
	if gpioMem[PINLEVEL_OFFSET+gpio/32]&(1<<(gpio%32)) != 0 {
		return HIGH
	}
	return LOW
}

func InputGPIOArray(gpios []int) []int {
	inputs := make([]int, len(gpios))
	for i, gpio := range gpios {
		inputs[i] = InputGPIO(gpio)
	}
	return inputs
}

func SetRisingEvent(gpio, enable int) {
	offset := RISING_ED_OFFSET + (gpio / 32)
	shift := (gpio % 32)
	if enable != 0 {
		gpioMem[offset] |= (1 << shift)
	} else {
		gpioMem[offset] &^= (1 << shift)
	}
}

func SetFallingEvent(gpio, enable int) {
	offset := FALLING_ED_OFFSET + (gpio / 32)
	shift := (gpio % 32)
	if enable != 0 {
		gpioMem[offset] |= (1 << shift)
	} else {
		gpioMem[offset] &^= (1 << shift)
	}
}

func SetHighEvent(gpio, enable int) {
	offset := HIGH_DETECT_OFFSET + (gpio / 32)
	shift := (gpio % 32)
	if enable != 0 {
		gpioMem[offset] |= (1 << shift)
	} else {
		gpioMem[offset] &^= (1 << shift)
	}
}

func SetLowEvent(gpio, enable int) {
	offset := LOW_DETECT_OFFSET + (gpio / 32)
	shift := (gpio % 32)
	if enable != 0 {
		gpioMem[offset] |= (1 << shift)
	} else {
		gpioMem[offset] &^= (1 << shift)
	}
}

func EventDetected(gpio int) int {
	offset := EVENT_DETECT_OFFSET + (gpio / 32)
	shift := (gpio % 32)
	if gpioMem[offset]&(1<<shift) != 0 {
		gpioMem[offset] = 1 << shift // Clear event
		return 1
	}
	return 0
}

func Cleanup() {
	unix.Munmap((*[BLOCK_SIZE]byte)(unsafe.Pointer(&gpioMem[0]))[:BLOCK_SIZE])
}

