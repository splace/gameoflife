package main

import (
	"os"
)
import "flag"

type loc struct{ x, y int }

type surroundingLiveCellCounter uint8

var liveCells map[loc]surroundingLiveCellCounter  
var deadCellsNextToLiveCell map[loc]surroundingLiveCellCounter

func main() {
	var cycles uint
	flag.UintVar(&cycles, "ticks", 1, "Ticks/Cycles")
	var help bool
	flag.BoolVar(&help,"help", false, "display help/usage.")
	flag.BoolVar(&help,"h", false, "display help/usage.")
	var source fileValue
    flag.Var(&source, "i", "source for the starting cell pattern, encoded in PNG image.(default:<Stdin>)")
    flag.Var(&source, "input", "source for the starting cell pattern, encoded in PNG image.(default:<Stdin>)")
	var sink newFileValue
    flag.Var(&sink, "o", "file for encoding result cell pattern, PNG image.(default:Stdout)")
    flag.Var(&sink, "output", "file for encoding result cell pattern, PNG image.(default:Stdout)")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	
    if source.File==nil {
    	var err error
		liveCells,err=DecodeCellsFromImages(os.Stdin)
		if 	err!=nil{panic(err)}
	}else{
    	var err error
		liveCells,err=DecodeCellsFromImages(source)
		if 	err!=nil{panic(err)}
	}	
	
	var t uint
	for t = 1; t < cycles ; t++ {
		if !tick(){break}
	}
	
    if sink.File==nil {
    	var err error
		EncodeCellsAsImage(os.Stdout, liveCells) 
		if 	err!=nil{panic(err)}
	}else{
    	var err error
		EncodeCellsAsImage(sink, liveCells) 
		if 	err!=nil{panic(err)}
	}	
}


func tick() (activity bool){
	deadCellsNextToLiveCell = make(map[loc]surroundingLiveCellCounter)
	var count surroundingLiveCellCounter
	var l loc
	
	for l = range liveCells {
		liveCells[l] = atOffset(l, 1, 0) + atOffset(l, 1, 1) + atOffset(l, 0, 1) + atOffset(l, -1, 1) + atOffset(l, -1, 0) + atOffset(l, -1, -1) + atOffset(l, 0, -1) + atOffset(l, 1, -1)
	}

	for l, count = range liveCells {
		if count > 3 || count < 2 {
			delete(liveCells, l)
			activity =true
		}
	}

	for l, count = range deadCellsNextToLiveCell {
		if count == 3 {
			liveCells[l] = 0
			activity =true
		}
	}
	return
}

func atOffset(l loc, dx, dy int8) surroundingLiveCellCounter {
	l.x += int(dx)
	l.y += int(dy)
	if _, in := liveCells[l]; in {
		return 1
	}
	if c, in := deadCellsNextToLiveCell[l]; in {
		deadCellsNextToLiveCell[l] = c + 1
	} else {
		deadCellsNextToLiveCell[l] = 1
	}
	return 0
}


