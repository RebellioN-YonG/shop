package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"shop/internal/cmd"
	_ "shop/internal/logic"
	_ "shop/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
