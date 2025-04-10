package ui

import (
	"evolution/internal/util"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type ViewManager struct {
	UI *ebitenui.UI

	SimulationImg *ebiten.Image
	GraphImg      *ebiten.Image

	GenerationLabel *widget.Text
	SurvivorsLabel  *widget.Text

	PauseButton *widget.Button

	MsptInput *widget.TextInput
}

type Handlers struct {
	PauseButton     widget.ButtonClickedHandlerFunc
	MsptSubmit      widget.TextInputChangedHandlerFunc
	SimulationClick widget.WidgetMouseButtonReleasedHandlerFunc
}

func NewViewManager(size int, h Handlers) *ViewManager {
	viewManager := new(ViewManager)

	viewManager.SimulationImg = ebiten.NewImage(size, size)
	viewManager.GraphImg = ebiten.NewImage(size, size)

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0xff, 0xff, 0xff, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(5)),
			widget.GridLayoutOpts.DefaultStretch(false, true),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true, true}),
		)),
	)

	viewManager.UI = &ebitenui.UI{
		Container: root,
	}

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
		widget.GraphicOpts.Image(viewManager.SimulationImg),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.MouseButtonReleasedHandler(h.SimulationClick),
		),
	)

	header := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(25),
			),
		),
	)

	viewManager.GenerationLabel = widget.NewText(widget.TextOpts.Text("", util.MplusNormalFace, color.Black))
	viewManager.SurvivorsLabel = widget.NewText(widget.TextOpts.Text("", util.MplusNormalFace, color.Black))

	header.AddChild(
		viewManager.GenerationLabel,
		viewManager.SurvivorsLabel,
	)

	left.AddChild(
		header,
		simulationGraphic,
	)

	viewManager.PauseButton = widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewBorderedNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255}, color.NRGBA{90, 90, 90, 255}, 3),
			Hover:   image.NewBorderedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, color.NRGBA{70, 70, 70, 255}, 3),
			Pressed: image.NewAdvancedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255})),
		}),
		widget.ButtonOpts.Text("Pause", util.MplusNormalFace, &widget.ButtonTextColor{
			Idle: color.Black,
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget.ButtonOpts.ClickedHandler(h.PauseButton),
	)

	viewManager.MsptInput = widget.NewTextInput(
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:      image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Highlight: image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 150, A: 255}),
		}),
		widget.TextInputOpts.Face(util.MplusNormalFace),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(util.MplusNormalFace, 2),
		),
		widget.TextInputOpts.Placeholder("0"),
		widget.TextInputOpts.SubmitHandler(h.MsptSubmit),
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

	graphGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(viewManager.GraphImg),
	)

	right.AddChild(
		viewManager.PauseButton,
		viewManager.MsptInput,
		graphGraphic,
	)

	root.AddChild(
		left,
		right,
	)

	return viewManager
}
