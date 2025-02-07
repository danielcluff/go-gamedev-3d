package main

import r "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Model    r.Model
	Position r.Vector3
}
type Player struct {
	*Entity
}

func PlayerCreate(model r.Model, position r.Vector3) *Player {
	player := &Player{
		Entity: &Entity{
			Model:    model,
			Position: position,
		},
	}
	return player
}
func (p Player) Draw() {
	r.DrawModel(p.Model, p.Position, 1, r.White)
}
