module golang-course/cmd/hw_02

go 1.16

replace golang-course/cmd/hw_02/fibonacci/dp => /fibonacci/dp

replace golang-course/cmd/hw_02/fibonacci/recursive => /fibonacci/recursive

require (
	golang-course/cmd/hw_02/fibonacci/dp v0.0.0-00010101000000-000000000000 // indirect
	golang-course/cmd/hw_02/fibonacci/recursive v0.0.0-00010101000000-000000000000 // indirect
)
