package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int = 0

	// for i := 0; i < len(birdsPerDay); i++ {
	// 	sum += birdsPerDay[i]
	// }

	for _, v := range birdsPerDay {
		sum += v
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	var sum int = 0

	for i, v := range birdsPerDay {
		if i >= (week-1)*7 && i < week*7 {
			sum += v
		}
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
}
