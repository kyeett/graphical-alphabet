package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten"

	"github.com/peterhellberg/gfx"
)

var header3 = `
&nbsp; | &nbsp; | &nbsp; |
:---------:|:----------:|:---------:|
`

var header4 = `
&nbsp; | &nbsp; |&nbsp; | &nbsp; |
:---------:|:----------:|:---------:|:----------:|
`

var header5 = `
&nbsp; | &nbsp; |&nbsp; | &nbsp; |&nbsp; |
:---------:|:----------:|:---------:|:----------:|:---------:|
`

func porterDuffComposition(s, d image.Image) error {
	src, _ := ebiten.NewImageFromImage(s, ebiten.FilterDefault)
	dst, _ := ebiten.NewImageFromImage(d, ebiten.FilterDefault)

	modes := []struct {
		name string
		op   ebiten.CompositeMode
	}{
		{"Src", ebiten.CompositeModeCopy},
		{"Atop", ebiten.CompositeModeSourceAtop},
		{"Over", ebiten.CompositeModeSourceOver},
		{"In", ebiten.CompositeModeSourceIn},
		{"Out", ebiten.CompositeModeSourceOut},

		{"Dest", ebiten.CompositeModeDestination},
		{"DestAtop", ebiten.CompositeModeCopy},
		{"DestOver", ebiten.CompositeModeDestinationOver},
		{"DestIn", ebiten.CompositeModeDestinationIn},
		{"DestOut", ebiten.CompositeModeDestinationOut},

		{"Clear", ebiten.CompositeModeClear},
		{"Xor", ebiten.CompositeModeXor},
		{"Lighter", ebiten.CompositeModeLighter},
	}

	var table string
	for i, o := range modes {
		filename := fmt.Sprintf("examples/%s.png", o.name)
		tmp, _ := ebiten.NewImageFromImage(dst, ebiten.FilterDefault)
		op := &ebiten.DrawImageOptions{}
		op.CompositeMode = o.op
		if err := tmp.DrawImage(src, op); err != nil {
			gfx.Log("couldn't draw image %s: %s", filename, err)
		}
		err := gfx.SavePNG(filename, tmp)
		if err != nil {
			return errors.Errorf("failed to save image:", err)
		}

		table += fmt.Sprintf("![example:%s](%s)<br>[%s](%s) |", o.name, filename, o.name, filename)
		if i%5 == 5-1 {
			table += "\n"
		}
	}

	ioutil.WriteFile("EXAMPLES.md", []byte(header5+table), 0644)
	return nil
}

func transformation(d, s image.Image) error {
	src, _ := ebiten.NewImageFromImage(s, ebiten.FilterDefault)
	dst, _ := ebiten.NewImageFromImage(d, ebiten.FilterDefault)

	flipHorizontal := &ebiten.DrawImageOptions{}
	flipHorizontal.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	flipHorizontal.GeoM.Scale(-1, 1)
	flipHorizontal.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))

	flipVertical := &ebiten.DrawImageOptions{}
	flipVertical.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	flipVertical.GeoM.Scale(1, -1)
	flipVertical.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))

	rotate90 := &ebiten.DrawImageOptions{}
	rotate90.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotate90.GeoM.Rotate(math.Pi / 2)
	rotate90.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))

	rotate180 := &ebiten.DrawImageOptions{}
	rotate180.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotate180.GeoM.Rotate(math.Pi)
	rotate180.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))

	rotate270 := &ebiten.DrawImageOptions{}
	rotate270.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotate270.GeoM.Rotate(3 * math.Pi / 2)
	rotate270.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))

	rotateDest90 := &ebiten.DrawImageOptions{}
	rotateDest90.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotateDest90.GeoM.Rotate(math.Pi / 2)
	rotateDest90.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))
	rotateDest90.CompositeMode = ebiten.CompositeModeDestinationOver

	rotateDest180 := &ebiten.DrawImageOptions{}
	rotateDest180.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotateDest180.GeoM.Rotate(math.Pi)
	rotateDest180.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))
	rotateDest180.CompositeMode = ebiten.CompositeModeDestinationOver

	rotateDest270 := &ebiten.DrawImageOptions{}
	rotateDest270.GeoM.Translate(-float64(src.Bounds().Dx()/2), -float64(src.Bounds().Dy()/2))
	rotateDest270.GeoM.Rotate(3 * math.Pi / 2)
	rotateDest270.GeoM.Translate(float64(src.Bounds().Dx()/2), float64(src.Bounds().Dy()/2))
	rotateDest270.CompositeMode = ebiten.CompositeModeDestinationOver

	modes := []struct {
		name string
		op   *ebiten.DrawImageOptions
	}{
		{"Normal", &ebiten.DrawImageOptions{}},
		{"Rotate90", rotate90},
		{"Rotate180", rotate180},
		{"Rotate270", rotate270},
		{"FlipHorizontal", flipHorizontal},
		{"FlipVertical", flipVertical},
	}
	var table string

	modes2 := []struct {
		name string
		op   *ebiten.DrawImageOptions
	}{
		{"Normal", &ebiten.DrawImageOptions{}},
		{"RotateDest90", rotateDest90},
		{"RotateDest180", rotateDest180},
		{"RotateDest270", rotateDest270},
	}

	var j int
	for i, o := range modes2 {
		filename := fmt.Sprintf("examples/%s.png", o.name)
		tmp, _ := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
		if err := tmp.DrawImage(dst, o.op); err != nil {
			gfx.Log("couldn't draw image %s: %s", filename, err)
		}
		err := gfx.SavePNG(filename, tmp)
		if err != nil {
			return errors.Errorf("failed to save image:", err)
		}

		table += fmt.Sprintf("![example:%s](%s)<br>[%s](%s) |", o.name, filename, o.name, filename)
		if i%4 == 4-1 {
			table += "\n"
		}
		j++
	}
	for i, o := range modes {
		filename := fmt.Sprintf("examples/%s.png", o.name)
		tmp, _ := ebiten.NewImageFromImage(dst, ebiten.FilterDefault)
		if err := tmp.DrawImage(src, o.op); err != nil {
			gfx.Log("couldn't draw image %s: %s", filename, err)
		}
		err := gfx.SavePNG(filename, tmp)
		if err != nil {
			return errors.Errorf("failed to save image:", err)
		}

		table += fmt.Sprintf("![example:%s](%s)<br>[%s](%s) |", o.name, filename, o.name, filename)
		if (i+j)%4 == 4-1 {
			table += "\n"
		}
	}

	ioutil.WriteFile("EXAMPLES_2.md", []byte(header4+table), 0644)

	return nil
}

type Images struct {
	src image.Image
	dst image.Image
}

func (i *Images) update(screen *ebiten.Image) error {
	porterDuffComposition(i.src, i.dst)
	transformation(i.src, i.dst)
	return gfx.ErrDone
}

func main() {
	srcImg, err := gfx.OpenPNG("source.png")
	if err != nil {
		log.Fatal("could not open source.png:", err)
	}
	dstImg, err := gfx.OpenPNG("dest.png")
	if err != nil {
		log.Fatal("could not open dest.png:", err)
	}

	imgs := Images{
		src: srcImg,
		dst: dstImg,
	}

	if err := ebiten.Run(imgs.update, 10, 10, 1, "renderer"); err != nil {
		logrus.Error(err)
	}

}
