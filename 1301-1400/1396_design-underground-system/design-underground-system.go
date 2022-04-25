package main

// Difficulty:
// Medium

// Tags:
// Hash

type UndergroundSystem struct {
	customers  map[int]*customer
	statistics map[string]map[string][]int
}

type customer struct {
	startStation string
	startTime    int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		customers:  map[int]*customer{},
		statistics: map[string]map[string][]int{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.customers[id]; ok {
		return
	}
	this.customers[id] = &customer{stationName, t}
	if _, ok := this.statistics[stationName]; !ok {
		this.statistics[stationName] = map[string][]int{}
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.customers[id]; !ok {
		return
	}
	customer := this.customers[id]
	delete(this.customers, id)
	if _, ok := this.statistics[customer.startStation][stationName]; !ok {
		this.statistics[customer.startStation][stationName] = make([]int, 2)
	}
	this.statistics[customer.startStation][stationName][0]++
	this.statistics[customer.startStation][stationName][1] += t - customer.startTime
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if _, ok := this.statistics[startStation]; !ok {
		return 0
	}
	if _, ok := this.statistics[startStation][endStation]; !ok {
		return 0
	}
	cnt := this.statistics[startStation][endStation][0]
	sum := this.statistics[startStation][endStation][1]
	return float64(sum) / float64(cnt)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
