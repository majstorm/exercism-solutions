package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, num := range birdsPerDay {
		sum += num
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[0+(week-1)*7 : 7+(week-1)*7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	k := birdsPerDay
	for i := 0; i < len(k); i++ {
		k[i] = k[i] + 1
		if i%2 == 1 {
			k[i] = k[i] - 1
		}
	}
	return k
}
