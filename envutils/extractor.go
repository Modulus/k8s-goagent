package envutils

import (
	"fmt"
	"os"
)

func ExtractPodFilter() string {
	filter := os.Getenv("POD_FILTER")

	fmt.Printf("POD_FILTER = %s\n", filter)

	return filter
}
