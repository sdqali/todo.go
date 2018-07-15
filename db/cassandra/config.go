package cassandra

import (
	"os"

	"github.com/gocql/gocql"
)

func GetCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Keyspace = "todo"
	cluster.ProtoVersion = 3
	cluster.DisableInitialHostLookup = true
	return cluster
}
