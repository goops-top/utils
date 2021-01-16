package kafka

import (
	"testing"

	"fmt"
	"strings"
)

type TopicDetail struct {
	NumPartitions     int32
	ReplicationFactor int16
	ReplicaAssignment map[int32][]int32
	ConfigEntries     map[string]*string
}

type TopicMetadata struct {
	Err        string
	Name       string
	IsInternal bool // Only valid for Version >= 1
	Partitions []*PartitionMetadata
}
type PartitionMetadata struct {
	Err             string
	ID              int32
	Leader          int32
	Replicas        []int32
	Isr             []int32
	OfflineReplicas []int32
}

func TestListTopic(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092", "172.29.203.106:9092", "172.29.203.86:9092"})
	defer admin.Close()
	// 原始的topic的map结构,应该比较好获取topic名称
	topics, _ := admin.ListTopic()
	for topic, tInfo := range topics {
		fmt.Println(topic, tInfo)
	}
	// 获取指定topic列表并解析到slice中
	topicsInfos, err := admin.ListTopicsInfo([]string{"data-metadata", "myapp-yum-log"})
	fmt.Println(err)
	for _, t := range topicsInfos {
		fmt.Println(t.Name, t.PartitionNum, t.Replication, t.ReplicaAssignment, t.ConfigEntries)
		fmt.Printf("Config:")
		for name, conf := range t.ConfigEntries {
			fmt.Printf("%v:%v,\n", name, *conf)
		}
	}
}

func TestDescribeTopics(t *testing.T) {
	// admin := NewClusterAdmin([]string{"172.29.203.62:9092","172.29.203.106:9092","172.29.203.86:9092"})
	admin := NewClusterAdmin([]string{"bgbiao.top:9092"})
	defer admin.Close()
	// 原始的topic元数据信息
	topicInfo, _ := admin.describeTopics([]string{"imfullpushflumelog"})
	for _, v := range topicInfo {
		fmt.Println(v.Name)
		for _, p := range v.Partitions {
			fmt.Println(p.ID, p.Leader, p.Replicas, p.Isr, p.OfflineReplicas)
		}
	}
	// 获取指定topic列表并解析到slice中
	topicMetas, _ := admin.DescribeTopics([]string{})
	for _, v := range topicMetas {
		fmt.Println(v)
	}
}

// describe the cluster
// return : []Broker,controllerId,error
// https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#Broker
func TestDescribeCluster(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	brokers, controllerId, clusterErr := admin.DescribeCluster()
	if clusterErr != nil {
		fmt.Printf("err:%v\n", clusterErr)
	}
	fmt.Println("current controllerd id:", controllerId)
	for _, broker := range brokers {
		fmt.Printf("broker:%v,broker_id:%v\n", broker.Addr(), broker.ID())
	}

}

// create topic with default config.
func TestCreateTopic(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()
	isok, err := admin.CreateTopic("gokafka-test")
	fmt.Println(isok, err)
}

// create topic with specifiy some params.
func TestCreateCustomTopic(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()
	isok, err := admin.CreateCustomTopic("gokafka-test-new", 10, 1, map[string]string{"unclean.leader.election.enable": "true", "retention.ms": "172800000"})
	fmt.Println(isok, err)
}

// create topic with the replicaAssign
// when you specified replicaAssignment ,the partitionNum and replications will be failed.
func TestCreateCustomTopicWithReplicaAssign(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()
	isok, err := admin.CreateCustomTopicWithReplicaAssign("gokafka-test-new",
		map[string]string{"unclean.leader.election.enable": "true", "retention.ms": "172800000"},
		map[int32][]int32{
			0: []int32{1, 2},
			1: []int32{2, 3},
			3: []int32{3, 1},
		})
	fmt.Printf("%v %+v\n", isok, err)
}

// update toipic with specifiy config
func TestUpdateTopicConfig(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()

	isok, err := admin.UpdateTopicConfig("--msg", map[string]string{"retention.mx": "888"}, false)

	fmt.Printf("%v %+v\n", isok, err)
}

// delete a topic
func TestDeleteTopic(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()
	for _, v := range []string{"xxb-test-123", "gokafka-test-new", "xxb-test-48"} {
		isok, err := admin.DeleteTopic(v)
		fmt.Println(isok, err)
	}
}

