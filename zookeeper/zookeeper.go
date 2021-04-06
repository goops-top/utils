/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : zookeeper.go
@Time   : 2021/04/06 10:37:56
@Update : 2021/04/06 10:37:56
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : None
*/
package zookeeper

import (
	"fmt"
	"time"

	zk "github.com/samuel/go-zookeeper/zk"
)

type zookeeperMeta struct {
	Servers        []string      `json:"servers"`
	Namespace      string        `json:"namespace"`
	SessionTimeout time.Duration `json:"timeout"`
}

// 使用function options 来构建整个参数
type zkOptions func(z *zookeeperMeta)

// 使用高阶函数来构造function option
func Timeout(timeout time.Duration) zkOptions {
	return func(z *zookeeperMeta) {
		z.SessionTimeout = timeout
	}
}

func Namespace(ns string) zkOptions {
	return func(z *zookeeperMeta) {
		z.Namespace = ns
	}
}

func NewZookeeperClusterMeta(servers []string, options ...zkOptions) *zookeeperMeta {
	zkIns := zookeeperMeta{
		Servers: servers,
	}

	for _, option := range options {
		option(&zkIns)
	}

	return &zkIns
}

func (z *zookeeperMeta) NewZookeeperConn() (*zk.Conn, error) {
	conn, _, err := zk.Connect(z.Servers, z.SessionTimeout)

	if err != nil {
		return nil, err
	}
	return conn, err
}

func ZkStateStringFormat(s *zk.Stat) string {
	/*
		   type Stat struct {
		   	Czxid          int64 // The zxid of the change that caused this znode to be created.
		   	Mzxid          int64 // The zxid of the change that last modified this znode.
		   	Ctime          int64 // The time in milliseconds from epoch when this znode was created.
		   	Mtime          int64 // The time in milliseconds from epoch when this znode was last modified.
		   	Version        int32 // The number of changes to the data of this znode.
		   	Cversion       int32 // The number of changes to the children of this znode.
		   	Aversion       int32 // The number of changes to the ACL of this znode.
			// 可以以此来判别是否为临时节点，如果不是临时节点，该值将为0
		   	EphemeralOwner int64 // The session id of the owner of this znode if the znode is an ephemeral node. If it is not an ephemeral node, it will be zero.
		   	DataLength     int32 // The length of the data field of this znode.
		   	NumChildren    int32 // The number of children of this znode.
		   	Pzxid          int64 // last modified children
		   }
	*/
	return fmt.Sprintf(
		`createTime: %d
		modifiedTime: %d
		version: %d
		Cversion: %d
		isEphemeral: %d
		dataLength: %d
		NumChildren: %d`,
		s.Ctime, s.Mtime, s.Version, s.Cversion, s.EphemeralOwner, s.DataLength, s.NumChildren)
}
