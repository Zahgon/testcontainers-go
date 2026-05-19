package nebulagraph

import (
	_ "embed"

	"github.com/testcontainers/testcontainers-go"
)

//go:embed activator.sh
var activatorScript string

const user = "root"

const (
	defaultNebulaConsoleImage = "vesoft/nebula-console:v3.8.0"
)

const (
	metadNetworkAlias    = "metad0"
	graphdNetworkAlias   = "graphd0"
	storagedNetworkAlias = "storaged0"
)

const (
	graphdPort   = "9669"
	metadPort    = "9559"
	storagedPort = "9779"

	graphdPortHTTP   = "19669"
	metadPortHTTP    = "19559"
	storagedPortHTTP = "19779"
)

func defaultGraphdContainerCustomizers(nw *testcontainers.DockerNetwork) []testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return nil
}

func defaultMetadContainerCustomizers(nw *testcontainers.DockerNetwork) []testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return nil
}

func defaultStoragedContainerCustomizers(nw *testcontainers.DockerNetwork) []testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return nil
}

func defaultActivatorContainerCustomizers(nw *testcontainers.DockerNetwork) []testcontainers.ContainerCustomizer {
	_ = "STUB: not implemented"
	return nil
}
