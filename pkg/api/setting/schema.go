package setting

import (
	"github.com/rancher/apiserver/pkg/types"
	"github.com/rancher/steve/pkg/schema"
	"github.com/rancher/steve/pkg/server"
	"github.com/rancher/steve/pkg/stores/proxy"

	"github.com/harvester/harvester/pkg/config"
)

func RegisterSchema(scaled *config.Scaled, server *server.Server, options config.Options) error {
	t := schema.Template{
		ID:        "harvesterhci.io.setting",
		Store:     proxy.NewProxyStore(server.ClientFactory, nil, server.AccessSetLookup),
		Customize: func(s *types.APISchema) {},
	}
	server.SchemaFactory.AddTemplate(t)
	return nil
}
