package main

import (
	"context"
	"github.com/liuzhaomax/maxblog-main/internal/app"
	"github.com/liuzhaomax/maxblog-main/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
