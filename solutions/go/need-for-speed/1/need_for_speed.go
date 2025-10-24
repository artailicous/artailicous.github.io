package speed

// Car represents a remote controlled car with battery, speed, and distance tracking
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// Track represents a race track with a specific distance
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// Check if there's enough battery to drive
	if car.battery < car.batteryDrain {
		return car // Car doesn't move if not enough battery
	}

	// Update car state after driving
	car.battery -= car.batteryDrain
	car.distance += car.speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// Calculate remaining distance to finish the track
	remainingDistance := track.distance - car.distance

	// If already finished, return true
	if remainingDistance <= 0 {
		return true
	}

	// Calculate how many drives are needed to finish
	drivesNeeded := (remainingDistance + car.speed - 1) / car.speed // Ceiling division

	// Check if we have enough battery for all needed drives
	return car.battery >= drivesNeeded*car.batteryDrain
}
