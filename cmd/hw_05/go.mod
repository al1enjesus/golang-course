module golang-course/cmd/hw_05

go 1.16

replace golang-course/cmd/hw_05/shape => /shape

replace golang-course/cmd/hw_05/shape/circle => /shape/circle

replace golang-course/cmd/hw_05/shape/rectangle => /shape/rectangle

require (
	golang-course/cmd/hw_05/shape v0.0.0-00010101000000-000000000000
	golang-course/cmd/hw_05/shape/circle v0.0.0-00010101000000-000000000000
	golang-course/cmd/hw_05/shape/rectangle v0.0.0-00010101000000-000000000000
)
