package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	total := v.Total
	var num = make(chan float64, 2)
	defer close(num)
	go getMemory(num)
	for {
		select {
		case data, ok := <-num:
			if !ok {
				fmt.Println("Channel has been closed")
				num = nil
			} else {
				ShowMemoryInOneLine(data, float64(total))
			}
		default:
			break
		}
	}
}
func ShowMemoryInOneLine(used, total float64) {
	fmt.Print(fmt.Sprintf("|Used : %0.2f | Total: %0.2f|\r", used, total))
}
func getMemory(channel chan float64) {
	for true {
		v, err := mem.VirtualMemory()
		if err != nil {
			_ = fmt.Errorf("ошибка получения данных о памяти")
		}
		used := float64(v.Used) / float64(1024) / float64(1024) / float64(1024)
		channel <- used
	}
}
