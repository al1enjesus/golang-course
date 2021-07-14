module golang-course/cmd/hw_02

go 1.16

replace golang-course/cmd/hw_02/fibonacci/fibo_dynamic => /fibonacci/fibo_dynamic

replace golang-course/cmd/hw_02/fibonacci/fibo_recursive => /fibonacci/fibo_recursive

replace golang-course/cmd/hw_02/fibonacci/fibo_simplest => /fibonacci/fibo_simplest

require (
	golang-course/cmd/hw_02/fibonacci/fibo_dynamic v0.0.0-00010101000000-000000000000 // indirect
	golang-course/cmd/hw_02/fibonacci/fibo_recursive v0.0.0-00010101000000-000000000000 // indirect
	golang-course/cmd/hw_02/fibonacci/fibo_simplest v0.0.0-00010101000000-000000000000 // indirect
)
