package controller

import (
	"builder_pattern/internal/config"
	"builder_pattern/internal/objectSettings"
)

type Controller struct {
	builder objectSettings.Builder
}

func NewController() *Controller {
	bldr := objectSettings.NewObjectBuilder()
	return &Controller{builder: bldr}
}

func (c *Controller) LoadObjectSettings(config *config.Config) {
	c.builder.SetColorBack(config.ColorBack())
	c.builder.SetColorExterior(config.ColorExterior())
	c.builder.SetColorVertex(config.ColorVertex())
	c.builder.SetHasUvMap(config.HasUvMap())
	c.builder.SetTextureLoaded(config.TextureLoaded())
}

func (c *Controller) GetObjectSettings() *objectSettings.ObjectSettings {
	return c.builder.GetObject()
}
