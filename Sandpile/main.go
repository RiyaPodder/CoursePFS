package main

import(
  "fmt"
  "strconv"
  "os"
  "log"
  "time"
)

func main(){

  fmt.Println("Sandpile Simulation!")

  //take user inputs of size, pile and placement
  size, err1 := strconv.Atoi(os.Args[1])
  if err1 != nil {
    panic("Problem converting size parameter to an integer!")
  }
  pile, err2 := strconv.Atoi(os.Args[2])
  if err2 != nil {
    panic("Problem converting pile parameter to an integer!")
  }
  placement := os.Args[3]
  fmt.Println("User input taken successfully!")

  //Initialize first board
  initial_gameboard := InitializeGameboard(size, pile, placement)

  //Call serial function
  fmt.Println("Start sandpile simulation in serial!")
  start := time.Now()
  configurations := SimulateSandpile(initial_gameboard)
  elapsed := time.Since(start)
  log.Printf("Time taken %s", elapsed)
  fmt.Println("Simulation run. Now drawing final stable configuration.")
  //in DrawPNGFinal, the second parameter 1 is the width of each cell in the board which is 1 pixel
  DrawPNGFinal(configurations, 1, "sandpile_serial.png")
  fmt.Println("Image drawn!")

  //Call parallel function
  fmt.Println("Start sandpile simulation in parallel!")
  start2 := time.Now()
  configurations2 := SimulateSandpileParallel(initial_gameboard)
  elapsed2 := time.Since(start2)
  log.Printf("Time taken %s", elapsed2)
  fmt.Println("Simulation run. Now drawing final stable configuration.")
  //in DrawPNGFinal, the second parameter 1 is the width of each cell in the board which is 1 pixel
  DrawPNGFinal(configurations2, 1, "sandpile_parallel.png")
  fmt.Println("Image drawn!")

  //Check if both serial and parallel are giving same output
  if GameboardsEqual(configurations, configurations2){
    fmt.Println("Serial and parallel results are equal!")
  }else{
    panic("Error: serial and parallel results are not equal!")
  }

}
