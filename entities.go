package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	Model    r.Model
	Position r.Vector3
	Discard  bool
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

type Laser struct {
	*Entity
	Direction r.Vector3
	Speed     float32
}

func LaserCreate(model r.Model, pos r.Vector3) *Laser {
	laser := &Laser{
		Entity: &Entity{
			Model:    model,
			Position: pos,
		},
		Direction: r.Vector3{X: 0, Y: 0, Z: -1},
		Speed:     LASER_SPEED,
	}
	return laser
}
func (l *Laser) Move(dt float32) {
	l.Position.Z += l.Speed * l.Direction.Z * dt
}
func (l *Laser) Update(dt float32) {
	l.Move(dt)
}
