package practice

import (
	"fmt"
	"testing"
)

func TestFloatFormat(t *testing.T) {
	num := float64(10)
	//strconv.ParseFloat()
	fmt.Println(fmt.Sprintf("%v", num))
}