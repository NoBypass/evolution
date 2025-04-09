package game

import (
	"bytes"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	image2 "image"
	"image/color"
	"image/png"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
}

type viewManager struct {
	ui *ebitenui.UI

	simulationImg *ebiten.Image

	generationLabel *widget.Text
	survivorsLabel  *widget.Text
}

func newViewManager(size int) *viewManager {
	img := image2.NewRGBA(image2.Rect(0, 0, size, size))
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	simulationImg, _, err := ebitenutil.NewImageFromReader(&buf)
	if err != nil {
		panic(err)
	}

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0xff, 0xff, 0xff, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(5)),
			widget.GridLayoutOpts.DefaultStretch(false, true),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true, true}),
		)),
	)

	left := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		)),
	)

	right := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 0, 128})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
			}),
		),
	)

	simulationGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(simulationImg),
		widget.GraphicOpts.WidgetOpts(),
	)

	header := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(25),
			),
		),
	)

	generationLabel := widget.NewText(widget.TextOpts.Text("", mplusNormalFace, color.Black))
	survivorsLabel := widget.NewText(widget.TextOpts.Text("", mplusNormalFace, color.Black))

	header.AddChild(
		generationLabel,
		survivorsLabel,
	)

	left.AddChild(
		header,
		simulationGraphic,
	)

	right.AddChild()

	root.AddChild(
		left,
		right,
	)

	return &viewManager{
		ui: &ebitenui.UI{
			Container: root,
		},

		simulationImg: simulationImg,

		generationLabel: generationLabel,
		survivorsLabel:  survivorsLabel,
	}
}
