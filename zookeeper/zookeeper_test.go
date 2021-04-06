package zookeeper

import (
	"fmt"
	"testing"
	"time"
)

func TestZKConn(t *testing.T) {
	servers := []string{"zookeeper.c.t.soulapp-inc.cn:2181"}
	clusterMeta := NewZookeeperClusterMeta(servers, Namespace("/dubbo"), Timeout(1*time.Second))
	fmt.Println(clusterMeta)

	conn, err := clusterMeta.NewZookeeperConn()
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(clusterMeta.Namespace)
	zkKeys, state, err := conn.Children(clusterMeta.Namespace)
	if err != nil {
		panic(err)
	}

	fmt.Println(zkKeys, ZkStateStringFormat(state))
	for _, key := range zkKeys {
		subKey := fmt.Sprintf("%s/%s", clusterMeta.Namespace, key)
		fmt.Println(subKey)
		data, s, _ := conn.Get(subKey)
		fmt.Println(string(data), ZkStateStringFormat(s))
	}
	//fmt.Println(conn, err)
}
