package helpers

import (
	"github.com/go-redis/redis/v7"
)

// ConnectToCluster :
func ConnectToCluster(clusterAddrs []string) *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: clusterAddrs,
	})

	return clusterClient
}

// ConnectToSingleNode :
func ConnectToSingleNode(transport string, address string, poolSize int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network:  transport,
		Addr:     address,
		PoolSize: poolSize,
	})

	return client
}
