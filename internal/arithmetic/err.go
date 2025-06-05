package arithmetic

type arError struct {
	place int
	msg   string
}

func (a *arError) Error() string {
	return a.msg
}

func NewarError(whatExpected string, place int, text []byte) error {
	return &arError{
		place: place,
		msg:   " ❌ ERROR: " + whatExpected + "\n" + string(text[:place-1]) + "[❗️Here]" + string(text[place-1:]),
	}
}
