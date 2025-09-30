package mvslibbeta

import (
	"fmt"
	"github.com/mbbm-slb/mvs_lib_gamma/v2"
)

func Beta() string {
	return fmt.Sprintf("Beta forwarding: %s", mvslibgamma.Gamma("beta uses v2 of gamma", "yay!"))
}