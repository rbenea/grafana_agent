package build

import (
	"github.com/grafana/agent/component/discovery"
	"github.com/grafana/agent/component/prometheus/exporter/oracledb"
	"github.com/grafana/agent/pkg/integrations/oracledb_exporter"
	"github.com/grafana/river/rivertypes"
)

func (b *IntegrationsConfigBuilder) appendOracledbExporter(config *oracledb_exporter.Config, instanceKey *string) discovery.Exports {
	args := toOracledbExporter(config)
	return b.appendExporterBlock(args, config.Name(), instanceKey, "oracledb")
}

func toOracledbExporter(config *oracledb_exporter.Config) *oracledb.Arguments {
	return &oracledb.Arguments{
		ConnectionString: rivertypes.Secret(config.ConnectionString),
		MaxIdleConns:     config.MaxIdleConns,
		MaxOpenConns:     config.MaxOpenConns,
		QueryTimeout:     config.QueryTimeout,
	}
}
