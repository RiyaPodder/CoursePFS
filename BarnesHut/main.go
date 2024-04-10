package main

import "fmt"

func main(){

/*
  var t Tree
  fmt.Println(len(t))
  t = InitializeQuadTree()
  fmt.Println(len(t))



  pNodes := InitializeNodes()
  fmt.Println(pNodes)
  s := InitializeSector(pNodes, 100.0)
  fmt.Println(s.members)
  fmt.Println(s.width)
  */


  //fmt.Println(MakeQuadTree(20.0))

  t := MakeQuadTree(20.0)

  var currentNode Node
  currentNode.coordinate.x = 16
  currentNode.coordinate.y = 18
  currentNode.mass = 1.0

  fmt.Println("Starting ComputeNetForce")
  forcecheck:= ComputeNetForce(0.5, &currentNode, t)
  fmt.Println(forcecheck)
  var accelcheck OrderedPair
  accelcheck.x = forcecheck.x/currentNode.mass
  accelcheck.y = forcecheck.y/currentNode.mass
  fmt.Println(accelcheck)

  //fmt.Println(exfun2(17))

  //check()
  //-3.3035010583270587e-13 + -2.0022240000000002e-12 + 7.415644444444445e-12 + 4.2716694064879356e-12 + 5.034210673857891e-13 + -7.135523414322418e-14
  //-4.719287226181513e-14 + -2.6696320000000006e-12 + -2.847779604325291e-12 + -2.5171053369289458e-12 + -2.675821280370907e-13
}
