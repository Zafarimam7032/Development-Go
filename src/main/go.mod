module github.com/src/main

go 1.18

replace example.com/src/packageFolder => ../packageFolder

require (
	example.com/src/CondtionalStatement v0.0.0-00010101000000-000000000000
	example.com/src/loop v0.0.0-00010101000000-000000000000
	example.com/src/packageFolder v0.0.0-00010101000000-000000000000
	github.com/src/variables v0.0.0-00010101000000-000000000000
)

replace github.com/src/variables => ../variables

replace example.com/src/loop => ../loop

replace example.com/src/CondtionalStatement => ../CondtionalStatement
