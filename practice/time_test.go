package practice

import (
	"log"
	"testing"
	"time"
)

func ifFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func TestParseDuration(t *testing.T) {
	d, err := time.ParseDuration("1s")
	ifFatal(err)
	t.Log(d.Nanoseconds())
	t.Log(d.Microseconds())
	t.Log(d.Milliseconds())

	//Output:
	//1000000000
	//1000000
	//1000
}

func TestTimeParse(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000"
	now := time.Now()
	t.Log(now.Format(layout))

	layout = "2006-01-02x15:04:05Zggg.0000"
	t.Log(now.Format(layout))
	//time.Parse()
}

