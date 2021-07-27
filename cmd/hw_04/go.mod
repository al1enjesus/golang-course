module golang-course/cmd/hw_04

go 1.16

replace golang-course/cmd/hw_04/shape => /shape

replace golang-course/cmd/hw_04/shape/circle => /shape/circle

replace golang-course/cmd/hw_04/shape/rectangle => /shape/rectangle

require (
	golang-course/cmd/hw_04/shape v0.0.0-00010101000000-000000000000 // indirect
	golang-course/cmd/hw_04/shape/circle v0.0.0-00010101000000-000000000000 // indirect
	golang-course/cmd/hw_04/shape/rectangle v0.0.0-00010101000000-000000000000 // indirect
)
