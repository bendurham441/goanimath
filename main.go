package main

import (
    "github.com/fogleman/gg"
    "fmt"
    "math"
)

const width, height int = 1920, 1080

func main() {
    base := Axes{
        xmin: -50,
        xmax: 50,
        ymin: -200,
        ymax: 500,
        xscale: 5,
        yscale: 50,
        GraphicalObject: GraphicalObject{
            width: 1920,
            height: 1080,
        },
    }
    dc := gg.NewContext(width, height)
    base.Draw(dc)
    base.PlotPoint(dc, Coord{1, 1})
    dc.Identity()
    base.LineGraphFunction(dc, 0.1, GraphFunction(func(n float64) (float64) {return 50 * math.Sin(n)}))
    dc.SavePNG("test.png")
}


type GraphicalObject struct {
    width float64
    height float64
    location Coord
}

type Axes struct {
    GraphicalObject
    xmin float64
    xmax float64
    ymin float64
    ymax float64
    xscale float64
    yscale float64
}

type Coord struct {
    x float64
    y float64
}

func (a *Axes) findOrigin() Coord {
    xrange := a.xmax - a.xmin
    yrange := a.ymax - a.ymin
    xcoord := a.width - a.xmax / xrange * a.width
    ycoord := a.ymax / yrange * a.height
    return Coord{xcoord, ycoord}
}

func (a *Axes) Draw(dc *gg.Context) error {
    orig := a.findOrigin()
    dc.SetRGB(1, 0.5, 0.5)
    dc.SetLineWidth(5)
    dc.DrawLine(orig.x, 0, orig.x, float64(height))
    dc.DrawLine(0, orig.y, float64(width), orig.y)
    for i := a.xscale; i < a.xmax; i = i+ a.xscale {
        tickWidth := a.xmax / (a.xmax - a.xmin) * float64(width) / (a.xmax / a.xscale)
        fmt.Println(tickWidth)
        tickX := orig.x + (i / a.xscale) * tickWidth
        dc.DrawLine(tickX, orig.y - 20, tickX, orig.y + 20)
    }
    for i := -a.xscale; i > a.xmin; i = i - a.xscale {
        tickWidth := a.xmin / (a.xmax - a.xmin) * float64(width) / (a.xmin / a.xscale)
        fmt.Println(tickWidth)
        tickX := orig.x + (i / a.xscale) * tickWidth
        dc.DrawLine(tickX, orig.y - 20, tickX, orig.y + 20)
    }
    for i := -a.yscale; i > -a.ymax; i = i - a.yscale {
        yRange := a.ymax - a.ymin
        tickWidth := a.ymax / (yRange) * float64(height) / (a.ymax / a.yscale)
        fmt.Println(tickWidth)
        tickY := orig.y + (i / a.yscale) * tickWidth
        dc.DrawLine(orig.x - 20, tickY, orig.x + 20, tickY)
    }
    for i := a.yscale; i < -a.ymin; i = i + a.yscale {
        yRange := a.ymax - a.ymin
        tickWidth := a.ymin / (yRange) * float64(height) / (a.ymin / a.yscale)
        fmt.Println(tickWidth)
        tickY := orig.y + (i / a.yscale) * tickWidth
        dc.DrawLine(orig.x - 20, tickY, orig.x + 20, tickY)
    }
    dc.Stroke()
    return nil
}

func (a *Axes) TranslatePoint(coord Coord) Coord {
    pointX := coord.x / (a.xmax - a.xmin) * float64(width)
    pointY := -coord.y / (a.ymax - a.ymin) * float64(height)
    return Coord{pointX, pointY}
}

func (a *Axes) LineGraphFunction(dc *gg.Context, inc float64, fn GraphFunction) error {
    for i := a.xmin; i < a.xmax - inc; i = i + inc {
        orig := a.findOrigin()
        dc.SetLineWidth(5)
        firstPoint := a.TranslatePoint(Coord{i, fn(i)})
        secondPoint := a.TranslatePoint(Coord{i + inc, fn(i + inc)})
        dc.DrawLine(orig.x + firstPoint.x, orig.y + firstPoint.y, orig.x + secondPoint.x, orig.y + secondPoint.y)
    }
    dc.Stroke()
    return nil
}

type GraphFunction func(float64) float64

func (a *Axes) PlotPoint(dc *gg.Context, coord Coord) (error) {
    orig := a.findOrigin()
    translatedPoint := a.TranslatePoint(coord)
    pointX := translatedPoint.x
    pointY := translatedPoint.y
    dc.DrawCircle(orig.x + pointX, orig.y - pointY, 10)
    dc.SetRGB(1, 1, 1)
    dc.Fill()
    return nil
}

