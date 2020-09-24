## kafka

说明: kafka 库主要封装了[sarama](github.com/Shopify/sarama)库中的常用方法，来对常用的kafka管理操作和简单的生产消费进行接口模拟，用以实现一些基本的功能。不过侧重点在于整个集群或topic详情信息窥探。

在kafka中，主要分为三种客户端，分别是: `生产者客户端`，`消费者客户端`，`集群管理客户端`

### 集群管理客户端

集群管理客户端对外通过`ClusterAdmin`接口暴露，主要包含如下常见的方法实现:

- `CreateTopic`: Topic创建
- `ListTopics`: 列出topic列表
- `DescribeTopics`: 查看topic的详情信息
- `DeleteTopic`: topic删除
- `CreatePartitions`: 创建分区(增加分区时手动指定分区和分区分配策略 要求: version 1.0.0 or higher)
- `AlterPartitionReassignments`: 修改分区副本分配策略(用在副本在节点间的均衡 要求: 2.4.0.0 or higher)
- `ListPartitionReassignments`: 列出分区副本分配策略(正在进行分配的副本分配信息 要求: 2.4.0.0 or higher)
- `DeleteRecords`: 删除偏移量小于给定偏移量的记录(要求: 0.11.0.0 or higher)
- `DescribeConfig`: 查看配置文件(获取指定资源的配置实体，参数类型为`ConfigResource`，对于一些敏感配置信息没有开放 要求: 0.11.0.0 or higher)
- `AlterConfig`: 修改配置文件(使用默认参数来更新指定资源配置)
- `CreateACL`: 创建ACL访问规则
- `ListAcls`: 列出ACL规则
- `DeleteACL`: 删除ACL规则
- `ListConsumerGroups`: 列出集群消费组
- `DescribeConsumerGroups`: 查看指定消费组的详情信息
- `ListConsumerGroupOffsets`: 列出指定消费组可用的offset
- `DeleteConsumerGroup`: 删除消费组
- `DescribeCluster`: 查看集群详情信息(broker节点，controller节点等信息)
- `DescribeLogDirs`: 查看topic的各个日志信息(获取指定broker列表里日志目录信息)
- `Close`: 关闭管理员，并关闭基础客户端

`注意:` 因为通常可能企业内部使用的集群版本各不相同，因此对于有兼容版本的问题，高版本api接口未进行封装

**已封装功能**

- [X] 创建topic (指定分区和副本数进行创建)
- [X] 列出集群全部Topic列表
- [X] 查看集群信息(broker列表,controler节点)
- [X] 查看topic详情(配置信息，分区信息，ISR等信息)
- [X] 查看集群的topic日志信息(每个节点上每个topic-part的大小)
- [X] 列出topic的消费组
- [X] 查看topic的消费组信息(指定group的topic的消费者情况)
- [X] 查看指定实体的配置文件(broker,topic)
- [X] 修改指定资源实体的配置文件


### 生产者客户端

在[sarama](github.com/Shopify/sarama) 中，对于生产者相关的接口来讲，仅提供了异步的接口，即`AsyncProducer`接口。

AsyncProducer接口主要实现了以下几个功能:

- `AsyncClose()`: 异步客户端关闭(该方法会触发生产者的关闭，只有当错误和成功的channel都关闭时才会关闭，相当于优雅关闭，该方法在消费者侧可以很好的保证消息不会丢失)
- `Close()`: 该方法会等待缓冲在buffer里的消息已经被刷新时才关闭
- `Input()`: 这是一个`ProducerMessage`指针类型的写入channel，用来将用户希望的消息写进去
- `Successes()`: 这是成功输出的channel(当Return.Successes为true时，用户必须从这里读取来判断是否写入成功，否则生产者将deadlock，建议在一个单独的select语句中发送和读取消息)
- `Errors()`: 返回给用户的错误输出channel(当该channel满时，必须先读出来，否则会死锁，当然也可设置`Producer.Return.Errors`为false来避免返回错误信息)

**已实现的封装接口**

- [X] 同步生产(syncProducer): 从字符串直接生产
- [ ] 同步生产(syncProducer): 从标准输入读取字符串进行



### 消费者客户端

`注意:` 其实对于消费者而言，通常会分为两种消费方式，一种是消费者直接消费，另外一种是使用消费者组来消费一个指定topic的数据，后者不会影响整个topic的数据一致性，相当于一种订阅模式。

在[sarama](github.com/Shopify/sarama)中，对于消息消费的两个接口分别为`Consumer`和`ConsumerGroup`两个接口，但是因为消费组中会涉及到更多的消息控制，整体会比较负责一些，通常在测试消息过程中可以直接使用`Consumer`接口，而在对生产topic数据进行消费或处理时需要注意使用`ConsumerGroup`.

Consumer 接口实现的几个主要功能:

- `Topics()`: 从集群元数据中返回一组可用的topic列表信息
- `Partitions(topic string)`: 返回指定topic的排序分区列表，等同于`Client.Partitions()`
- `ConsumePartition(topic string,part int32,offset int64)`: 使用指定的topic/part的offset上创建一个`PartitionConsumer`，offset可以是字面的offset或者`OffsetNewest`或`OffsetOldest`的内置变量
- `HighWaterMarks()`: 返回每个topic和分区当前的高水位(分区间的一致性无法保证，因为高位水是单独标记的)
- `Close()`: 关闭消费者客户端

ConsumerGroup 接口实现的几个主要功能:

- `Consume(ctx context.Context,topics []string,handler ConsumerGroupHandler) error`: 为指定的topic列表加入一个消费者集群，然后通过`ConsumerGroupHandler`开启一个阻塞的`ConsumerGroupSession`，但是需要注意每个会话有一个完整的生命周期
- `Errors()`: 返回一个错误类型的读channel，如果想自定义一些错误类型可以设置`Consumer.Return.Errors=true` 然后读取出来自定义实现
- `Close()`: 停止消费组并分离正在运行的会话

[consumerGroupSession声明周期](https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#ConsumerGroup)


**封装已实现接口**

- [X] 消费者组消费，采用`gokafka`消费组
- [X] 指定消费者组
- [X] 指定消费位置`earliest` 和 `latest(default)`


