package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: <program> <args>\n Available args:\n - vec3\n - mat4\n - transformation")
	}

	switch os.Args[1] {
	case "vec3":
		vec3Examples()
	case "mat4":
		mat4Examples()
	case "transformation":
		TransformationMatExample()
	default:
		log.Fatal("Usage: <program> <args>\n Available args:\n - vec3\n - mat4\n - transformation")
	}
}
