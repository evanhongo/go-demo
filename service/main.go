package service

import (
	"fmt"
	"sync"
	"time"
)

type MeatProcessingFactory struct {
	workerCh chan string
}

func (f *MeatProcessingFactory) ProcessMeat(meats []string) {
	var (
		wg = sync.WaitGroup{}
	)

	// 遍歷肉品清單，每個肉品都要 從 buffer channel(workerCh) 中找到一個 "空閒的工人" 處理肉品
	for _, meat := range meats {
		// 這裡如果阻塞，就代表所有工人都在忙，找不到空閒的
		worker := <-f.workerCh
		wg.Add(1)

		// 找到有空閒的工人，分派肉品給他
		go func(meat string) {
			// 處理完肉品就把工人放回 buffer channel
			defer func() {
				wg.Done()
				f.workerCh <- worker
			}()

			f.processMeat(worker, meat)
		}(meat)
	}

	// 等待所有肉品處理完畢
	wg.Wait()
}

func (f *MeatProcessingFactory) processMeat(worker string, meat string) {
	fmt.Printf("%v 在 %v 取得%v\n", worker, time.Now().Format("2006-01-02 15:04:05"), meat)
	// 模擬肉品處理，牛肉1秒，豬肉2秒，雞肉3秒
	switch meat {
	case "牛肉":
		time.Sleep(time.Second)
	case "豬肉":
		time.Sleep(2 * time.Second)
	case "雞肉":
		time.Sleep(3 * time.Second)
	default:
	}
	fmt.Printf("%v 在 %v 處理完%v\n", worker, time.Now().Format("2006-01-02 15:04:05"), meat)
}

func NewMeatProcessingFactory(workers []string) *MeatProcessingFactory {
	f := &MeatProcessingFactory{}

	// 初始化工廠，用 buffered channel 來存 "工人"，buffer 上限為5個
	f.workerCh = make(chan string, len(workers))
	for _, worker := range workers {
		f.workerCh <- worker
	}
	return f
}
