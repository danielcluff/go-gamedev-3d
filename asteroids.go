package main

import (
	"math"
	"math/rand"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	Direction r.Vector3
	Speed     float32
	*Entity
}

// add rotation
func AsteroidCreate(textures []r.Texture2D) *Asteroid {
	radius := float32(math.Max(math.Min(1.2, rand.Float64()*1.2), 0.8))
	model := r.LoadModelFromMesh(r.GenMeshSphere(radius, 8, 8))
	randomTextureInded := rand.Intn(len(textures))
	texture := textures[randomTextureInded]
	r.SetMaterialTexture(&model.GetMaterials()[0], r.MapAlbedo, texture)

	// generate random position
	position := r.Vector3{}
	position.X = (rand.Float32()*2 - 1) * 10
	position.Y = (rand.Float32()*2 - 1) / 1.25
	position.Z = -20

	asteroid := &Asteroid{
		Direction: r.Vector3{X: 0, Y: 0, Z: 1},
		Speed:     rand.Float32() + 3,
		Entity: &Entity{
			Model:    model,
			Position: position,
		},
	}
	return asteroid
}
func (p *Asteroid) Move(dt float32) {
	p.Position.Z += p.Speed * p.Direction.Z * dt
}
func (p *Asteroid) Update(dt float32) {
	p.Move(dt)
	// transform rotation
}
func (p Asteroid) Draw() {
	r.DrawModel(p.Model, p.Position, 1, r.White)
}
