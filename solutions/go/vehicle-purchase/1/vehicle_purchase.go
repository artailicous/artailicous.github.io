package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}else{
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var recommendedVehicle string

	// Compare option1 and option2 lexicographically.
	// The vehicle that is "smaller" (comes first alphabetically) is chosen.
	if option1 < option2 {
		recommendedVehicle = option1
	} else {
		recommendedVehicle = option2
	}

	// Format and return the decision string.
	// The problem examples show that the chosen vehicle is always part of the output string,
	// e.g., "Toyota Corolla is clearly the better choice."
	return recommendedVehicle + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	/*
		For a rough estimate, assume if the vehicle is less than 3 years old,
		it costs 80% of the original price it had when it was brand new.
		If it is at least 10 years old, it costs 50%.
		If the vehicle is at least 3 years old but less than 10 years,
		it costs 70% of the original price.
	*/

	// Check for the *less than 3 years old* condition (80%)
	if age < 3 {
		return originalPrice * 0.8
	}

	// Check for the *at least 10 years old* condition (50%)
	// Note: We use >= 10 for "at least 10 years old"
	if age >= 10 {
		return originalPrice * 0.5
	}

	// The remaining case is when age is *at least 3 years old but less than 10* (70%)
	// Since the previous checks handled age < 3 and age >= 10,
	// any age reaching this point must be in the range [3, 10), so no explicit check is needed.
	return originalPrice * 0.7
}