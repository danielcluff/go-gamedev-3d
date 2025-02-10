package main

import (
	"math"
	"math/rand"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	Radius        float32
	DeathTimer    *Timer
	Hit           bool
	RotationAxis  r.Vector3
	RotationSpeed r.Vector3
	*Entity
}

// add rotation
func AsteroidCreate(textures []r.Texture2D) *Asteroid {
	radius := float32(math.Max(math.Min(1.2, rand.Float64()*1.2), 0.8))
	model := r.LoadModelFromMesh(r.GenMeshSphere(radius, 8, 8))
	shader := r.LoadShader("", filepath.Join("assets", "shaders", "flash.fs"))
	model.Materials.Shader = shader
	randomTextureInded := rand.Intn(len(textures))
	texture := textures[randomTextureInded]
	r.SetMaterialTexture(&model.GetMaterials()[0], r.MapAlbedo, texture)

	// generate random position
	position := r.Vector3{}
	position.X = (rand.Float32()*2 - 1) * 10
	position.Y = (rand.Float32()*2 - 1) / 1.25
	position.Z = -20

	// generate rotation
	rotation := r.Vector3{}
	rotation.X = rand.Float32()
	rotation.Y = rand.Float32()
	rotation.Z = rand.Float32()

	asteroid := &Asteroid{
		Radius:        radius,
		RotationAxis:  rotation,
		RotationSpeed: r.Vector3{X: rand.Float32()*2 - 1, Y: rand.Float32()*2 - 1, Z: rand.Float32()*2 - 1},
		Entity: &Entity{
			Direction: r.Vector3{X: 0, Y: 0, Z: 1},
			Speed:     rand.Float32() + 3,
			Model:     model,
			Position:  position,
		},
	}
	asteroid.DeathTimer = TimerCreate(0.25, false, false, asteroid.Remove)
	return asteroid
}
func (a *Asteroid) Flash() {
	a.Direction.Z = 0
	flashLoc := r.GetShaderLocation(a.Model.Materials.Shader, "flash")
	r.SetShaderValue(a.Model.Materials.Shader, flashLoc, []float32{1, 0}, r.ShaderUniformVec2)
}
func (a *Asteroid) Rotate(dt float32) {
	a.RotationAxis.X += 200 * dt
	a.RotationAxis.Y += 20 * dt
	a.RotationAxis.Z += 20 * dt
}
func (a *Asteroid) Update(dt float32) {
	a.DeathTimer.Update()
	if !a.Hit {
		a.Entity.Update(dt)
		a.RotationAxis.X += a.RotationSpeed.X * dt
		a.RotationAxis.Y += a.RotationSpeed.Y * dt
		a.RotationAxis.Z += a.RotationSpeed.Z * dt
		a.Model.Transform = r.MatrixRotateXYZ(a.RotationAxis)
	}
}
func (a *Asteroid) Remove() {
	a.Discard = true
}
