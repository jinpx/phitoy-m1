package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 定义有效时间段
	validPeriods := []struct {
		start string
		end   string
	}{
		{"10:00", "12:30"},
		{"13:30", "19:05"},
	}

	// 通道用于键盘事件
	keyPressDetected := make(chan bool)

	// 启动全局按键监听
	go startKeyListener(keyPressDetected)

	// 主循环：模拟按键
	interval := 25 * time.Second // 每 25 秒检查一次
	for {
		select {
		case <-keyPressDetected:
			// 如果检测到键盘或鼠标事件，跳过本次模拟操作
		case <-time.After(interval):
			// 检查是否在有效时间段内
			if isWithinValidPeriod(validPeriods) {
				// 在有效时间段内，执行模拟操作
				executeKeyActions()
			} else {
				fmt.Println("当前时间不在有效时间段内，跳过模拟操作...")
			}
		}
	}
}

// isWithinValidPeriod 检查当前时间是否在指定的时间段内
func isWithinValidPeriod(periods []struct {
	start string
	end   string
}) bool {
	now := time.Now()
	for _, period := range periods {
		// 将时间段解析为当天的时间
		startTime, _ := time.ParseInLocation("15:04", period.start, now.Location())
		startTime = time.Date(now.Year(), now.Month(), now.Day(), startTime.Hour(), startTime.Minute(), 0, 0, now.Location())

		endTime, _ := time.ParseInLocation("15:04", period.end, now.Location())
		endTime = time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), 0, 0, now.Location())

		// 检查当前时间是否在时间段内
		if now.After(startTime) && now.Before(endTime) {
			return true
		}
	}
	return false
}

// startKeyListener 启动全局按键监听
func startKeyListener(keyPressDetected chan bool) {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		// 监听按键和鼠标点击事件
		if ev.Kind == hook.KeyDown || ev.Kind == hook.MouseDown {
			keyPressDetected <- true
		}
	}
}

// executeKeyActions 模拟上下键操作
func executeKeyActions() {
	randomMoves := rand.Intn(126) + 25 // 随机生成 25 到 150 次
	fmt.Printf("Performing key actions: pressing 'down' %d times, followed by pressing 'up' 30 times.\n", randomMoves)
	for i := 0; i < randomMoves; i++ {
		robotgo.KeyTap("down")
		time.Sleep(100 * time.Millisecond)
	}
	for i := 0; i < 30; i++ {
		robotgo.KeyTap("up")
		time.Sleep(100 * time.Millisecond)
	}
}
