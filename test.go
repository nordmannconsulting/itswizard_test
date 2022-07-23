package itswizard_test

import (
	"itswizard_test/basic_structs"
	"itswizard_test/greeting"
)

func SayHelloWithValue(name string) (string, error) {
	x := basic_structs.NewTestStructure(name, 17)
	x.Increment()
	tmp, err := x.GetName()
	if err != nil {
		return "", err
	}
	return greeting.Greet(tmp), nil
}
