package validations

type Validation interface {
	Func() string
	Args() map[string]interface{}
}