// get the broker base info
func TestGetBrokerIdList(t *testing.T) {
	admin := NewClusterAdmin([]string{"offline-kafka-1.soulapp.cn:9092"})
	controllerId, brokerIds, brokerInfos := admin.GetBrokerIdList()
	fmt.Println("controller:", controllerId)
	fmt.Println("broker list:", brokerIds)
	for _, v := range brokerInfos {
		fmt.Printf("id:%v\t\t broker:%v\t\n", v.BrokerId, v.BrokerIp)
	}
}

// describe the broker logdir
// return map[int32][]DescribeLogDirsResponseDirMetadata,error
// https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#DescribeLogDirsResponseDirMetadata
func TestDescribeLogDirs(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.16.162.63:9092"})
	defer admin.Close()
	_, brokerIds, _ := admin.GetBrokerIdList()
	fmt.Println(brokerIds)

	brokerLogs, brokerLogErr := admin.DescribeLogDirs(brokerIds)
	if brokerLogErr != nil {
		fmt.Printf("err:%v\n", brokerLogErr)
	}
	for id, logdirs := range brokerLogs {
		// 其实需要用broker_id找到对应的ip地址
		fmt.Printf("broker_id:%v\n", id)
		for _, logdir := range logdirs {
			for _, topicLog := range logdir.Topics {
				for _, topicPart := range topicLog.Partitions {
					fmt.Printf("logPath:%v\n", logdir.Path)
					fmt.Printf("pardId:%v-%v logSize:%v offsetLag:%v isTemp:%v\n", topicLog.Topic, topicPart.PartitionID, topicPart.Size, topicPart.OffsetLag, topicPart.IsTemporary)
				}
			}
		}
	}

}

// 获取指定broker id下的日志详情
func TestGetLogFromBrokers(t *testing.T) {
	admin := NewClusterAdmin([]string{"172.16.162.63:9092"})
	defer admin.Close()
	admin.GetLogFromBrokers([]int32{1})
}

func TestGetLogFromTopic(t *testing.T) {
	admin := NewClusterAdmin([]string{"bgbiao.top:9092"})
	//admin := NewClusterAdmin([]string{"172.29.203.62:9092"})
	defer admin.Close()
	for _, data := range admin.GetLogFromTopic("ablogflumelog") {
		fmt.Println(data.BrokerIp)
		fmt.Println(data.LogDatas)
	}

	for _, topicsDatas := range admin.GetLogFromTopics([]string{"ablogflumelog", "abresp"}) {
		fmt.Printf("topic:%v\n", topicsDatas.Name)
		for _, data := range topicsDatas.LogData {
			fmt.Println(data.BrokerIp)
			fmt.Println(data.LogDatas)
		}
	}
}

// 获取消费组列表
func TestListConsumerGroups(t *testing.T) {
	admin := NewClusterAdmin([]string{"bgbiao.top:9092"})
	defer admin.Close()

	// map[string]string
	consumerGroups, _ := admin.ListConsumerGroups()
	for k, v := range consumerGroups {
		fmt.Println(k, v)
	}

	consumerGroup, _ := admin.ListConsumerGroup()
	fmt.Println("####################################")
	fmt.Printf("%v\n", strings.Join(consumerGroup, "\n"))
}

// 获取消费组详情信息
func TestDescribeConsumerGroups(t *testing.T) {
	admin := NewClusterAdmin([]string{"bgbiao.top:9092"})
	defer admin.Close()

	abc, _ := admin.DescribeConsumerGroup([]string{"elasticsearch-kafka", "auto-commentV2-0"})

	for _, v := range abc {
		fmt.Println(v)
	}
}

// 获取消费组的offset消息
func TestListConsumerGroupOffSets(t *testing.T) {
	admin := NewClusterAdmin([]string{"bgbiao.top:9092"})
	defer admin.Close()

	topicPart := map[string][]int32{"test-topic": {1, 2}}
	topicPartOffset, _ := admin.ListConsumerGroupOffSets("elasticsearch-kafka", topicPart)
	fmt.Println(topicPartOffset)

	offset, _ := admin.ListConsumerGroupOffSet("active_action_user_groupid", "baseeventlog")
	fmt.Println(offset)
}
