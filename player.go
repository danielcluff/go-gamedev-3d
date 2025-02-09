package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Direction  r.Vector3
	Speed      float32
	ShootLaser func(pos r.Vector3)
	*Entity
}

// need to create a bounding box for the model and move it with the player
// need to create a flat plane for the floor texture
// check for collisions in the killboxes and remove entities
func PlayerCreate(model r.Model, position r.Vector3, shootLaser func(pos r.Vector3)) *Player {
	player := &Player{
		Direction:  r.Vector3{X: 0, Y: 0, Z: 0},
		Speed:      12,
		ShootLaser: shootLaser,
		Entity: &Entity{
			Model:    model,
			Position: position,
		},
	}
	return player
}
func (p *Player) Input() {
	p.Direction.X = float32(BoolToInt(r.IsKeyDown(r.KeyRight))) - float32(BoolToInt(r.IsKeyDown(r.KeyLeft)))

	if r.IsKeyPressed(r.KeySpace) {
		p.ShootLaser(p.Position)
	}
}
func (p *Player) Move(dt float32) {
	p.Position.X += p.Speed * p.Direction.X * dt
}
func (p *Player) Constraint() {
	if p.Position.X > 10 {
		p.Position.X = 10
	} else if p.Position.X < -10 {
		p.Position.X = -10
	}
}
func (p *Player) Update(dt float32) {
	p.Input()
	p.Move(dt)
	p.Constraint()
}
