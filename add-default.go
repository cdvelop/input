package input

const DefaultValidateFunction = `userFormTyping(this)`

type sourceData interface {
	SourceData() map[string]string
}
