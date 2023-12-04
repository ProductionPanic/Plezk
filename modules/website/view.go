package website

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"plezk/lib/common"
	"plezk/lib/domain"
)

func View(dom domain.Domain) int {
	common.Cls()
	fmt.Println(dom.Info().GetInfoString())
	_, _, err := keyboard.GetSingleKey()
	if err != nil {
		return 0
	}
	return 0
}
