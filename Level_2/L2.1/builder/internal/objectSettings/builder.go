package objectSettings

import (
	"builder_pattern/pkg/vector"
)

type Builder interface {
	SetColorExterior(config vector.Vector3D)
	SetColorBack(config vector.Vector3D)
	SetColorVertex(config vector.Vector3D)
	SetTypeCamera(cameraType string)
	SetTypeProjections(projectionType string)
	SetHasUvMap(status bool)
	SetTextureLoaded(status bool)
	GetObject() *ObjectSettings
}
