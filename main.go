package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"time"
// )

// func clearScreen() {
// 	switch runtime.GOOS {
// 	case "linux", "darwin": // Unix-like systems
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		_ = cmd.Run()
// 	case "windows":
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		_ = cmd.Run()
// 	default:
// 		// Unsupported platform; clear the screen by printing newlines
// 		fmt.Print("\033[H\033[2J")
// 	}
// }

// type Graphics2D struct {
// 	Width  int
// 	Height int
// 	Arr    [][]string
// 	FPS    int
// }

// func (g *Graphics2D) NewGraphics2D(w, h, fps int) {
// 	g.Width = w
// 	g.Height = h
// 	g.FPS = fps
// 	g.Arr = make([][]string, g.Height)
// 	g.NewArray()
// }

// func (g *Graphics2D) NewArray() {
// 	for i := range g.Arr {
// 		g.Arr[i] = make([]string, g.Width)
// 	}
// 	for i := range g.Arr {
// 		for j := range g.Arr[i] {
// 			g.Arr[i][j] = "."
// 		}
// 	}
// }

// func (g *Graphics2D) Calculate() {
// 	arr := g.Arr
// 	if arr[g.Height-1][g.Width-1] != "." {
// 		g.NewArray()
// 	}
// 	hMax := g.Height - 1
// 	hMin := 0
// 	wMax := g.Width - 1
// 	wMin := 0
// 	for i := 0; i < g.Height/5; i++ {
// 		h := rand.Intn(hMax-hMin+1) + hMin
// 		w := rand.Intn(wMax-wMin+1) + wMin
// 		luck := rand.Intn(3-1+1) + 1
// 		if luck == 1 {
// 			arr[h][w] = "*"
// 		}else if luck == 2 {
// 			arr[h][w] = "@"
// 		}else if luck == 3 {
// 			arr[h][w] = "?"
// 		}

// 	}
// 	// for i := 0; i < g.Height; i++ {

// 	// 	if arr[i][0] == "." {
// 	// 		for j := 0; j < g.Width; j++ {
// 	// 			arr[i][j] = "*"

// 	// 		}
// 	// 		break
// 	// 	}

// 	// }

// }
// func (g *Graphics2D) Print() {
// 	sleepTime := time.Duration(1000/g.FPS) * time.Millisecond
// 	for {
// 		for _, v := range g.Arr {
// 			for _, v2 := range v {
// 				fmt.Print(v2)
// 			}
// 			fmt.Println()
// 		}
// 		time.Sleep(sleepTime)
// 		clearScreen()

// 	}
// }

// func (g *Graphics2D) Start() {
// 	go func() {
// 		for {
// 			g.Calculate()
// 			time.Sleep(1 * time.Millisecond)
// 		}
// 	}()

// 	go func() {

// 		g.Print()

// 	}()

// }
// func main() {
// 	var g Graphics2D
// 	g.NewGraphics2D(100, 50, 10)
// 	g.Start()
// 	select {}
// }
