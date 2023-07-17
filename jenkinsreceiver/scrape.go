package jenkinscireceiver

import (
	"context"
	"time"

	jenkins "github.com/yosida95/golang-jenkins"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/splunk/collector-workshop-example/jenkinsreceiver/internal/metadata"
)

type scraper struct {
	mb       *metadata.MetricsBuilder
	settings component.TelemetrySettings

	client *jenkins.Jenkins
}

func newScraper(cfg *Config, set receiver.CreateSettings) (scraperhelper.Scraper, error) {
	s := &scraper{
		settings: set.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, set),
	}
	return scraperhelper.NewScraper(
		metadata.Type,
		s.scrape,
		scraperhelper.WithStart(func(ctx context.Context, h component.Host) error {
			client, err := cfg.ToClient(h, s.settings)
			if err != nil {
				return err
			}
			// The collector provides a means of injecting authentication
			// on our behalf, so this will ignore the libraries approach
			// and use the configured http client with authentication.
			s.client = jenkins.NewJenkins(nil, cfg.Endpoint)
			s.client.SetHTTPClient(client)
			return nil
		}),
	)
}

func (s *scraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	jobs, err := s.client.GetJobs()
	if err != nil {
		return pmetric.Metrics{}, err
	}

	now := pcommon.NewTimestampFromTime(time.Now())
	s.mb.RecordJenkinsJobsCountDataPoint(now, int64(len(jobs)))

	for _, job := range jobs {
		// Fetch the most recent build
		var (
			build  = job.LastCompletedBuild
			status = metadata.AttributeJobStatusUnknown
		)

		switch build.Result {
		case "aborted", "not_built", "unstable":
			// Do nothing
		case "success":
			status = metadata.AttributeJobStatusSuccess
		case "failure":
			status = metadata.AttributeJobStatusFailed
		}

		s.mb.RecordJenkinsJobDurationDataPoint(
			now,
			int64(job.LastCompletedBuild.Duration),
			job.Name,
			status,
		)

		if len(build.ChangeSet.Items) == 0 {
			continue
		}
		change := build.ChangeSet.Items[0]

		s.mb.RecordJenkinsJobCommitDeltaDataPoint(
			now,
			int64(build.Timestamp-change.Timestamp),
			job.Name,
			status,
		)
	}

	return s.mb.Emit(), nil
}
