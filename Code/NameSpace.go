package main

func NP(pref string)(func(string)string){
	return func(name string)string{
		return pref+name
	}
}

var(
	intPrefixName = NP("I")
	IntSet = intPrefixName("Set")
	IntAdd = intPrefixName("Add")
	IntSub = intPrefixName("Sub")
	IntMull = intPrefixName("Mull")
	IntDiv = intPrefixName("Div")
	IntMoreThen = intPrefixName("MT")
	IntEquality = intPrefixName("E")
	IntMoreThen_Equality = intPrefixName("MTE")
	IntToBool = intPrefixName("ToBool")
	IntToFloat = intPrefixName("ToFloat")
	IntPrint = intPrefixName("Print")
)

var(
	floatPrefixName = NP("F")
	FloatSet = floatPrefixName("Set")
	FloatAdd = floatPrefixName("Add")
	FloatSub = floatPrefixName("Sub")
	FloatMull = floatPrefixName("Mull")
	FloatDiv = floatPrefixName("Div")
	FloatMoreThen = floatPrefixName("MT")
	FloatEquality = floatPrefixName("E")
	FloatMoreThen_Equality = floatPrefixName("MTE")
	FloatToBool = floatPrefixName("ToBool")
	FloatToInt = floatPrefixName("ToInt")
	BoolPrint = floatPrefixName("Print")
)

var(
	boolPrefixName = NP("B")
	BoolSet = boolPrefixName("Set")
	BoolAdd = boolPrefixName("Add")
	BoolMull = boolPrefixName("Mull")
	BoolInvert = boolPrefixName("Invert")
	BoolMoreThen = boolPrefixName("MT")
	BoolEquality = boolPrefixName("E")
	BoolMoreThen_Equality = boolPrefixName("MTE")
	BoolToFloat = boolPrefixName("ToFloat")
	BoolToInt = boolPrefixName("ToInt")
	FloatPrint = boolPrefixName("Print")
)

var(
	systemPrefixName = NP("")
	EndPoint = systemPrefixName("End")
	JMP = systemPrefixName("JMP")
	If = systemPrefixName("If")
)