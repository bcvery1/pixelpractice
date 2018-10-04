package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/bcvery1/pixelpractice/movebox/common"
)

const (
	ip          = "0.0.0.0"
	port        = 9021
	portUpdates = 9022
)

var (
	clients atomic.Value
	src     = rand.NewSource(time.Now().UnixNano())
	r       = rand.New(src)
)

func genColour() color.RGBA {
	log.Println("Generating new colour")
	red := r.Int()
	green := r.Int()
	blue := r.Int()
	tmpColour := color.RGBA{uint8(red), uint8(green), uint8(blue), 255}

	for {
		tmpClients := clients.Load().(map[color.RGBA]*common.Player)
		if _, ok := tmpClients[tmpColour]; !ok {
			log.Println("New colour works")
			break
		}

		red = r.Int()
		green = r.Int()
		blue = r.Int()
		tmpColour = color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
		log.Printf("Generated temp colour: %v\n", tmpColour)
	}
	log.Printf("Generated colour %v\n", tmpColour)

	return tmpColour
}

func movementUpdates() {
	service := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.ListenPacket("udp", service)
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	for {
		msg := make([]byte, 258)
		_, _, err := listener.ReadFrom(msg)
		if err != nil && err != io.EOF {
			log.Printf("Failed to read bytes from packet: %v\n", err)
			continue
		}

		p := common.Player{}
		if err := p.DecodeFrom(msg); err != nil {
			log.Printf("Could not decode: %v\n", err)
			continue
		}

		tmpClients := clients.Load().(map[color.RGBA]*common.Player)
		if _, ok := tmpClients[p.Colour]; ok {
			tmpClients[p.Colour] = &p
			clients.Store(tmpClients)
		} else {
			log.Printf("Could not find colour %v, ignoring", p.Colour)
			continue
		}
	}
}

func allocateColours() {
	service := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Panicln(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v\n", err)
			continue
		}
		log.Printf("Got connection from %v\n", c.RemoteAddr())

		p := common.Player{
			Colour: genColour(),
		}

		log.Printf("Allocated %v\n", p.Colour)

		tmpClients := clients.Load().(map[color.RGBA]*common.Player)
		tmpClients[p.Colour] = &p
		clients.Store(tmpClients)

		pB, err := p.Byte()
		if err != nil {
			log.Printf("Could not create player bytes: %v\n", err)
			c.Close()
			continue
		}
		c.Write(pB)

		c.Close()
	}
}

func provideUpdatesFactory() {
	service := fmt.Sprintf("%s:%d", ip, portUpdates)
	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Panicln(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v\n", err)
			continue
		}

		log.Printf("Got connection from %v\n", c.RemoteAddr())

		p := common.Player{}
		pb := make([]byte, 256)
		_, err = c.Read(pb)
		if err != nil {
			log.Panicln(err)
		}
		p.DecodeFrom(pb)
		log.Printf("Received colour as %v\n", p.Colour)

		go provideUpdates(p, c)
	}
}

func provideUpdates(thisPlayer common.Player, conn net.Conn) {
	log.Printf("Providing updates to %v\n", conn.RemoteAddr())
	defer conn.Close()
	defer log.Printf("%v disconnected\n", conn.RemoteAddr())
	defer removePlayer(thisPlayer)

	tmpClients := clients.Load().(map[color.RGBA]*common.Player)
	if _, ok := tmpClients[thisPlayer.Colour]; !ok {
		log.Printf("Did not find %v in map, closing connection", thisPlayer.Colour)
		return
	}

	log.Println("Starting update loop")

	for {
		tmpClients := clients.Load().(map[color.RGBA]*common.Player)
		for _, p := range tmpClients {
			if p.Colour == thisPlayer.Colour {
				continue
			}

			pBytes, err := p.Byte()
			if err != nil {
				log.Printf("Could not decode player: %v\n", err)
				continue
			}

			written, err := conn.Write(pBytes)
			if err != nil {
				log.Printf("Failed to write to connection: %v", err)
				return
			}
			log.Printf("Wrote %d bytes\n", written)
		}

		time.Sleep(20 * time.Millisecond)
	}

	log.Println("No longer updating")
}

func removePlayer(p common.Player) {
	log.Printf("Removing player %v", p)
	tmpClients := clients.Load().(map[color.RGBA]*common.Player)
	if _, ok := tmpClients[p.Colour]; ok {
		delete(tmpClients, p.Colour)
		clients.Store(tmpClients)
	}
}

func main() {
	clients.Store(make(map[color.RGBA]*common.Player))
	go movementUpdates()
	go allocateColours()
	go provideUpdatesFactory()

	quitC := make(chan os.Signal, 1)
	signal.Notify(quitC, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT)
	<-quitC
}
