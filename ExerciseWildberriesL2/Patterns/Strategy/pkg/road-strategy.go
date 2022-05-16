package pkg

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 /* произвольная цифра среднего времени*/ * trafficJam
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Trrafic jam:[%d] Total:[%d] Total time[%d] min \n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}
