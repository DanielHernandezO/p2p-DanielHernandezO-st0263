package shell

import (
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/config"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/delivery"
)

type shellCommands struct{}

func NewShellCommands() delivery.Strategy {
	return &shellCommands{}
}

func (s *shellCommands) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()
	mapActions(dependencies)
}
