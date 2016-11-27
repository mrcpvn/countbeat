package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/mrcpvn/countbeat/beater"
)

func main() {
	err := beat.Run("countbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
