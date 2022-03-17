package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

const maxNames = 676000

var nameAvailability = make(map[string]bool)
var randSrc = rand.NewSource(time.Now().UnixNano())

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(nameAvailability) == maxNames {
		return "", errors.New("No available names")
	}

	randomNumber := fmt.Sprintf("%03d", rand.New(randSrc).Intn(1000))
	randomLetters := fmt.Sprintf("%c%c", rune(65+rand.New(randSrc).Intn(26)), rune(65+rand.New(randSrc).Intn(26)))
	if ok, _ := nameAvailability[randomLetters+randomNumber]; !ok {
		r.name = randomLetters + randomNumber
		nameAvailability[r.name] = true
	} else {
		r.name, _ = r.Name()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	delete(nameAvailability, r.name)
	r.name = ""
}
