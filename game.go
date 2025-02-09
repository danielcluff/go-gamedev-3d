package main

import (
	"fmt"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

// have the player moving along a grid and create meteors ahead of the player

type Game struct {
	Camera        r.Camera3D
	Player        *Player
	Lasers        []*Laser
	Asteroids     []*Asteroid
	AsteroidTimer *Timer
	KillWall      [2]*Entity
	Assets
}
type Assets struct {
	Models
	Audio
	Textures
	Shaders
	Font r.Font
}
type Models struct {
	player r.Model
	laser  r.Model
}
type Textures struct {
	asteroids []r.Texture2D
	floor     r.Texture2D
	light     r.Texture2D
}
type Shaders struct {
	flash r.Shader
}
type Audio struct {
	music     r.Music
	laser     r.Sound
	explosion r.Sound
}

func (g *Game) Init() {
	r.InitWindow(WINDOW_WIDTH/3*2, WINDOW_HEIGHT/3*2, "3d space shooter")
	r.InitAudioDevice()
	g.ImportAssets()

	// Camera
	g.Camera.Position = r.Vector3{X: -4, Y: 8, Z: 6}
	g.Camera.Target = r.Vector3{X: 0, Y: 0, Z: -1}
	g.Camera.Up = r.Vector3{X: 0, Y: 1, Z: 0}
	g.Camera.Fovy = 45
	g.Camera.Projection = r.CameraPerspective
	g.KillWall[0] = CreateKillWall(-5)
	g.KillWall[1] = CreateKillWall(21)
	g.Player = PlayerCreate(g.player, r.Vector3{X: 0, Y: 0, Z: 0}, g.ShootLaser)
	g.AsteroidTimer = TimerCreate(0.4, true, true, g.CreateAsteroid)
	// r.PlayMusicStream(g.Audio.music)
}
func (g *Game) Update() {
	dt := r.GetFrameTime()
	g.Player.Update(dt)
	g.AsteroidTimer.Update()
	for i := range g.Asteroids {
		roid := g.Asteroids[i]
		roid.Update(dt)
	}
	for i := range g.Lasers {
		laser := g.Lasers[i]
		laser.Update(dt)
	}
	g.CheckCollisions()
	// r.UpdateMusicStream(g.Audio.music)
}
func (g *Game) Draw() {
	r.BeginDrawing()
	r.ClearBackground(BG_COLOR)
	r.BeginMode3D(g.Camera)

	g.Player.Draw()
	g.DrawAsteroids()
	g.DrawLasers()

	r.EndMode3D()
	r.EndDrawing()

}
func (g *Game) Run() {
	g.Init()
	defer r.CloseWindow()
	defer r.CloseAudioDevice()
	// defer r.UnloadAudioStream(g.Audio.music.Stream)

	for !r.WindowShouldClose() {

		g.Update()
		g.Draw()
	}
}
func (g *Game) ImportAssets() {
	g.Models.player = r.LoadModel(filepath.Join("assets", "models", "ship.glb"))
	g.Models.laser = r.LoadModel(filepath.Join("assets", "models", "laser.glb"))

	g.Shaders.flash = r.LoadShader("", filepath.Join("assets", "shaders", "flash.fs"))

	txtr := []string{"red", "green", "orange", "purple"}
	for i := range txtr {
		asset := r.LoadTexture(filepath.Join("assets", "textures", fmt.Sprintf("%v.png", txtr[i])))
		g.Assets.Textures.asteroids = append(g.Assets.Textures.asteroids, asset)
	}
	g.Assets.Textures.floor = r.LoadTexture(filepath.Join("assets", "textures", "dark.png"))
	g.Assets.Textures.light = r.LoadTexture(filepath.Join("assets", "textures", "light.png"))

	g.Assets.Audio.music = r.LoadMusicStream(filepath.Join("assets", "audio", "music.wav"))
	g.Assets.Audio.laser = r.LoadSound(filepath.Join("assets", "audio", "laser.wav"))
	g.Assets.Audio.explosion = r.LoadSound(filepath.Join("assets", "audio", "explosion.wav"))

	g.Assets.Font = r.LoadFontEx(filepath.Join("assets", "font", "Stormfaze.otf"), FONT_SIZE, nil, 0)

}
func (g *Game) DiscardEntities() {

}
func (g *Game) CheckCollisions() {
	for a := range g.Asteroids {
		roid := *g.Asteroids[a]
		// check collision if true
		roid.Discard = true
	}
}
func (g *Game) DrawLasers() {
	for _, laser := range g.Lasers {
		laser.Draw()
	}
}
func (g *Game) DrawAsteroids() {
	for _, roid := range g.Asteroids {
		roid.Draw()
	}
}
func (g *Game) CreateAsteroid() {
	roid := AsteroidCreate(g.asteroids)
	g.Asteroids = append(g.Asteroids, roid)
}
func (g *Game) ShootLaser(pos r.Vector3) {
	laser := LaserCreate(g.Models.laser, pos)
	g.Lasers = append(g.Lasers, laser)
	r.PlaySound(g.Audio.laser)
}
