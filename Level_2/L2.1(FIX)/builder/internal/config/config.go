package config

import "builder_pattern/pkg/vector"

type Config struct {
	colorExterior vector.Vector3D
	colorBack     vector.Vector3D
	colorVertex   vector.Vector3D
	// настройка камеры
	typeCamera      string
	typeProjections string
	//настройка текстуры
	hasUvMap      bool
	textureLoaded bool
}

func (c Config) ColorExterior() vector.Vector3D {
	return c.colorExterior
}

func (c Config) ColorBack() vector.Vector3D {
	return c.colorBack
}

func (c Config) ColorVertex() vector.Vector3D {
	return c.colorVertex
}

func (c Config) TypeCamera() string {
	return c.typeCamera
}

func (c Config) TypeProjections() string {
	return c.typeProjections
}

func (c Config) HasUvMap() bool {
	return c.hasUvMap
}

func (c Config) TextureLoaded() bool {
	return c.textureLoaded
}

func NewConfig() Config {
	conf := Config{
		colorExterior:   vector.Vector3D{X: 255, Y: 255, Z: 255},
		colorBack:       vector.Vector3D{},
		colorVertex:     vector.Vector3D{X: 255},
		typeCamera:      "FOCUS",
		typeProjections: "PERSPECTIVE",
		hasUvMap:        false,
		textureLoaded:   false,
	}
	return conf
}
