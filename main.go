package main

import (
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	_ "github.com/cellargalaxy/msg_gateway/sdk"
)

func main() {
	fmt.Println(util.GenId())
}
