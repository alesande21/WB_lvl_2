package view

type Widget interface {
	Open() string
}

type TempWindowFactory struct {
	colorWindow   *ColorSettings
	textureWindow *TextureSettings
	cameraWindow  *CameraSettings
}

func NewTempWindowFactory() *TempWindowFactory {
	return &TempWindowFactory{}
}

func (t *TempWindowFactory) GetWindow(windowName string) Widget {
	switch windowName {
	case "color":
		if t.colorWindow == nil {
			t.colorWindow = NewColorSettings()
		}
		return t.colorWindow
	case "camera":
		if t.cameraWindow == nil {
			t.cameraWindow = NewCameraSettings()
		}
		return t.cameraWindow
	case "texture":
		if t.textureWindow == nil {
			t.textureWindow = NewTextureSettings()
		}
		return t.textureWindow
	}
	return t.colorWindow
}
