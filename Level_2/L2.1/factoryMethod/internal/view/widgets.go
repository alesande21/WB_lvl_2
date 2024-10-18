package view

type CameraSettings struct {
	name string
}

func (c CameraSettings) Open() string {
	return c.name
}

func NewCameraSettings() *CameraSettings {
	str := "Отркыты настройки камеры."
	return &CameraSettings{name: str}
}

type ColorSettings struct {
	name string
}

func (c ColorSettings) Open() string {
	return c.name
}

func NewColorSettings() *ColorSettings {
	str := "Отркыты настройки цвета."
	return &ColorSettings{name: str}
}

type TextureSettings struct {
	name string
}

func (c TextureSettings) Open() string {
	return c.name
}

func NewTextureSettings() *TextureSettings {
	str := "Отркыты настройки текстуры."
	return &TextureSettings{name: str}
}
