package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	return kind == "car" || kind == "truck"
	// panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {

	choosen := option2
	if option1 < option2 {
		choosen = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", choosen)

	//panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	rate := 70.0
	if age < 3 {
		rate = 80.0
	}

	if age >= 10 {
		rate = 50.0
	}
	return originalPrice * rate / 100
	//	panic("CalculateResellPrice not implemented")
}
