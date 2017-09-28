package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	)

var liveColor = color.White
var emptyColor = color.Black 

func DecodeCellsFromImages(r io.Reader)(c map[loc]surroundingLiveCellCounter,err error){
	img,_,err:=image.Decode(r)
	if err!=nil{return}
	c = make(map[loc]surroundingLiveCellCounter)
	SetCells(img,img.Bounds(),&c)
	// look for any image encodings, following on, in the same Reader, but dont return any error from these.
	// allows easy overlaying several grids
	var eerr error
	for {
		img,_,eerr=image.Decode(r)
		if eerr!=nil{return}
		SetCells(img,img.Bounds(),&c)
	}
	return
}

func SetCells(i image.Image,ib image.Rectangle,cells *map[loc]surroundingLiveCellCounter){
	emptyColor:=i.ColorModel().Convert(emptyColor)
	for x:=ib.Min.X;x<ib.Max.X;x++{
		for y:=ib.Min.Y;y<ib.Max.Y;y++{
			if i.At(x,y) != emptyColor {
				(*cells)[loc{x,y}]=0
			}		
		}
	}
}


func EncodeCellsAsImage(w io.Writer,c map[loc]surroundingLiveCellCounter) error{
	return png.Encode(w, Image{c, CellsBounds(c),liveColor, emptyColor,color.RGBAModel}) 
}

func EncodeCellsAsSizedImage(w io.Writer,c map[loc]surroundingLiveCellCounter,size uint) error{
	return png.Encode(w, Image{c, image.Rect(-int(size/2),-int(size/2),int(size/2),int(size/2)),liveColor, emptyColor,color.RGBAModel}) 
}

func CellsBounds(c map[loc]surroundingLiveCellCounter)image.Rectangle{
	limits:=image.ZR
	for l:=range c{
		if l.x>=limits.Max.X {limits.Max.X=l.x+1}
		if l.x<limits.Min.X {limits.Min.X=l.x}
		if l.y>=limits.Max.Y {limits.Max.Y=l.y+1}
		if l.y<limits.Min.Y {limits.Min.Y=l.y}
	}
	return limits
}


type Image struct {
	Cells map[loc]surroundingLiveCellCounter
	size  image.Rectangle
	in,out color.Color
	colorModel color.Model
}

func (i Image) Bounds() image.Rectangle {
	return i.size
}

func (i Image) At(xp, yp int) color.Color {
	if _,in:=i.Cells[loc{xp,yp}];in{
		return i.in
	}
	return i.out
}


func (i Image) ColorModel() color.Model { 
	return i.colorModel
}

