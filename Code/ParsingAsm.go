package main

import(
	"fmt"
	"strings"
	"strconv"
)

func AsmParser(lines string)(AsmBuilder, error){
	
	ret := AsmBuilder{}
	
	lines = Trim(Trim(Trim(lines, " ", 1), string(10),0),`	`, 0)
	
	wood := ParsToWood(lines)
	
	var errors ParsingErrors
	
	var(
		parsBody = func(text string){
			for _, line := range strings.Split(text, ";") {
				components := strings.Split(line, " ")
				if len(components) == 0 {
					continue
				}
				com, err := commandTest(components)
				
				if err != nil {
					errors.Push(err)
					return
				}
				ret.Commands = append(
					ret.Commands,
					com,
				)
			}
			return
		}
		parsIntList = func(text string){
			for _, element := range strings.Split(text, ";") {
				value, err := strconv.Atoi(element)
				errors.Push(err)
				ret.Ints = append(ret.Ints, VmInt(value))
			}
			return
		}
		parsBoolList = func(text string){
			for _, element := range strings.Split(text, ";") {
				value, err := strconv.ParseBool(element)
				errors.Push(err)
				ret.Bools = append(ret.Bools, VmBool(value))
			}
			return
		}
		parsFloatList = func(text string){
			for _, element := range strings.Split(text, ";") {
				value, err := strconv.ParseFloat(element, 64)
				errors.Push(err)
				ret.Floats = append(ret.Floats, VmFloat(value))
			}
			return
		}
	)
	
	{
		testersData := map[string]func(e Char){
			"Int" : func(e Char){
				parsIntList(Trim(Trim(e.Body, " ", 1), `\n`,0))
			},
			"Float" : func(e Char){
				parsFloatList(Trim(Trim(e.Body, " ", 1), `\n`,0))
			},
			"Bool" : func(e Char){
				parsBoolList(Trim(Trim(e.Body, " ", 1), `\n`,0))
			},
		}
		
		testersMain := map[string]func(e Char){
			"Data" : func(e Char){
				TestStr(e.Link, testersData)
			},
			"Body" : func(e Char){
				parsBody(Trim(Trim(e.Body, " ", 1), `\n`,0))
			},
		}
		
		TestZones(
			wood,
			testersMain,
		)
	}
	
	return ret, errors.Test()
}

type ParsingErrors []error

func(pr ParsingErrors)Error()(ret string){
	ret = "["
	for i := range pr{
		ret += "<"+pr[i].Error()+">"
	}
	ret += "]"
	return
}

func(pr ParsingErrors)Push(add error)(ParsingErrors){
	return append(pr, add)
}

func(pr ParsingErrors)Test()(error){
	if len(pr)!= 0 {
		return pr
	}
	var count int
	for i := range pr {
		if pr[i] != nil {count++}
	}
	if count != 0 {
		return pr
	}
	return nil
}

func TestZones( e Char, tester map[string]func(Char) ){
	var ok bool
	if e.Type {
		if len(e.Link) > 0 {
			if !(e.Link[0].Type){
				if TestStr(e.Link, tester) {
					ok = true
				}
			}
		}
		if !ok {
			for i := range e.Link {
				TestZones(*e.Link[i], tester)
			}
		}
	}
}

func TestStr( elements Line, tester map[string]func(Char) ) bool{
	for k, _ := range tester{
		for e, _ := range elements{
			if strings.Contains(elements[e].Body, k) {
				if len(elements) > e+1 {
					(tester)[k](*(elements[e+1]))
					delete(tester, k)
				}
			}
		}
	}
	if len(tester) == 0 {
		return true
	}
	return false
}

func Trim(text string, f string, max int)(ret string){
	var te = []rune(text)
	var b int
	for i := range te {
		if string(te[i]) == f {
			if b < max {
				ret += string(te[i])
			}
			b++
		}else{
			b = 0
			ret += string(te[i])
		}
	}
	return
}

func commandTest(arg []string)(Command, error){
	comId, ok := (*CommandsNames)[arg[CLF]]
	if !ok {
		return Command{}, fmt.Errorf("Undefine command")
	}
	var rg1, rg2, rg3 int
	if len(arg) > RegisterA{
		rg1, _ = strconv.Atoi(arg[RegisterA])
	}
	if len(arg) > RegisterB{
		rg2, _ = strconv.Atoi(arg[RegisterB])
	}
	if len(arg) > RegisterC{
		rg3, _ = strconv.Atoi(arg[RegisterC])
	}
	
	return Command{
		comId,
		uint16(rg1),
		uint16(rg2),
		uint16(rg3),
	}, nil
}