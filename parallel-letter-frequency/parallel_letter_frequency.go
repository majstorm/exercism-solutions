package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	c := make(chan FreqMap)
	fm := FreqMap{}

	for _, str := range l {
		go func(s string) {
			c <- Frequency(s)
		}(str)
	}

	for i := 0; i < len(l); i++ {
		for let, cnt := range <-c {
			fm[let] += cnt
		}
	}
	return fm
}
