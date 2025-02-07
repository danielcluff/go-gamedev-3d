package main

import (
	"fmt"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

// have the player moving along a grid and create meteors ahead of the player

type Game struct {
	Camera r.Camera3D
	Player Player
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
	r.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "3d space shooter")
	r.InitAudioDevice()
	g.ImportAssets()

	// Camera
	g.Camera.Position = r.Vector3{X: -4, Y: 8, Z: 6}
	g.Camera.Target = r.Vector3{X: 0, Y: 0, Z: -1}
	g.Camera.Up = r.Vector3{X: 0, Y: 1, Z: 0}
	g.Camera.Fovy = 45
	g.Camera.Projection = r.CameraPerspective

	g.Player = *PlayerCreate(g.player, r.Vector3{X: 0, Y: 0, Z: 0})
	// r.PlayMusicStream(g.Audio.music)
}
func (g *Game) Update() {
	// dt = r.GetFrameTime()
	// r.UpdateMusicStream(g.Audio.music)
}
func (g *Game) Draw() {
	r.BeginDrawing()
	r.ClearBackground(BG_COLOR)
	r.BeginMode3D(g.Camera)

	g.Player.Draw()

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

	g.Assets.Audio.music = r.LoadMusicStream(filepath.Join("assets", "models", "laser.wav"))
	g.Assets.Audio.laser = r.LoadSound(filepath.Join("assets", "models", "laser.wav"))
	g.Assets.Audio.explosion = r.LoadSound(filepath.Join("assets", "models", "explosion.wav"))

	g.Assets.Font = r.LoadFontEx(filepath.Join("assets", "font", "Stormfaze.otf"), FONT_SIZE, nil, 0)

}
func (g *Game) DiscardEntities() {

}
