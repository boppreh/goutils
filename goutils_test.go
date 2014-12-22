package goutils

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	fmt.Println(SearchUrl("google.com/robots.txt", `Allow: (.+)`))
}
