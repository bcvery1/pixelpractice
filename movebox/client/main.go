package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net"
	"sync/atomic"
	"time"

	"github.com/bcvery1/pixelpractice/movebox/common"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	winWidth    = 1280
	winHeight   = 720
	initX       = 128
	initY       = 72
	ip          = "10.0.12.6"
	port        = 9021
	portUpdates = 9022
)

var (
	backgroundColour = colornames.Whitesmoke
	player           = common.Player{}
	playerSpeed      = 150.0
	otherPlayers     atomic.Value
)

func updateOtherPlayers() {
	log.Println("Starting get-updates loop")

	dest := fmt.Sprintf("%s:%d", ip, portUpdates)
	c, err := net.DialTimeout("tcp", dest, 5*time.Second)
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()
	log.Println("Established server connection")

	pBytes, err := player.Byte()
	if err != nil {
		log.Panicln(err)
	}
	_, err = c.Write(pBytes)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Sent this player to server for identification")

	for {
		pB := make([]byte, 256)
		_, err := c.Read(pB)
		if err != nil {
			fmt.Printf("Failed to read from connection: %v\n", err)
			continue
		}

		p := common.Player{}
		p.DecodeFrom(pB)
		tmpPlayers := otherPlayers.Load().(map[color.RGBA]*common.Player)
		tmpPlayers[p.Colour] = &p
		otherPlayers.Store(tmpPlayers)
	}
}

func register() {
	log.Println("Registering with server")

	dest := fmt.Sprintf("%s:%d", ip, port)
	c, err := net.DialTimeout("tcp", dest, 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	log.Println("Got connection")

	pB, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}

	player.DecodeFrom(pB)
	player.X = initX
	player.Y = initY

	log.Printf("Assigned colour: %v\n", player.Colour)
}

func sendPos() {
	c, err := net.Dial("udp", "127.0.0.1:9021")
	defer c.Close()
	if err != nil {
		log.Panic(err)
	}

	msgB, err := player.Byte()
	if err != nil {
		log.Panic(err)
	}
	c.Write(msgB)
}

func drawPlayers(t pixel.Target) {
	tmpPlayers := otherPlayers.Load().(map[color.RGBA]*common.Player)
	for _, p := range tmpPlayers {
		p.Draw(t)
	}
}

func run() {
	// Register
	register()

	otherPlayers.Store(make(map[color.RGBA]*common.Player))
	go updateOtherPlayers()

	cfg := pixelgl.WindowConfig{
		Title:       "MoveBoxes",
		Bounds:      pixel.R(0, 0, winWidth, winHeight),
		VSync:       true,
		Undecorated: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(backgroundColour)

		player.Draw(win)
		drawPlayers(win)

		if win.Pressed(pixelgl.KeyA) {
			player.X -= playerSpeed * dt
			if player.X < 0.0 {
				player.X = 0.0
			}
		}
		// Try move left
		if win.Pressed(pixelgl.KeyD) {
			player.X += playerSpeed * dt
			if player.X > win.Bounds().Max.X-common.PlayerWidth {
				player.X = win.Bounds().Max.X - common.PlayerWidth
			}
		}
		// Try move down
		if win.Pressed(pixelgl.KeyS) {
			player.Y -= playerSpeed * dt
			if player.Y < 0.0 {
				player.Y = 0.0
			}
		}
		// Try move up
		if win.Pressed(pixelgl.KeyW) {
			player.Y += playerSpeed * dt
			if player.Y > win.Bounds().Max.Y-common.PlayerHeight {
				player.Y = win.Bounds().Max.Y - common.PlayerHeight
			}
		}

		win.Update()
		sendPos()
	}
}

func main() {
	pixelgl.Run(run)
}
