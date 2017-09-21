package main

import (
	"os"
)
import "flag"
import "time"
import "log"

type loc struct{ x, y int }

type surroundingLiveCellCounter uint8

var liveCells map[loc]surroundingLiveCellCounter
var deadCellsNextToLiveCell map[loc]surroundingLiveCellCounter

func main() {
	var cycles uint
	flag.UintVar(&cycles, "ticks", 1, "Ticks/Cycles")
	var logInterval time.Duration
	flag.DurationVar(&logInterval, "interval", time.Second, "time between log status reports")
	var help bool
	flag.BoolVar(&help, "help", false, "display help/usage.")
	flag.BoolVar(&help, "h", false, "display help/usage.")
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

	var c uint
	doLog := time.NewTicker(logInterval)
	go func() {
		for _ = range doLog.C {
			log.Printf("\t#%d\talive:%d", c, len(liveCells))
		}
	}()

	if source.File == nil {
		var err error
		log.Printf("Loading:<<StdIn>>")
		liveCells, err = DecodeCellsFromImages(os.Stdin)
		if err != nil {
			panic(err)
		}

	} else {
		var err error
		log.Printf("Loading:%q", &source)
		liveCells, err = DecodeCellsFromImages(source)
		if err != nil {
			panic(err)
		}
	}
	log.Printf("\t#%d\talive:%d", 0, len(liveCells))

	for c = 0; c < cycles; c++ {
		if !tick() {
			log.Print("Unchanging")
			break
		}
	}

	log.Printf("\t#%d\talive:%d", c, len(liveCells))

	doLog.Stop()
	if sink.File == nil {
		log.Printf("Saving:<<StdOut>>")
		EncodeCellsAsImage(os.Stdout, liveCells)
	} else {
		log.Printf("Saving:%q", &sink)
		EncodeCellsAsImage(&sink, liveCells)
	}
}

func tick() (activity bool) {
	deadCellsNextToLiveCell = make(map[loc]surroundingLiveCellCounter)
	var count surroundingLiveCellCounter
	var l loc

	for l = range liveCells {
		liveCells[l] = atOffset(l, 1, 0) + atOffset(l, 1, 1) + atOffset(l, 0, 1) + atOffset(l, -1, 1) + atOffset(l, -1, 0) + atOffset(l, -1, -1) + atOffset(l, 0, -1) + atOffset(l, 1, -1)
	}

	for l, count = range liveCells {
		if count > 3 || count < 2 {
			delete(liveCells, l)
			activity = true
		}
	}

	for l, count = range deadCellsNextToLiveCell {
		if count == 3 {
			liveCells[l] = 0
			activity = true
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

