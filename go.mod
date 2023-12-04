module github.com/cdvelop/input

go 1.20

require (
	github.com/cdvelop/model v0.0.76
	github.com/cdvelop/timetools v0.0.25
)

require github.com/cdvelop/strings v0.0.7 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/strings => ../strings
