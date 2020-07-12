package hostname

import (
	"os"
)

// WithUser returns a string with the user and hostname
// e.g user@machine
func WithUser() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	}

	user := os.Getenv("USER")

	return user + "@" + host
}

// String returns the hostname as reported by the kernel
func String() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	}
	return host
}
