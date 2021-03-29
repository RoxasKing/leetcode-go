package main

/*
  Design a parking system for a parking lot. The parking lot has three kinds of parking spaces: big, medium, and small, with a fixed number of slots for each size.

  Implement the ParkingSystem class:
    1. ParkingSystem(int big, int medium, int small) Initializes object of the ParkingSystem class. The number of slots for each parking space are given as part of the constructor.
    2. bool addCar(int carType) Checks whether there is a parking space of carType for the car that wants to get into the parking lot. carType can be of three kinds: big, medium, or small, which are represented by 1, 2, and 3 respectively. A car can only park in a parking space of its carType. If there is no space available, return false, else park the car in that size space and return true.

  Example 1:
    Input
      ["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
      [[1, 1, 0], [1], [2], [3], [1]]
    Output
      [null, true, true, false, false]
    Explanation
      ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
      parkingSystem.addCar(1); // return true because there is 1 available slot for a big car
      parkingSystem.addCar(2); // return true because there is 1 available slot for a medium car
      parkingSystem.addCar(3); // return false because there is no available slot for a small car
      parkingSystem.addCar(1); // return false because there is no available slot for a big car. It is already occupied.

  Constraints:
    1. 0 <= big, medium, small <= 1000
    2. carType is 1, 2, or 3
    3. At most 1000 calls will be made to addCar

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/design-parking-system
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
