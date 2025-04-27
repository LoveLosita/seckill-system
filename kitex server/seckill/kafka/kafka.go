package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"kitex-server/seckill/dao"
	"kitex-server/seckill/kitex_gen/seckill"
	"log"
	"strconv"
	"time"
)

func StartKafkaConsumer() {
	var emptyStatus seckill.Status
	// 1. 创建消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9094"},
		Topic:    "seckill_requests",
		GroupID:  "seckill_group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
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
		Addr:         kafka.TCP("localhost:9094"), // Kafka 服务器地址
		Topic:        "seckill_requests",          // Kafka 主题
		BatchSize:    100,                         //批次大小
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
