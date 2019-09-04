package main

type screenManager struct {
	screens map[string]*screen
	currentScreen string
}

var _smInstance *screenManager

func smInstance() *screenManager {

	if _smInstance == nil {

		_smInstance = &screenManager{screens: make(map[string]*screen, 0)} // <-- not thread safe
	}

	return _smInstance
}

func (sm screenManager) addScreen(name string, screen screen) {
	sm.screens[name] = &screen
}

func (sm screenManager) getScreen(name string) screen {
	return *sm.screens[name]
}

func (sm *screenManager) switchScreen(name string) {
	// Unload old screen
	if len(sm.currentScreen) > 0 {
		currentScreen := *sm.screens[sm.currentScreen]
		currentScreen.unload()
	}

	sm.currentScreen = name

	// Load new screen
	newCurrentScreen := *sm.screens[sm.currentScreen]
	newCurrentScreen.load()

}

func (sm screenManager) tick() {
	screen := *sm.screens[sm.currentScreen]
	screen.tick()
}
