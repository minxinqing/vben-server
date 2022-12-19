package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "test-gf/internal/packed"

	_ "test-gf/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"test-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
