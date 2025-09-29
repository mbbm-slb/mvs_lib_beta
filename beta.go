package mvslibbeta

import (
	"fmt"
	"github.com/mbbm-slb/mvs_lib_gamma"
)

func Beta() string {
	return fmt.Sprintf("Beta forwarding: %s", mvslibgamma.Gamma())
}