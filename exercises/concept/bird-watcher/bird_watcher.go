package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for i := range birdsPerDay {
        sum += birdsPerDay[i]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum := 0
    for i := (week - 1) * 7; i < len(birdsPerDay); i++ {
        if i <= week * 7 {
            sum += birdsPerDay[i]
        } else {
            break
        }
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i]++
    }
    return birdsPerDay
}
