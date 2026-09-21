package app

import "fmt"

func VersionedGreeting(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("teltonika-playground says hello, %s", name)
}
