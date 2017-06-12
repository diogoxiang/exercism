package clock

const testVersion = 4

// You can find more details and hints in the test file.

type Clock int // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	c := Clock((hour*60 + minute) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
