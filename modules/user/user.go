package user

import "os"

func String() string {
	user := os.Getenv("USER")
	return user
}
