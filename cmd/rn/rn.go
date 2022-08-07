package rn

import (
	"context"
	"github.com/arielhenryson/rn-agent/pkg/api"
	"github.com/arielhenryson/rn-agent/pkg/rnprovider"
	"github.com/sirupsen/logrus"
	cli "github.com/virtual-kubelet/node-cli"
	logruscli "github.com/virtual-kubelet/node-cli/logrus"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	logruslogger "github.com/virtual-kubelet/virtual-kubelet/log/logrus"
)

var (
	buildVersion = "N/A"
	buildTime    = "N/A"
)

func Start() {
	ctx := cli.ContextWithCancelOnSignal(context.Background())

	logger := logrus.StandardLogger()
	log.L = logruslogger.FromLogrus(logrus.NewEntry(logger))
	logConfig := &logruscli.Config{LogLevel: "info"}

	node, err := cli.New(ctx,
		cli.WithCLIVersion(buildVersion, buildTime),
		cli.WithProvider("rn", func(cfg provider.InitConfig) (provider.Provider, error) {
			return rnprovider.NewMockProvider(
				"rn",
				cfg.OperatingSystem,
				cfg.InternalIP,
				cfg.DaemonPort,
			)
		}),
		cli.WithPersistentFlags(logConfig.FlagSet()),
		cli.WithPersistentPreRunCallback(func() error {
			return logruscli.Configure(logConfig, logger)
		}),
	)

	if err != nil {
		log.G(ctx).Fatal(err)
	}

	if err := node.Run(ctx, "--provider", "rn", "--nodename", "rn"); err != nil {
		log.G(ctx).Fatal(err)
	}

	api.StartApiServer()
}
