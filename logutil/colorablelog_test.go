package logutil_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jhunters/goassist/logutil"
)

func TestLog(t *testing.T) {
	log := logutil.CreateLogger("info", logutil.GREEN)
	fmt.Fprintf(log, "Hello %s", "World")

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log2 := logutil.CreateLoggerToFile("logfile", f, logutil.GREEN)
	log2.Write([]byte("hello world\n"))

	time.Sleep(1 * time.Second)
}
