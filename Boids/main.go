package main

import (
	"os"
	"fmt"
	"strconv"
	"gifhelper"
)

func main() {
	//Place your code here.


  //take user input

  numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil{
		panic(err1)
	}
	if numBoids<0 {
		panic("Negative num of boids given")
	}


  skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil{
		panic(err2)
	}

  initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
  if err3 != nil{
  	panic(err3)
  }

  maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
  if err4 != nil{
    panic(err4)
  }

  numGens, err5 := strconv.Atoi(os.Args[5])
  if err5 != nil{
    panic(err5)
  }
  if numBoids<0 {
    panic("Negative num of gens given")
  }

  proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
  if err6 != nil{
    panic(err6)
  }

  separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
  if err7 != nil{
    panic(err7)
  }

  alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
  if err8 != nil{
    panic(err8)
  }

  cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
  if err9 != nil{
    panic(err9)
  }

  time, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10 != nil{
		panic(err10)
	}

	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil{
		panic(err11)
	}

	drawingFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12 != nil{
		panic(err12)
	}

  fmt.Println("Input taken!")

  initialSky := InitializeSky(skyWidth, initialSpeed, maxBoidSpeed, proximity, separationFactor, alignmentFactor, cohesionFactor, numBoids)

  fmt.Println("Initial sky is set!")

  timePoints := SimulateBoids(initialSky, numGens, time)

  fmt.Println("Simulation done!")

  images := AnimateSystem(timePoints, canvasWidth, drawingFrequency)

	fmt.Println("Images drawn")


	fmt.Println("Draw gif")

	gifhelper.ImagesToGIF(images, "boids")

	fmt.Println("Animated gif produced")
	fmt.Println("Done")


}
