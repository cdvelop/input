package input

func Phone() *number {
	// Phone `pattern="^[0-9]{7,11}$"`
	return Number(`min="7"`, `max="11"`)
}
