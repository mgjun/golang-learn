package day01

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"time"
)

func main() {
	// 1. 压测时长&速率
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 4 * time.Second

	// 2. 压测接口
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:9100/",
	})

	// 3. 启动压测并收集结果
	var metrics vegeta.Metrics
	attacker := vegeta.NewAttacker()
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}

	// 4. 处理结果
	metrics.Close()

	// 5. 打印感兴趣的指标
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
