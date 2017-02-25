package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/closedgl"
)

func main() {
	closedgl.Run(piirrä, 640, 640, "Liikkuvaa kuvaa")
}

func piirrä(aika float64) {
	gl.Begin(gl.LINES)

	gl.Color3d(1, 1, 1)
	gl.Vertex2d(-1, 0)
	gl.Vertex2d(aika-1, 0)

	gl.End()
}
