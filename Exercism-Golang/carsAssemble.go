package main

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	rateSuccess := successRate / 100
	return rateSuccess * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	rateSuccess := successRate / 100
	perhour := rateSuccess * float64(productionRate)
	return int(perhour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTen := carsCount / 10
	remainders := carsCount % 10
	cost := (groupOfTen*95000 + remainders*10000)
	return uint(cost)
}

fmt.Prin
