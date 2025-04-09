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

	pauseButton *widget.Button

	msptInput *widget.TextInput
}

type handlers struct {
	pauseButton widget.ButtonClickedHandlerFunc
	msptSubmit  widget.TextInputChangedHandlerFunc
}

func newViewManager(size int, h handlers) *viewManager {
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
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		)),
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

	pauseButton := widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewBorderedNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255}, color.NRGBA{90, 90, 90, 255}, 3),
			Hover:   image.NewBorderedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, color.NRGBA{70, 70, 70, 255}, 3),
			Pressed: image.NewAdvancedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255})),
		}),
		widget.ButtonOpts.Text("Pause", mplusNormalFace, &widget.ButtonTextColor{
			Idle: color.Black,
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget.ButtonOpts.ClickedHandler(h.pauseButton),
	)

	msptInput := widget.NewTextInput(
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:      image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Highlight: image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 150, A: 255}),
		}),
		widget.TextInputOpts.Face(mplusNormalFace),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(mplusNormalFace, 2),
		),
		widget.TextInputOpts.Placeholder("0"),
		widget.TextInputOpts.SubmitHandler(h.msptSubmit),
		widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
			filteredText := ""
			for _, r := range args.InputText {
				if r >= '0' && r <= '9' {
					filteredText += string(r)
				}
			}

			if filteredText != args.InputText {
				args.TextInput.SetText(filteredText)
			}
		}),
	)

	right.AddChild(
		pauseButton,
		msptInput,
	)

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

		pauseButton: pauseButton,

		msptInput: msptInput,
	}
}
