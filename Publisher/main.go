package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	RecordedAt  time.Time `json:"recorded_at"`
}

func main() {
	// MQTT配置（连接到您的公网IP）
	opts := MQTT.NewClientOptions().AddBroker("tcp://8.222.186.212:1883")
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic("MQTT连接失败: " + token.Error().Error())
	}
	defer client.Disconnect(250)

	rand.Seed(time.Now().UnixNano())

	for {
		// 生成随机温湿度数据
		data := SensorData{
			Temperature: 20 + rand.Float64()*15, // 20~35°C
			Humidity:    40 + rand.Float64()*30, // 40~70%
			RecordedAt:  time.Now().UTC(),       // 当前时间
		}

		// 序列化为JSON
		payload, _ := json.Marshal(data)

		// 发布到MQTT Topic
		topic := "sensor/data"
		token := client.Publish(topic, 1, false, payload)
		token.Wait()
		fmt.Printf("[%s] 发送数据: %s\n", time.Now().Format("15:04:05"), string(payload))

		time.Sleep(5 * time.Second) // 每5秒发送一次
	}
}
