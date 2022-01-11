package main

import(
	"fmt"
	"log"
)

type Command [4]uint16

type (
	VmInt int
	VmFloat float64
	VmBool bool
	
	AsmBuilder struct{
		Ints []VmInt
		Floats []VmFloat
		Bools []VmBool
		Commands []Command
	}
)



var(
	System_Point int = 0
	CommandLine = []Command{}
	Int_Register   = []VmInt{}
	Float_Register = []VmFloat{}
	Bool_Register  = []VmBool{}
)

const(
	CLF int = iota
	RegisterA
	RegisterB
	RegisterC
)

func main(){
	
	if builder, error := AsmParser(AsmCode,); error != nil {
		log.Fatal(
			error,
		)
	}else{
		Int_Register   = builder.Ints
		Float_Register = builder.Floats
		Bool_Register  = builder.Bools
		CommandLine    = builder.Commands
	}
	
	
	if false{
		fmt.Println(Int_Register)
		fmt.Println(Bool_Register)
		fmt.Println(Float_Register)
	}
	
	MainLoop:
	for {
		
		if len(CommandLine) <= System_Point {
			fmt.Println("Main loop error : System point ",System_Point," is more then command count ", len(CommandLine))
			break MainLoop
		}
		
		switch CommandLine[System_Point][CLF]{
			
			// INT
			
			case Int_RegisterA_Move_ToRegisterB :
				Int_Register[int(CommandLine[System_Point][RegisterB])] = Int_Register[int(CommandLine[System_Point][RegisterA])]
			
			
			case Int_RegisterA_Add_RegisterB_ToRegisterC :
				Int_Register[int(CommandLine[System_Point][RegisterC])] = Int_Register[int(CommandLine[System_Point][RegisterA])] + Int_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Int_RegisterA_Sub_RegisterB_ToRegisterC :
				Int_Register[int(CommandLine[System_Point][RegisterC])] = Int_Register[int(CommandLine[System_Point][RegisterA])] - Int_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Int_RegisterA_Mull_RegisterB_ToRegisterC :
				Int_Register[int(CommandLine[System_Point][RegisterC])] = Int_Register[int(CommandLine[System_Point][RegisterA])] * Int_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Int_RegisterA_Div_RegisterB_ToRegisterC :
				Int_Register[int(CommandLine[System_Point][RegisterC])] = Int_Register[int(CommandLine[System_Point][RegisterA])] / Int_Register[int(CommandLine[System_Point][RegisterB])]
			
			
			case Int_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Int_Register[int(CommandLine[System_Point][RegisterA])] > Int_Register[int(CommandLine[System_Point][RegisterB])])
			
			case Int_RegisterA_Equality_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Int_Register[int(CommandLine[System_Point][RegisterA])] == Int_Register[int(CommandLine[System_Point][RegisterB])])
			
			case Int_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Int_Register[int(CommandLine[System_Point][RegisterA])] >= Int_Register[int(CommandLine[System_Point][RegisterB])])
			
			
			case Int_RegisterA_Convert_To_Bool_RegisterB :
				if Int_Register[int(CommandLine[System_Point][RegisterA])] == 0 {
					Bool_Register[int(CommandLine[System_Point][RegisterB])] = false
				}else{
					Bool_Register[int(CommandLine[System_Point][RegisterB])] = true
				}
			
			case Int_RegisterA_Convert_To_Float_RegisterB :
				Float_Register[int(CommandLine[System_Point][RegisterB])] = VmFloat(Int_Register[int(CommandLine[System_Point][RegisterA])])
			
			// FLOAT
			
			case Float_RegisterA_Move_ToRegisterB :
				Float_Register[int(CommandLine[System_Point][RegisterB])] = Float_Register[int(CommandLine[System_Point][RegisterA])]
			
			
			case Float_RegisterA_Add_RegisterB_ToRegisterC :
				Float_Register[int(CommandLine[System_Point][RegisterC])] = Float_Register[int(CommandLine[System_Point][RegisterA])] + Float_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Float_RegisterA_Sub_RegisterB_ToRegisterC :
				Float_Register[int(CommandLine[System_Point][RegisterC])] = Float_Register[int(CommandLine[System_Point][RegisterA])] - Float_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Float_RegisterA_Mull_RegisterB_ToRegisterC :
				Float_Register[int(CommandLine[System_Point][RegisterC])] = Float_Register[int(CommandLine[System_Point][RegisterA])] * Float_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Float_RegisterA_Div_RegisterB_ToRegisterC :
				Float_Register[int(CommandLine[System_Point][RegisterC])] = Float_Register[int(CommandLine[System_Point][RegisterA])] / Float_Register[int(CommandLine[System_Point][RegisterB])]
			
			
			case Float_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Float_Register[int(CommandLine[System_Point][RegisterA])] > Float_Register[int(CommandLine[System_Point][RegisterB])])
			
			case Float_RegisterA_Equality_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Float_Register[int(CommandLine[System_Point][RegisterA])] == Float_Register[int(CommandLine[System_Point][RegisterB])])
			
			case Float_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = (Float_Register[int(CommandLine[System_Point][RegisterA])] >= Float_Register[int(CommandLine[System_Point][RegisterB])])
			
			
			case Float_RegisterA_Convert_To_Bool_RegisterB :
				if Float_Register[int(CommandLine[System_Point][RegisterA])] == 0 {
					Bool_Register[int(CommandLine[System_Point][RegisterB])] = false
				}else{
					Bool_Register[int(CommandLine[System_Point][RegisterB])] = true
				}
			
			case Float_RegisterA_Convert_To_Int_RegisterB :
				Int_Register[int(CommandLine[System_Point][RegisterB])] = VmInt(Float_Register[int(CommandLine[System_Point][RegisterA])])
			
			// BOOL
			
			case Bool_RegisterA_Move_ToRegisterB :
				Bool_Register[int(CommandLine[System_Point][RegisterB])] = Bool_Register[int(CommandLine[System_Point][RegisterA])]
			
			
			case Bool_RegisterA_Add_RegisterB_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = Bool_Register[int(CommandLine[System_Point][RegisterA])] || Bool_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Bool_RegisterA_Mull_RegisterB_ToRegisterC :
				Bool_Register[int(CommandLine[System_Point][RegisterC])] = Bool_Register[int(CommandLine[System_Point][RegisterA])] && Bool_Register[int(CommandLine[System_Point][RegisterB])]
			
			case Bool_RegisterA_Invert_ToRegisterB :
				Bool_Register[int(CommandLine[System_Point][RegisterB])] = !Bool_Register[int(CommandLine[System_Point][RegisterA])]
			
			
			case Bool_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC :
				var a, b int8
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					a = 1
				}
				if Bool_Register[int(CommandLine[System_Point][RegisterB])] {
					b = 1
				}
				if a > b {
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = true
				}else{
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = false
				}
			
			case Bool_RegisterA_Equality_RegisterB_Bool_ToRegisterC :
				var a, b int8
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					a = 1
				}
				if Bool_Register[int(CommandLine[System_Point][RegisterB])] {
					b = 1
				}
				if a == b {
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = true
				}else{
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = false
				}
			
			case Bool_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC :
				var a, b int8
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					a = 1
				}
				if Bool_Register[int(CommandLine[System_Point][RegisterB])] {
					b = 1
				}
				if a >= b {
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = true
				}else{
					Bool_Register[int(CommandLine[System_Point][RegisterC])] = false
				}
			
			
			case Bool_RegisterA_Convert_To_Int_RegisterB :
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					Int_Register[int(CommandLine[System_Point][RegisterB])] = 1
				}else{
					Int_Register[int(CommandLine[System_Point][RegisterB])] = 0
				}
				
			
			case Bool_RegisterA_Convert_To_Float_RegisterB :
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					Float_Register[int(CommandLine[System_Point][RegisterB])] = 1.
				}else{
					Float_Register[int(CommandLine[System_Point][RegisterB])] = 0.
				}
			
			
			case System_EndPoint :
				break MainLoop
			
			case System_Point_Set_Int_RegisterA :
				System_Point = int(
					Int_Register[int(CommandLine[System_Point][RegisterA])]-1,
				)
			
			case System_If_RegisterA_Set_RegisterB_Else_RegisterC :
				if Bool_Register[int(CommandLine[System_Point][RegisterA])] {
					System_Point = int(
						Int_Register[int(CommandLine[System_Point][RegisterB])]-1,
					)
				}else{
					System_Point = int(
						Int_Register[int(CommandLine[System_Point][RegisterC])]-1,
					)
				}
			
			
			case Print_Int_RegisterA :
				fmt.Println(
					Int_Register[int(CommandLine[System_Point][RegisterA])],
				)
			case Print_Float_RegisterA :
				fmt.Println(
					Float_Register[int(CommandLine[System_Point][RegisterA])],
				)
			case Print_Bool_RegisterA :
				fmt.Println(
					Bool_Register[int(CommandLine[System_Point][RegisterA])],
				)
			
			default :
				fmt.Print("Main loop error : indefine command index : ", CommandLine[System_Point][0])
		}
		
		System_Point++
	}
}