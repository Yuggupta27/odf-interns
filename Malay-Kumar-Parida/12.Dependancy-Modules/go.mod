module github.com/malayparida2000/Dependancy-Modules

go 1.17

//go mod tidy pulls the required packages and creates go.sum with the checksums

//requireed modules

//replace moudules

//exclude modules

// Can upgrade packages by go get and smenatic version tags

// go get rsc.io/sampler@v1.3.1

//go list -m all lists out all dependancies

//go mod tidy removes any unused packages

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
