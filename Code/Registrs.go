package main

import(
	"fmt"
)

func IntRegistrTest(a int)(VmInt, error){
	if len(Int_Register) > a {
		return Int_Register[a], nil
	}
	return 0, fmt.Errorf("Int Registr: index overflov")
}

func FloatRegistrTest(a int)(VmFloat, error){
	if len(Float_Register) > a {
		return Float_Register[a], nil
	}
	return 0., fmt.Errorf("Float Registr: index overflov")
}

func BoolRegistrTest(a int)(VmBool, error){
	if len(Bool_Register) > a {
		return Bool_Register[a], nil
	}
	return false, fmt.Errorf("Bool Registr: index overflov")
}