package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
)

type SensorData struct {
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	RecordedAt  time.Time `json:"recorded_at"`
}

func main() {
	// PostgreSQL连接配置（根据您的服务器信息修改）
	connStr := "host=111.111.111 port=5432 user=postgres password=111 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal("数据库Ping失败: ", err)
	}
	log.Println("成功连接到PostgreSQL!")

	// MQTT订阅配置
	opts := MQTT.NewClientOptions().AddBroker("tcp://8.222.186.212:1883")
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT连接失败: ", token.Error())
	}

	// 订阅Topic
	token := client.Subscribe("sensor/data", 1, func(client MQTT.Client, msg MQTT.Message) {
		var data SensorData
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			log.Printf("JSON解析失败: %v (原始数据: %s)", err, msg.Payload())
			return
		}

		// 写入PostgreSQL
		_, err := db.Exec(
			`INSERT INTO sensor_data (temperature, humidity, recorded_at) 
			 VALUES ($1, $2, $3)`,
			data.Temperature, data.Humidity, data.RecordedAt,
		)
		if err != nil {
			log.Printf("数据库写入失败: %v", err)
		} else {
			log.Printf("数据已存储: %.1f°C, %.1f%%", data.Temperature, data.Humidity)
		}
	})
	token.Wait()

	// 保持运行
	log.Println("等待传感器数据...")
	select {}
}
