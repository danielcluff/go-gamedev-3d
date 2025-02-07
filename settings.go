package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH, WINDOW_HEIGHT = 1920, 1080
const PLAYER_SPEED = 7
const LASER_SPEED = 9
const METEOR_TIMER_DURATION = 0.4
const FONT_SIZE = 60
const FONT_PADDING = 60

var BG_COLOR = r.Black
var METEOR_SPEED_RANGE = [2]int{3, 4}
