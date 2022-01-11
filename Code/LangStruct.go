package main

func NewSintSystem()(func(string)uint16, *map[string]uint16){
	namespace := make(map[string]uint16)
	return func(name string)uint16{
		if _,ok := namespace[name]; ok {
			return 0xffff
		}
		var size = uint16(len(namespace))
		namespace[name] = size
		return size
	}, &namespace
}
