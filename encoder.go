package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"
)


var liveColor = color.White
var emptyColor = color.Black 

func DecodeCellsFromImages(r io.Reader)(c map[loc]surroundingLiveCellCounter,err error){
	img,_,err:=image.Decode(r)
	if err!=nil{return}
	c = make(map[loc]surroundingLiveCellCounter)
	SetCells(img,&c)
	// look for any image encodings following on from the same Reader, abort but dont return errors.
	var eerr error
	for {
		img,_,eerr=image.Decode(r)
		if eerr!=nil{return}
		SetCells(img,&c)
	}
	return
}

func SetCells(i image.Image,cells *map[loc]surroundingLiveCellCounter){
	ib:=i.Bounds()
	emptyColor:=i.ColorModel().Convert(emptyColor)
	for x:=ib.Min.X;x<ib.Max.X;x++{
		for y:=ib.Min.Y;y<ib.Max.Y;y++{
			if i.At(x,y) != emptyColor {
				(*cells)[loc{x,y}]=0
			}		
		}
	}
}


func EncodeCellsAsImage(w io.Writer,c map[loc]surroundingLiveCellCounter)(err error){
	return png.Encode(w, RGBAImage{NewDepictionAll(c, liveColor, emptyColor)}) 
}

type Depiction struct {
	Cells map[loc]surroundingLiveCellCounter
	size  image.Rectangle
	in,out color.Color
}

func NewDepictionAll(cs map[loc]surroundingLiveCellCounter, below, above color.Color) Depiction {
	limits:=image.ZR
	for l:=range cs{
		if l.x>=limits.Max.X {limits.Max.X=l.x+1}
		if l.x<limits.Min.X {limits.Min.X=l.x}
		if l.y>=limits.Max.Y {limits.Max.Y=l.y+1}
		if l.y<limits.Min.Y {limits.Min.Y=l.y}
	}
	return Depiction{cs,limits, below, above}
}

func (i Depiction) Bounds() image.Rectangle {
	return i.size
}

func (i Depiction) At(xp, yp int) color.Color {
	if _,in:=i.Cells[loc{xp,yp}];in{
		return i.in
	}
	return i.out
}




// a Depictor is an image.Image without a colormodel, so is more general.
// by being embedded in one of the helper wrappers you get an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// RGBA depiction wrapper
type RGBAImage struct {
	Depictor
}

func (i RGBAImage) ColorModel() color.Model { return color.RGBAModel }

// gray depiction wrapper.
type GrayImage struct {
	Depictor
}

func (i GrayImage) ColorModel() color.Model { return color.GrayModel }

// plan9 paletted, depiction wrapper.
type Plan9PalettedImage struct {
	Depictor
}

func (i Plan9PalettedImage) ColorModel() color.Model { return color.Palette(palette.Plan9) }

// WebSafe paletted, depiction wrapper.
type WebSafePalettedImage struct {
	Depictor
}

func (i WebSafePalettedImage) ColorModel() color.Model { return color.Palette(palette.WebSafe) }



