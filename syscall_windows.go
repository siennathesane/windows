package windows

import "syscall"

// Getenv get environment variables via syscall.
func Getenv(s string) (string, bool) {
	return syscall.Getenv(s)
}

func itoa(x int) string {
	if x < 0 {
		return "-" + itoa(-x)
	}

	var buffer[32]byte

	i := len(buffer) - 1

	for x >= 10 {
		buffer[i]  = byte(x % 10 + '0')
		i--
		x /= 10
	}

	buffer[i] = byte(x + '0')

	return string(buffer[i:])
}