package main

import (
	"fmt"
	"testing"
)

func TestViper(t *testing.T) {
	conf := ReadConfigViper()
	fmt.Println(conf)
}
