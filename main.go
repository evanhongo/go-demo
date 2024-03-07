package main

import (
	"go-demo/service"
)

func main() {
	// 肉品初始化
	meats := make([]string, 22)
	for i := 0; i < 10; i++ {
		meats[i] = "牛肉"
	}

	for i := 10; i < 17; i++ {
		meats[i] = "豬肉"
	}

	for i := 17; i < 22; i++ {
		meats[i] = "雞肉"
	}

	// 創建一個處理肉品加工的工廠，並分配 5 位員工給工廠
	workers := []string{"A", "B", "C", "D", "E"}
	svc := service.NewMeatProcessingFactory(workers)

	// 工廠開始處理肉品
	svc.ProcessMeat(meats)

}
