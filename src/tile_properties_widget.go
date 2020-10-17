package ttme

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilePropertiesWidget struct {
	x, y int
	properties []tileProperty
	selectedProperty int
	editMode, propertyNameEditMode, propertyValueEditMode bool
	propertyName, propertyValue string
	propertyColor r.Color
}

func NewTilePropertiesWidget(x, y int) *tilePropertiesWidget {
	newWidget := tilePropertiesWidget{x: x, y: y}
	newWidget.properties = append(newWidget.properties, tileProperty{Name: "none"})
	newWidget.properties = append(newWidget.properties, tileProperty{Name: "blocking", Value: "true", Color: r.Green})
	newWidget.Reset()

	return &newWidget
}

func (tpw *tilePropertiesWidget) Draw() {
	editMode := tpw.editMode
	propertyNameEditMode := tpw.propertyNameEditMode
	propertyValueEditMode := tpw.propertyValueEditMode

	r.GuiLabel(r.Rectangle{X: float32(tpw.x), Y: float32(tpw.y + 40), Width: 40, Height: 20}, "Property name:")
	propertyNameEditMode, tpw.propertyName = r.GuiTextBox(r.Rectangle{X: float32(tpw.x + 90), Y: float32(tpw.y + 40), Width: 80, Height: 20}, tpw.propertyName, 50, tpw.propertyNameEditMode)

	r.GuiLabel(r.Rectangle{X: float32(tpw.x), Y: float32(tpw.y + 70), Width: 40, Height: 20}, "Property value:")
	propertyValueEditMode, tpw.propertyValue = r.GuiTextBox(r.Rectangle{X: float32(tpw.x + 90), Y: float32(tpw.y + 70), Width: 80, Height: 20}, tpw.propertyValue, 50, tpw.propertyValueEditMode)

	tpw.propertyColor = r.GuiColorPicker(r.Rectangle{X: float32(tpw.x), Y: float32(tpw.y + 100), Width: 100, Height: 100}, tpw.propertyColor)

	if (r.GuiButton(r.Rectangle{X: float32(tpw.x + 150), Y: float32(tpw.y) + 175, Width: 95, Height: 25}, "Create property")) {
		tpw.properties = append(tpw.properties, tileProperty{Name: tpw.propertyName, Value: tpw.propertyValue, Color: tpw.propertyColor})
		tpw.Reset()
	}

	selectBounds := r.Rectangle{X: float32(tpw.x), Y: float32(tpw.y), Width: 200, Height: 20}
	editMode, tpw.selectedProperty = r.GuiDropdownBox(selectBounds, tpw.AsText(), tpw.selectedProperty, tpw.editMode)

	if editMode {
		tpw.editMode = !tpw.editMode
	}

	if propertyNameEditMode {
		tpw.propertyNameEditMode = !tpw.propertyNameEditMode
	}

	if propertyValueEditMode {
		tpw.propertyValueEditMode = !tpw.propertyValueEditMode
	}
}

func (tpw *tilePropertiesWidget) Reset() {
	tpw.selectedProperty = 0
	tpw.editMode = false
	tpw.propertyNameEditMode = false
	tpw.propertyValueEditMode = false
	tpw.propertyName = ""
	tpw.propertyValue = ""
	tpw.propertyColor = r.Green
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

		text += baseText + tpw.properties[i].Name
	}

	return text
}
