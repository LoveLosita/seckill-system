package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"kitex-server/seckill/dao"
	"kitex-server/seckill/kitex_gen/seckill"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"kitex-server/seckill/dao"
	"kitex-server/seckill/kitex_gen/seckill"
	"log"
	"strconv"
	"time"
)

func StartKafkaConsumer() {
	// 2. 处理消息
	for {
		//2.1.获取消息
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Failed to read message: %v\n", err)
			continue
		}
		var request map[string]string
		if err := json.Unmarshal(m.Value, &request); err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}
		productID := request["productID"]
		orderID := request["orderID"]
		intProductID, err := strconv.ParseInt(productID, 10, 64)
		//2.2.扣减数据库中的库存
		status := dao.DeductStockInMysql(intProductID)
		//2.3.写入订单状态（下单成功/失败）到数据库
		if status != emptyStatus {
			status = dao.AddOrderStatusToMysql(orderID, intProductID, "failed")
			log.Fatal("dao.DeductStockInMysql error: ", status)
		} else {
			status = dao.AddOrderStatusToMysql(orderID, intProductID, "success")
		}
	}
}

func AddMsgToKafka(productID int64, orderID string) error {
	// 1. 创建生产者
	writer := &kafka.Writer{
		Addr:         kafka.TCP("host.docker.internal:19094"), // Kafka 服务器地址
		Topic:        "seckill_requests",                      // Kafka 主题
		BatchSize:    100,                                     //批次大小
		BatchTimeout: 100 * time.Millisecond,
	}
	// 2. 发送消息
	// 创建秒杀请求消息
	message := map[string]string{
		"productID": strconv.FormatInt(productID, 10),
		"orderID":   orderID,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// 发送到 Kafka
	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: messageBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}
	return nil
}
*/

func StartKafkaConsumer() {
	var emptyStatus seckill.Status
	// 配置信息
	topic := "my-topic"
	groupID := "my-consumer-group"
	brokers := []string{"172.23.75.194:9094"}
	// 创建SASL认证机制
	mechanism := plain.Mechanism{
		Username: "user",
		Password: "password",
	}
	// 配置Dialer (不使用TLS，因为我们使用的是SASL_PLAINTEXT)
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	// 创建Reader配置
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         brokers,
		Topic:           topic,
		GroupID:         groupID,           // 消费者组ID
		MinBytes:        10e3,              // 10KB 最小批处理大小
		MaxBytes:        10e6,              // 10MB 最大批处理大小
		MaxWait:         1 * time.Second,   // 最长等待时间
		StartOffset:     kafka.FirstOffset, // 从最早的消息开始（可选用 kafka.LastOffset 从最新的开始）
		ReadLagInterval: -1,                // 禁用滞后报告
		Dialer:          dialer,            // 使用带SASL的dialer
	})
	// 捕获中断信号以优雅退出
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// 创建上下文，允许我们控制消费循环
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 在单独的goroutine中处理信号
	go func() {
		sig := <-sigchan
		fmt.Printf("捕获到信号: %v, 正在关闭消费者...\n", sig)
		cancel()
	}()
	fmt.Println("开始消费消息，按 Ctrl+C 停止...")
	// 消费消息循环
	for {
		select {
		case <-ctx.Done():
			fmt.Println("上下文已取消，退出消费循环")
			if err := reader.Close(); err != nil {
				log.Fatalf("关闭reader失败: %v", err)
			}
			return
		default:
			// 读取消息
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				// 检查是否因为上下文取消而中断
				if ctx.Err() != nil {
					continue
				}
				log.Printf("读取消息失败: %v", err)
				continue
			}
			// 处理消息
			fmt.Printf("收到消息: 主题=%s, 分区=%d, 偏移量=%d, 键=%s, 值=%s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			// 这里可以添加您的业务逻辑来处理消息
			//2.1.获取消息
			var request map[string]string
			if err := json.Unmarshal(m.Value, &request); err != nil {
				fmt.Printf("Failed to unmarshal message: %v\n", err)
				continue
			}
			productID := request["productID"]
			orderID := request["orderID"]
			intProductID, err := strconv.ParseInt(productID, 10, 64)
			//2.2.扣减数据库中的库存
			status := dao.DeductStockInMysql(intProductID)
			//2.3.写入订单状态（下单成功/失败）到数据库
			if status != emptyStatus {
				status2 := dao.AddOrderStatusToMysql(orderID, intProductID, "failed")
				if status2 != emptyStatus {
					fmt.Println("dao.AddOrderStatusToMysql error: ", status2)
				}
			} else {
				status2 := dao.AddOrderStatusToMysql(orderID, intProductID, "success")
				if status2 != emptyStatus {
					fmt.Println("dao.AddOrderStatusToMysql error: ", status)
				}
			}
			// kafka-go 自动处理提交偏移量，除非您使用了CommitMessages方法手动控制
		}
	}
}

func AddMsgToKafka(productID int64, orderID string) error {
	topic := "my-topic"
	partition := 0
	// 创建SASL认证机制（使用用户名和密码）
	mechanism := plain.Mechanism{
		Username: "user",
		Password: "password",
	}
	// 创建无TLS的Dialer（因为我们配置的是SASL_PLAINTEXT）
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	// 连接至Kafka集群的Leader节点
	conn, err := dialer.DialLeader(context.Background(), "tcp", "172.23.75.194:9094", topic, partition)
	if err != nil {
		return fmt.Errorf("failed to dial leader:%s", err)
	}
	// 设置发送消息的超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	message := map[string]string{
		"productID": strconv.FormatInt(productID, 10),
		"orderID":   orderID,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	// 发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: messageBytes},
	)
	if err != nil {
		return fmt.Errorf("failed to write messages:%s", err)
	}
	fmt.Println("write messages success")

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return nil
}
