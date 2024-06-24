package main

import (
	_ "Literary_Snowlands/internal/logic"
	_ "Literary_Snowlands/internal/logic/author"
	_ "Literary_Snowlands/internal/logic/book"
	_ "Literary_Snowlands/internal/logic/home"
	_ "Literary_Snowlands/internal/logic/middleware"
	_ "Literary_Snowlands/internal/logic/news"
	_ "Literary_Snowlands/internal/logic/search"
	_ "Literary_Snowlands/internal/logic/user"
	_ "Literary_Snowlands/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gogf/gf/v2/os/gctx"

	"Literary_Snowlands/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
