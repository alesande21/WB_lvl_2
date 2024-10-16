package objectSettings

import "builder_pattern/pkg/vector"

type ObjectSettings struct {
	// настройки цвета, линий и вершин (Class Color)
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
