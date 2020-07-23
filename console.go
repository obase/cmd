package cmd

import (
	"fmt"
	"os"
	"time"
)

const ISO_DATE_TIME = "2006-01-02 15:04:05"

func PrintInfo(format string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [I] %s\n", time.Now().Format(ISO_DATE_TIME), fmt.Sprintf(format, val...))
}

func PrintError(format string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [E] %s\n", time.Now().Format(ISO_DATE_TIME), fmt.Sprintf(format, val...))
}
