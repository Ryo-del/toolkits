package ping

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/go-ping/ping"
)

type Config struct {
	Ping PingConfig `json:"ping"`
}

type PingConfig struct {
	Packages int    `json:"packages"`
	Interval int    `json:"interval"`
	Source   string `json:"source"`
}

var config PingConfig
var mu sync.Mutex

// Результат последнего пинга
var lastStats *ping.Statistics

func init() {
	file, err := os.Open("internal/settings/main.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}
	config = cfg.Ping
}

func RunPing() {
	mu.Lock()
	defer mu.Unlock()

	pinger, err := ping.NewPinger(config.Source)
	if err != nil {
		lastStats = nil
		return
	}

	pinger.Count = config.Packages
	err = pinger.Run()
	if err != nil {
		lastStats = nil
		return
	}
	stats := pinger.Statistics()
	lastStats = stats
}

func GetPingStatus() string {
	mu.Lock()
	defer mu.Unlock()

	if lastStats == nil {
		return "⚠️ Error"
	}
	if lastStats.PacketsRecv > 0 {
		return "✅ Alive"
	}
	return "❌ Unreachable"
}

/*func GetLoss() int {
	mu.Lock()
	defer mu.Unlock()

	if lastStats == nil {
		return -1
	}
	return int(lastStats.PacketLoss)
}*/

func GetSource() string {
	return config.Source
}
func GetInterval() int {
	return config.Interval
}
