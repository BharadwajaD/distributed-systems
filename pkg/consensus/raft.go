package consensus

import "BharadwajaD/DistSys/pkg/base"

type RaftKVMap struct {
	KVMap map[string]string
	node *base.Connector
}

func (rmap *RaftKVMap) Put(key, value string) {
}

/*
Leader, Candidate, Follower
If a follower doesn't hear a heartbeat from leader, they become candidate and ask
for votes (election - term)
Leader with highest term is valid.. Lower term => node broke


Log Replication
- two phase commit ( followers ack msg recieved and then leader commits into its log)
-
*/
