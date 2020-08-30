package ttme

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilePropertiesWidget struct {
	x, y int
	properties []tileProperty
	selectedProperty int
	editMode bool
}

func NewTilePropertiesWidget(x, y int) *tilePropertiesWidget {
	newWidget := tilePropertiesWidget{x: x, y: y}
	newWidget.properties = append(newWidget.properties, tileProperty{name: "none"})
	newWidget.properties = append(newWidget.properties, tileProperty{name: "blocking", value: "true", color: r.Green})
	newWidget.selectedProperty = 0
	newWidget.editMode = false

	return &newWidget
}

func (tpw *tilePropertiesWidget) Draw() {
	editMode := tpw.editMode
	selectBounds := r.Rectangle{X: float32(tpw.x), Y: float32(tpw.y), Width: 200, Height: 20}
	editMode, tpw.selectedProperty = r.GuiDropdownBox(selectBounds, tpw.AsText(), tpw.selectedProperty, tpw.editMode)

	if editMode {
		tpw.editMode = !tpw.editMode
	}
}

func (tpw *tilePropertiesWidget) Unset() {
	tpw.selectedProperty = 0
}

func (tpw tilePropertiesWidget) SelectedProperty() tileProperty {
	return tpw.properties[tpw.selectedProperty]
}

func (tpw tilePropertiesWidget) Selected() bool {
	return tpw.selectedProperty == 0
}

func (tpw tilePropertiesWidget) AsText() string {
	text := ""

	for i := 0; i < len(tpw.properties); i++ {
		baseText := ""

		if i > 0 {
			baseText += ";"
		}

		text += baseText + tpw.properties[i].name
	}

	return text
}
