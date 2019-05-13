/*Conway's Game Of Life

Uses maps/dicts for storage, so speed/memory use proportional to number of live cells.

in/out storage using png files.

save sequences of images for making movies. 
*/

// example arguments:  -i="./media/glider" -pipeMovie -wrap -size=200 -ticks=500


// TODO add terminal gocli display?

package main

import (
	"os"
)
import "flag"
import "time"
import "log"
import "strconv"
import "path/filepath"

import "github.com/splace/fsflags"

type loc struct{ x, y int }

type surroundingLiveCellCounter uint8

var liveCells map[loc]surroundingLiveCellCounter
var deadCellsNextToLiveCell map[loc]surroundingLiveCellCounter

var limit int
var wrap bool

func main() {
	var source fsflags.FileValue
	flag.Var(&source, "i", "source for the starting cell pattern, encoded in PNG image.(default:<Stdin>)")
	flag.Var(&source, "input", flag.Lookup("i").Usage)
	var sink fsflags.CreateFileValue
	flag.Var(&sink, "o", "file for encoding result cell pattern, PNG image.(default:Stdout)")
	flag.Var(&sink, "output", flag.Lookup("o").Usage)
	var cycles uint
	flag.UintVar(&cycles, "ticks", 1, "Ticks/Cycles")
	var logInterval time.Duration
	flag.DurationVar(&logInterval, "interval", time.Second, "time between log status reports")
	var wrap bool
	flag.BoolVar(&wrap, "w", false, "sets arena to (s)ize.")
	flag.BoolVar(&wrap, "wrap", false, flag.Lookup("w").Usage)
	var pipeMovie bool
	flag.BoolVar(&pipeMovie, "p", false,"send snapshot images to Stdout.")
	flag.BoolVar(&pipeMovie, "pipeMovie",false, flag.Lookup("p").Usage)
	var movie fsflags.NewOverwriteDirValue
	flag.Var(&movie, "m", "directory for snapshot frames, PNG images.")
	flag.Var(&movie, "movie", flag.Lookup("m").Usage)
	var size uint
	flag.UintVar(&size, "s", 32,"size of snapshots.")
	flag.UintVar(&size, "size",32,flag.Lookup("s").Usage)
	limit=int(size/2)
	var ticksSnapshot uint
	flag.UintVar(&ticksSnapshot, "f", 1,"ticks for each snapshot image.")
	flag.UintVar(&ticksSnapshot, "frameTicks", 1,flag.Lookup("f").Usage)
	var help bool
	flag.BoolVar(&help, "h", false, "display help/usage.")
	flag.BoolVar(&help, "help", false, flag.Lookup("h").Usage)
	var logToo fsflags.CreateFileValue
	flag.Var(&logToo, "log", "progress log destination.(default:Stderr)")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if pipeMovie {
		movie.File=os.Stdout
	}

	if logToo.File == nil {
		logToo.File = os.Stderr
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

	if movie.File!=nil {
		log.Printf("Saving snapshot images to :%q", &movie)
	}


	log.Printf("\t#%d\talive:%d", 0, len(liveCells))


	switch movie.File{
	case nil:
		for c = 0; c < cycles; c++ {
			if anyChanged:=tick(); !anyChanged {
				log.Print("Unchanging")
				break
			}
		}
	case os.Stdout:
		for c = 0; c < cycles; c++ {
			if anyChanged:=tick(); !anyChanged {
				log.Print("Unchanging")
				break
			}
			if c%ticksSnapshot ==0 {
				EncodeCellsAsSizedImage(os.Stdout, liveCells,size)							
			}
		}
	default:
		for c = 0; c < cycles; c++ {
			if anyChanged:=tick(); !anyChanged {
				log.Print("Unchanging")
				break
			}
			if c%ticksSnapshot ==0 {
				frameFile,err:=os.Create(filepath.Join(movie.File.Name(),strconv.FormatUint(uint64(c),10)+".png"))
				if err!=nil{
					log.Printf("\t#%d\tUnable to save frame:%s", c,err)
				}else{
					EncodeCellsAsSizedImage(frameFile, liveCells,size)		
					frameFile.Close()		
				}
			}
		}
	}

	log.Printf("\t#%d\talive:%d", c, len(liveCells))

	doLog.Stop()
	if sink.File == nil {
		log.Printf("Saving:<<StdOut>>")
		if wrap {
			EncodeCellsAsSizedImage(os.Stdout, liveCells,size)				
		}else{
			EncodeCellsAsImage(os.Stdout, liveCells)
		}
		
	} else {
		log.Printf("Saving:%q", &sink)
		if wrap {
			EncodeCellsAsSizedImage(&sink, liveCells,size)				
		}else{
			EncodeCellsAsImage(&sink, liveCells)
		}
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
	if wrap {
		l.x+=limit
		l.y+=limit
		l.x %=limit
		l.y %=limit
		l.x-=limit
		l.y-=limit
	}
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

