package jenkinscireceiver

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
)

func TestConfigTag(t *testing.T) {
	t.Parallel()

	require.NoError(t, componenttest.CheckConfigStruct(&Config{}))
}
