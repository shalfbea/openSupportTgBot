package main

import (
	"context"

	"github.com/shalfbea/openSupportTgBot/pkg/config"
	"github.com/shalfbea/openSupportTgBot/pkg/logger"
	"github.com/shalfbea/openSupportTgBot/pkg/telegram"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func testRun(cfg *config.Config, logger logger.Logger) {
	logger.Infof("Here would be bot starting... Started with config: %+#v\n", *cfg)
}

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.InitLogger,
			telegram.NewBot,
		),
		fx.Invoke(
			testRun,
			telegram.RunBot,
		),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
