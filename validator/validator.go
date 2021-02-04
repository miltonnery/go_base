package validator

// ok represents types capable of validating
// themselves.
type ok interface {
	OK() error
}
