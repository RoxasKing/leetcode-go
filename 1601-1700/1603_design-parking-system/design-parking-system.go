package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.big == 0 || carType == 2 && this.medium == 0 || carType == 3 && this.small == 0 {
		return false
	}
	switch carType {
	case 1:
		this.big--
	case 2:
		this.medium--
	case 3:
		this.small--
	}
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
