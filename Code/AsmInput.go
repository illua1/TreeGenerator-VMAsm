package main

import(
	//"strconv"
)

var AsmCode = `
Data{
	Int{10;1;0;1;4},
	Float{},
	Bool{1},
},
Body{
	ISub 0 1 0;IPrint 0;IMTE 0 3 0;If 0 2 4;End
}`
/*
type AsmPrinter AsmBuilder

func ZoneFuctory(open, close string)func(...string)string{
	return func(body ...string)(ret string){
		ret = open
		for i:=range body {
			ret += body[i]+","
		}
		return ret+close
	}
}

var(
	IntZone = ZoneFuctory("Int:{","}")
	FloatZone = ZoneFuctory("Float:{","}")
	BoolZone = ZoneFuctory("Bool:{","}")
	DataZone = ZoneFuctory("Data:{","}")
	CommandsZone = ZoneFuctory("Body:{","}")
)

func(asmP AsmPrinter)String()(ret string){
	ret = DataZone(
		IntZone(
			asmP.IntsToString(),
		),
		FloatZone(
			asmP.FloatsToString(),
		),
		BoolZone(
			asmP.BoolsToString(),
		),
	)
	ret += CommandsZone(
		asmP.Commands,
	)
	return
}

func(asmP AsmPrinter)IntsToString()(ret string){
	for i := range asmP.Ints {
		ret += strconv.Itoa(asmP.Ints[i])+","
	}
	ret = ret[:len(ret-len(":"))]
	return
}
func(asmP AsmPrinter)FloatsToString()(ret string){
	for i := range asmP.Floats {
		ret += strconv.FormatFloat(asmP.Floats[i], 'E', -1, 64)+","
	}
	ret = ret[:len(ret-len(":"))]
	return
}
func(asmP AsmPrinter)BoolsToString()(ret string){
	for i := range asmP.Bools {
		ret += strconv.FormatBool(asmP.Bools[i])+","
	}
	ret = ret[:len(ret-len(":"))]
	return
}

*/