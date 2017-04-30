package validations

// ValidationProvider provides validation behaviour
type Provider interface {
	Validate() error
}
