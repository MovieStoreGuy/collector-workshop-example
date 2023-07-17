package jenkinscireceiver

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/splunk/collector-workshop-example/jenkinsreceiver/internal/metadata"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		newDefaultConfig,
		receiver.WithMetrics(
			newMetricsReceiver,
			metadata.MetricsStability,
		),
	)
}
func newMetricsReceiver(
	ctx context.Context,
	set receiver.CreateSettings,
	cfg component.Config,
	next consumer.Metrics,
) (receiver.Metrics, error) {
	conf, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("can not convert to *Config")
	}

	s, err := newScraper(conf, set)
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(
		&conf.ScraperControllerSettings,
		set,
		next,
		scraperhelper.AddScraper(s),
	)
}
