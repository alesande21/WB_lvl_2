package objectSettings

import (
	"builder_pattern/pkg/vector"
)

type ObjectBuilder struct {
	settings ObjectSettings
}

func (o *ObjectBuilder) SetColorExterior(config vector.Vector3D) {
	o.settings.colorExterior = config
}

func (o *ObjectBuilder) SetColorBack(config vector.Vector3D) {
	o.settings.colorBack = config
}

func (o *ObjectBuilder) SetColorVertex(config vector.Vector3D) {
	o.settings.colorVertex = config
}

func (o *ObjectBuilder) SetTypeCamera(cameraType string) {
	o.settings.typeCamera = cameraType
}

func (o *ObjectBuilder) SetTypeProjections(projectionType string) {
	o.settings.typeProjections = projectionType
}

func (o *ObjectBuilder) SetHasUvMap(status bool) {
	o.settings.hasUvMap = status
}

func (o *ObjectBuilder) SetTextureLoaded(status bool) {
	o.settings.textureLoaded = status
}

func (o *ObjectBuilder) GetObject() *ObjectSettings {
	return &o.settings
}

func NewObjectBuilder() *ObjectBuilder {
	return &ObjectBuilder{settings: ObjectSettings{
		colorExterior:   vector.Vector3D{},
		colorBack:       vector.Vector3D{},
		colorVertex:     vector.Vector3D{},
		typeCamera:      "",
		typeProjections: "",
		hasUvMap:        true,
		textureLoaded:   true,
	}}
}
