package consensus

import "BharadwajaD/DistSys/pkg/base"

type LeaderWrapper struct {
	node *base.Connector
}

func (lw *LeaderWrapper) IsLeader() bool {
	return false
}

func (lw *LeaderWrapper) ElectLeader() {
}
