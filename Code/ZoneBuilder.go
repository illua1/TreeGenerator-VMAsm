package main

// https://go.dev/play/p/kOKhCHfhkXS
import (
	"fmt"
	"strconv"
)

type Char struct {
	Body  string
	Type  bool
	Types int
	Link  Line
}

func (c Char) String() (ret string) {
	ret = "{'" + c.Body + "'," + strconv.FormatBool(c.Type) + ", " + strconv.Itoa(c.Types) + " <" + c.Link.String() + ">" + "}"
	return
}

type Line []*Char

func (l Line) String() (ret string) {
	for i := range l {
		ret += l[i].String() + ", "
	}
	return
}

type zoneTest uint8

const (
	ZonePass zoneTest = iota
	ZoneOpen
	ZoneClose
)

func NewZone(open, close string, id int) func(string) (zoneTest, int) {
	return func(str string) (zoneTest, int) {
		if str == open {
			return ZoneOpen, id
		}
		if str == close {
			return ZoneClose, id
		}
		return ZonePass, -1
	}
}

var (
	MathZone   = NewZone("(", ")", 0)
	StructZone = NewZone("{", "}", 1)
	TextZone   = NewZone("'", "'", 2)
)
/*
func main() {
	var str = "ABD{123:assd,fss(dsed,sfe33,411),5235,wfe{sdvsdv}} + (123, 123,123)"
	var Wood = Pars(str)
	fmt.Println(str)
	fmt.Println(Wood)
}
*/

func ParsToWood(text string) Char {
	var strR = []rune(text)
	ret := make(chan rune)
	go func() {
		for i := range strR {
			ret <- strR[i]
		}
		close(ret)
	}()
	var r, _ = ParsDo(ret, -1)
	return r
}

func ParsDo(rets <-chan rune, types int) (Char, int) {
	var ret Char
	var buffer []rune
	for r := range rets {
		zoneProp, index := Test(r)
		if !((zoneProp == ZoneOpen) || (zoneProp == ZoneClose)) {
			buffer = append(buffer, r)
		} else {
			if zoneProp == ZoneOpen {
				var f Char
				f.Body = string(buffer)
				buffer = []rune("")
				ret.Link = append(ret.Link, &f)
				ret.Type = true
				element, r := ParsDo(rets, index)
				if r != index {
					fmt.Println("SINTAX ERROR")
				}
				ret.Link = append(ret.Link, &element)
			}
			if zoneProp == ZoneClose {
				if types != index {
					fmt.Println("SINTAX ERROR")
				}
				ret.Types = types
				ret.Body = string(buffer)
				return ret, types
			}
		}
	}
	ret.Body = string(buffer)
	ret.Types = types
	return ret, -1
}

func Test(ch rune) (zoneTest, int) {
	if r, i := MathZone(string(ch)); r != ZonePass {
		return r, i
	}
	if r, i := TextZone(string(ch)); r != ZonePass {
		return r, i
	}
	if r, i := StructZone(string(ch)); r != ZonePass {
		return r, i
	}
	return ZonePass, -1
}