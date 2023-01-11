package day14

type reindeer struct {
	name             string
	speed            int
	flyDuration      int
	flyDurationLeft  int
	restDuration     int
	restDurationLeft int
	distanceTraveled int
	points           int
}

func (r *reindeer) race(sec int) {
	if r.flyDurationLeft > 0 {
		r.fly()
	} else {
		if r.restDurationLeft > 0 {
			r.rest()
		} else {
			r.resetLeftDurations()
			r.fly()
		}
	}
}

func (r *reindeer) fly() {
	r.distanceTraveled += r.speed
	r.flyDurationLeft--
}

func (r *reindeer) rest() {
	r.restDurationLeft--
}

func (r *reindeer) resetLeftDurations() {
	r.flyDurationLeft = r.flyDuration
	r.restDurationLeft = r.restDuration
}
