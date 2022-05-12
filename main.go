package main

func main() {

}

/*
 - main functions CANNOT be imported by other packages.
 - The "exported" identifiers defined in a main package CANNOT be accessed by or leaked to other packages.
 - It therefore does not matter (ito isolation/encapsulation) whether identifiers here are exported or not.
 - It is however suggested to follow the "private-is-private-public-is-public" convention.
*/

// MakeParams is an exported function, but cannot be accessed outside this main package.
func MakeParams() Params {
	return Params{
		P1: 1,
		P2: 2.0,
	}
}

// Params is a value (DTO).
// Values only contain data, values do not contain logic.
// Values are treated as "immutable", so their fields are public.
type Params struct {
	P1 int
	P2 float64
}
