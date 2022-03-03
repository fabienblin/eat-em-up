package main

import (
	"flag"
	"os"
	// "regexp"
	"os/exec"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const MAX_PIECE_HEIGHT = 5
const MAX_PIECE_WIDTH = 5

var g1Flag = flag.String("g1", "", "First gang to be served")
var g2Flag = flag.String("g2", "", "Second gang to be served")
var mapFlag = flag.String("map", "", "A neighbourhood to take over")

var hood *grid = new(grid)

type grid struct {
	x int
	y int
	grid [][] byte
}

func main () {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	hood.loadMapFile(*mapFlag)
	var gangN byte = '1'
	var next string
	for true {
		piece := generatePiece(gangN)
		if gangN == '1' {
			g1 := exec.Command("./" + *g1Flag, piece.toString())
			out, err := g1.Output()
			check(err)
			playPosition(string(out), piece, gangN)
			gangN = '2'
		} else {
			g2 := exec.Command("./" + *g2Flag, piece.toString())
			out, err := g2.Output()
			check(err)
			playPosition(string(out), piece, gangN)
			gangN = '1'
		}
		fmt.Println(piece.toString())
		fmt.Println(hood.toString())

		fmt.Println("Press ENTER")
		fmt.Scanln(&next)
	}

}

func (g *grid) initGrid(y int, x int) grid {
	g.x = x
	g.y = y
	g.grid = make([][]byte, y)
	for i := 0; i < y; i++ {
		g.grid[i] = make([]byte, x)
	}
	return *g
}

func generatePiece(gangN byte) grid {
	var piece grid = new(grid).initGrid(rand.Intn(MAX_PIECE_HEIGHT) + 1, rand.Intn(MAX_PIECE_WIDTH) + 1)

	for i := 0; i < piece.y; i++ {
		for j := 0; j < piece.x; j++ {
			if (rand.Intn(2) == 0) {
				piece.grid[i][j] = '0'
			} else {
				piece.grid[i][j] = gangN
			}
		}
	}

	return piece
}

func (g *grid) loadMapFile(path string) grid {
	data, err := os.ReadFile(path)
    check(err)

	split := strings.Split(string(data), "\n")

	g.initGrid(len(split), len(split[0]))

	for i := 0; i < len(split); i++ {
		g.grid[i] = []byte(split[i])
	}

	return *g
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func isOutOfBoundries(posX int, posY int, piece grid) bool {
	return (	posY < 0 ||
				posX < 0 ||
				posY > hood.y - piece.y ||
				posX > hood.x - piece.x)
}

func isValidPosition(posX int, posY int, piece grid, gangN byte) bool {
	for i := 0; i < piece.y; i++ {
		for j := 0; j < piece.x; j++ {
			compare := piece.grid[i][j]
			if (compare == gangN){
				with := hood.grid[posY + i][posX + j]
				if (with != gangN && with != '0') {
					return false
				}
			}
		}
	}
	return true
}

func playPosition(pos string, piece grid, gangN byte) {
	posSplit := strings.Split(pos, ":")
	posY, _  := strconv.Atoi(posSplit[0])
	posX, _ := strconv.Atoi(posSplit[1])
	
	
	if (isOutOfBoundries(posX, posY, piece) || !isValidPosition(posX, posY, piece, gangN)) {
		return
	}

	for i := 0; i < piece.y; i++ {
		for j := 0; j < piece.x; j++ {
			if (piece.grid[i][j] == gangN) {
				hood.grid[posY + i][posX + j] = piece.grid[i][j]
			}
		}
	}
}

func (g *grid) toString() string {
	var s string = ""
	for _, row := range g.grid {
		s += string(row) + "\n"
	}
	return s
}