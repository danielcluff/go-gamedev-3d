package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	Model     r.Model
	Position  r.Vector3
	Direction r.Vector3
	Speed     float32
	Discard   bool
}

func (e *Entity) GetBoundingBox() r.BoundingBox {
	boundingBox := r.GetMeshBoundingBox(e.Model.GetMeshes()[0])
	minBoundary := r.Vector3Add(e.Position, boundingBox.Min)
	maxBoundary := r.Vector3Add(e.Position, boundingBox.Max)
	return r.BoundingBox{Min: minBoundary, Max: maxBoundary}
}
func (e *Entity) Move(dt float32) {
	e.Position.X += e.Speed * e.Direction.X * dt
	e.Position.Y += e.Speed * e.Direction.Y * dt
	e.Position.Z += e.Speed * e.Direction.Z * dt
}
func (e *Entity) Update(dt float32) {
	e.Move(dt)
}
func (e Entity) Draw() {
	r.DrawModel(e.Model, e.Position, 1, r.White)
}

func CreateKillWall(z float32) *Entity {
	model := r.LoadModelFromMesh(r.GenMeshCube(10, 2, 1))
	wall := &Entity{
		Model:    model,
		Position: r.Vector3{X: 0, Y: 0, Z: z},
	}
	return wall
}

func LaserCreate(model r.Model, pos r.Vector3) *Entity {
	laser := &Entity{
		Model:     model,
		Position:  pos,
		Direction: r.Vector3{X: 0, Y: 0, Z: -1},
		Speed:     LASER_SPEED,
	}
	return laser
}
func FloorCreate(texture r.Texture2D) *Entity {
	model := r.LoadModelFromMesh(r.GenMeshCube(32, 1, 32))
	r.SetMaterialTexture(model.Materials, r.MapAlbedo, texture)
	floor := &Entity{
		Model:    model,
		Position: r.Vector3{X: 6, Y: -2, Z: -8},
	}
	return floor
}
