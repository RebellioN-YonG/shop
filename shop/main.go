package main

import (
	_ "shop/internal/packed"

	_ "shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"shop/internal/cmd"
	// _ "github.com/gogf/gf/v2/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
