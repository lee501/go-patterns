package builder

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	director := Director{}

	//car
	car := &Car{}
	director.SetBuilder(car)
	director.Construct()
	vehicle := director.builder.GetVehicle()
	//vehicle = car.GetVehicle()
	fmt.Println(vehicle)
	if vehicle.Wheels != 4 {
		t.Errorf("car wheels must be 4, but get %d\n", vehicle.Wheels)
	}
	if vehicle.Seats != 4 {
		t.Errorf("car seats must be 4, but get %d\n", vehicle.Seats)
	}
	if vehicle.Structure != "Car" {
		t.Errorf("vehicle structure must be Car, but get %s\n", vehicle.Structure)
	}
}
