package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionFloat := float64(productionRate)
    successFraction := successRate / 100.0
    WorkingCarsPerHour := productionFloat * successFraction
    return WorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	WorkingCarsPerHour := int(CalculateWorkingCarsPerHour(productionRate, successRate))
    WorkingCarsPerMinute := WorkingCarsPerHour / 60
    return WorkingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    PriceForTen := (carsCount / 10) * 95000
    PriceForRemainder := (carsCount % 10) * 10000
    PriceForTotal := uint(PriceForTen + PriceForRemainder)
    return PriceForTotal
	
}
