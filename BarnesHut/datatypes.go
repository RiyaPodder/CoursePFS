package main

const G = 6.67408e-11

type Node struct{ //Star

  coordinate OrderedPair //position
  mass float64
}

type OrderedPair struct{

  x,y float64
}

type Tree *Sector //QuadTree

type Sector struct{ //Node

  members []*Node //keep track of original nodes which make the dummy node, not included in starter code
  node     *Node
  //width float64
  //coordinate OrderedPair
  //mass float64
  child [4]*Sector
  sec_quad Quadrant

}

type Quadrant struct {
	x     float64 //bottom left corner x coordinate
	y     float64 //bottom right corner y coordinate
	width float64
}
