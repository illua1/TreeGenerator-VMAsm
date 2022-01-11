package main

var(
	// MainNameSpace
	Lang, CommandsNames = NewSintSystem()
	// INT
		//INT_MOVE
			// B = A
	Int_RegisterA_Move_ToRegisterB = Lang(IntSet)
		//INT_MATH
			// C = A+B
	Int_RegisterA_Add_RegisterB_ToRegisterC = Lang(IntAdd)
			// C = A-B
	Int_RegisterA_Sub_RegisterB_ToRegisterC = Lang(IntSub)
			// C = A*B
	Int_RegisterA_Mull_RegisterB_ToRegisterC = Lang(IntMull)
			// C = A/B
	Int_RegisterA_Div_RegisterB_ToRegisterC = Lang(IntDiv)
	
		//INT_COMPAIR
			// C = A > B
	Int_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC = Lang(IntMoreThen)
			// C = A == B
	Int_RegisterA_Equality_RegisterB_Bool_ToRegisterC = Lang(IntEquality)
			// C = A >= B
	Int_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC = Lang(IntMoreThen_Equality)
	
		//INT_CAST_TO
			// B = Bool(A)
	Int_RegisterA_Convert_To_Bool_RegisterB = Lang(IntToBool)
            // B = Float(A)
	Int_RegisterA_Convert_To_Float_RegisterB = Lang(IntToFloat)
	
	// FLOAT
		//FLOAT_MOVE
			// B = A
	Float_RegisterA_Move_ToRegisterB = Lang(FloatSet)
		//FLOAT_MATH
			// C = A+B
	Float_RegisterA_Add_RegisterB_ToRegisterC = Lang(FloatAdd)
			// C = A-B
	Float_RegisterA_Sub_RegisterB_ToRegisterC = Lang(FloatSub)
			// C = A*B
	Float_RegisterA_Mull_RegisterB_ToRegisterC = Lang(FloatMull)
			// C = A/B
	Float_RegisterA_Div_RegisterB_ToRegisterC = Lang(FloatDiv)
	
		//FLOAT_COMPAIR
			// C = A > B
	Float_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC = Lang(FloatMoreThen)
			// C = A == B
	Float_RegisterA_Equality_RegisterB_Bool_ToRegisterC = Lang(FloatEquality)
			// C = A >= B
	Float_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC = Lang(FloatMoreThen_Equality)
	
		//FLOAT_CAST_TO
			// B = Bool(A)
	Float_RegisterA_Convert_To_Bool_RegisterB = Lang(FloatToBool)
			// B = Int(A)
	Float_RegisterA_Convert_To_Int_RegisterB = Lang(FloatToInt)
	
	// BOOL
		//BOOL_MOVE
			// B = A
	Bool_RegisterA_Move_ToRegisterB = Lang(BoolSet)
	
		//BOOL_MATH
			// C = A+B
	Bool_RegisterA_Add_RegisterB_ToRegisterC = Lang(BoolAdd)
			// C = A*B
	Bool_RegisterA_Mull_RegisterB_ToRegisterC = Lang(BoolMull)
			// B = !A
	Bool_RegisterA_Invert_ToRegisterB = Lang(BoolInvert)
	
		//BOOL_COMPAIR
			// C = A > B
	Bool_RegisterA_MoreThen_RegisterB_Bool_ToRegisterC = Lang(BoolMoreThen)
			// C = A == B
	Bool_RegisterA_Equality_RegisterB_Bool_ToRegisterC = Lang(BoolEquality)
			// C = A >= B
	Bool_RegisterA_MoreThen_Equality_RegisterB_Bool_ToRegisterC = Lang(BoolMoreThen_Equality)
	
		//BOOL_CAST_TO
			// B = Bool(A)
	Bool_RegisterA_Convert_To_Int_RegisterB = Lang(BoolToInt)
			// B = Float(A)
	Bool_RegisterA_Convert_To_Float_RegisterB = Lang(BoolToFloat)
	
	// SYSTEM
		//SYSTEM_STOP
			// STOP
	System_EndPoint = Lang(EndPoint)
		//SYSTEM_JMP
			// JMP -> A
	System_Point_Set_Int_RegisterA = Lang(JMP)
		//SYSTEM_IF_ELSE_JMP
			// If A (BOOL) {JMP -> B(INT)}else{GMP -> C(INT)}
	System_If_RegisterA_Set_RegisterB_Else_RegisterC = Lang(If)
	
	/*
	//DATA
		//DATA_ALLOC_INT
	Data_Alloc_Int_IndexToIntRegisterA
		//DATA_COUNT_OF_INT_ALLOC
	Data_Alloc_Int_CountRegisterA_IndexToIntRegisterB
	
		//DATA_ALLOC_FLOAT
	Data_Alloc_FLOAT_IndexToIntRegisterA
		//DATA_COUNT_OF_FLOAT_ALLOC
	Data_Alloc_FLOAT_CountRegisterA_IndexToIntRegisterB
	
		//DATA_ALLOC_BOOL
	Data_Alloc_Bool_IndexToIntRegisterA
		//DATA_COUNT_OF_BOOL_ALLOC
	Data_Alloc_Bool_CountRegisterA_IndexToIntRegisterB
	*/
	
		//PRINT
	Print_Int_RegisterA = Lang(IntPrint)
	Print_Float_RegisterA = Lang(FloatPrint)
	Print_Bool_RegisterA = Lang(BoolPrint)
)