package tensms

import (
	"log"

	"testing"
)

func TestOption(t *testing.T) {
	SetInfo("xxx", "xxx")
	res, err := GetTpl([]int{182993, 112873})
	log.Printf("%#v\n%#v\n", res, err)
}
