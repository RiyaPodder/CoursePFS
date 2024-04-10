package main
import (
"fmt"
"math"
)

//InitializeQuadTree appends a dummy object to a new tree. This dummy object is the root of the tree.
//Input: First Sector
//Output: Initial Tree with one root node only
/*
func InitializeQuadTree(s Sector) Tree{

  pointerToFirstSector := &s

  var t Tree
  t = append(t, pointerToFirstSector)
  return t
}
*/
//InitializeNodes initializes the masses and coordinate positions of a set of nodes.
//Input: Nothing, hard coded the values
//Output: Slice of pointers to the nodes intialized

func InitializeNodes() []*Node{

  var A, B, C, D, E, F, G Node

  A.coordinate.x = 2
  A.coordinate.y = 16
  A.mass = 1.0
  B.coordinate.x = 16
  B.coordinate.y = 18
  B.mass = 1.0
  C.coordinate.x = 19
  C.coordinate.y = 18
  C.mass = 1.0
  D.coordinate.x = 19
  D.coordinate.y = 16
  D.mass = 1.0
  E.coordinate.x = 17
  E.coordinate.y = 13
  E.mass = 1.0
  F.coordinate.x = 12
  F.coordinate.y = 3
  F.mass = 1.0
  G.coordinate.x = 10
  G.coordinate.y = 10
  G.mass = 5.0

  var pA, pB, pC, pD, pE, pF, pG *Node
  pA = &A
  pB = &B
  pC = &C
  pD = &D
  pE = &E
  pF = &F
  pG = &G

  var pNodes []*Node
  pNodes = append(pNodes, pA, pB, pC, pD, pE, pF, pG)

  return pNodes
}


//Sectors!
//Intialize Sector 1
//Input: all the node pointers
//Output: first sector
func InitializeSector(pNodes []*Node, quad Quadrant) Sector{

  var s Sector
  s.members = pNodes
  s.sec_quad = quad
  var dummyfirst Node
  s.node = &dummyfirst
  s.node.mass = CalcMass(s.members)
  s.node.coordinate = CalcCoordinate(s.members, s.node.mass)
  return s

}

//Recursive function which breaks the sectors into single element or nil element sectors
//Input: First Sector
//Output: Tree
func MakeQuadTree(width float64) Tree {

  pNodes := InitializeNodes()
  var quad Quadrant
  quad.x = 0.0
  quad.y = 0.0
  quad.width = width
  s := InitializeSector(pNodes, quad)

  var qt Tree
  qt = &s

  RecFunc(qt)
  printMems(qt,0)
  return qt

}



func RecFunc(s Tree) int{

  //i++
  //*t = append(*t,s)
  //fmt.Println("len(s.members) ",len(s.members))
  //s := *ps

  //fmt.Println(i)
  if len(s.members) == 0 || len(s.members) == 1{
    //return t
    return 1
  }else{
    //fmt.Println("hey!")
    //go through each member and put it in one of the child sectors and also update width of the sector
    //w := s.width
    //fmt.Println("width ",w)
    //do we need to make slice s.child?
    //s.child = make([]*Sector,4)

    //set quadrant
    //var northwest, northeast, southwest, southeast Quadrant
    var nw, ne, sw, se Quadrant
    /*
    nw = &northwest
    ne = &northeast
    sw = &southwest
    se = &southeast
*/
    nw, ne, sw, se = SetQuadrant(nw,ne,sw,se, s.sec_quad)
    fmt.Println("The quadrants ", nw, ne, sw, se)


    for i := 0; i < 4; i++{
      var dummy Sector
      s.child[i] = &dummy
    }

    for _,member := range s.members{
      //check the four conditions using member's coordinates
      x := member.coordinate.x
      y := member.coordinate.y
      //fmt.Println("inside for width ",w)
      //fmt.Println("inside for x", x)
      //fmt.Println("inside for y", y)
/*
      if member == nil{
        fmt.Println("Panic!")
      }else{
        fmt.Println("Continue!")
      }
*/
      if x >= nw.x && x <= (nw.x+nw.width) && y >= nw.y && y <= (nw.y+nw.width){
        fmt.Println("One!")
        s.child[0].members = append(s.child[0].members, member)
        s.child[0].sec_quad = nw
      }else if x > ne.x && x <= (ne.x+ne.width) && y >= ne.y && y <= (ne.y+ne.width){
        fmt.Println("Two!")
        s.child[1].members = append(s.child[1].members, member)
        s.child[1].sec_quad = ne
      }else if x >= sw.x && x <= (sw.x+sw.width) && y >= sw.y && y < (sw.y+sw.width){
        fmt.Println("Three!")
        s.child[2].members = append(s.child[2].members, member)
        s.child[2].sec_quad = sw
      }else if x > se.x && x <= (se.x+se.width) && y >= se.y && y < (se.y+se.width){
        fmt.Println("Four!")
        s.child[3].members = append(s.child[3].members, member)
        s.child[3].sec_quad = se
      }
    }
    for i := 0; i < 4; i++{
      var dummyNode Node
      s.child[i].node = &dummyNode
    }
    for _,childSector := range s.child{

        //childSector.width = w/2.0
        //var dummyNode Node
        //childSector.node = &dummyNode
        childSector.node.mass = CalcMass(childSector.members)
        childSector.node.coordinate = CalcCoordinate(childSector.members, childSector.node.mass)
    }
    for _,childSector := range s.child{

        RecFunc(childSector)
        //return t
    }
    return 0
  }
}

func CalcMass(members []*Node) float64{

  var centerOfMass float64

  for _,m := range members{

    centerOfMass += m.mass
  }

  return centerOfMass
}

func CalcCoordinate(members []*Node, totalMass float64) OrderedPair{

  var centerOfGravity OrderedPair
  var sum OrderedPair

  for _,m := range members{
    sum.x += m.coordinate.x * m.mass
    sum.y += m.coordinate.y	* m.mass
  }

	centerOfGravity.x = sum.x/totalMass
  centerOfGravity.y = sum.y/totalMass

  return centerOfGravity
}

func SetQuadrant(nw,ne,sw,se,currentQuad Quadrant) (Quadrant, Quadrant, Quadrant, Quadrant){

  //north-west quadrant
  nw.x = currentQuad.x
  nw.y = currentQuad.y + currentQuad.width/2.0
  nw.width = currentQuad.width/2.0

  //north-east quadrant
  ne.x = currentQuad.x + currentQuad.width/2.0
  ne.y = currentQuad.y + currentQuad.width/2.0
  ne.width = currentQuad.width/2.0

  //south-west quadrant
  sw.x = currentQuad.x
  sw.y = currentQuad.y
  sw.width = currentQuad.width/2.0

  //south-east quadrant
  se.x = currentQuad.x + currentQuad.width/2.0
  se.y = currentQuad.y
  se.width = currentQuad.width/2.0

  return nw, ne, sw, se


}

func printMems(s Tree, i int) int{

  i++
  fmt.Println("sector ",i)

  //fmt.Println of coordinates of members of s
  for _, m := range s.members{
    fmt.Println(m.coordinate)
  }
  if s.node.mass == 0.0{
    fmt.Println("Empty node")
  }
  fmt.Println("sector coordinate and mass ", s.node)
  fmt.Println("sector quadrant", s.sec_quad)
  //fmt.Println("mass ", s.node.mass)

  if s.child[0] == nil && s.child[1] == nil && s.child[2] == nil && s.child[3] == nil{
    return 0
  }else{
    for _,c := range s.child{
      printMems(c,i)
    }
    return 1
  }
}
/*
func exfun2(a int)int{

  if a<10{
    if a<5{
      return 5
    }else{
      return 10
    }
  }else{
    if a<20{
      return 20
    }else{
      return 15
    }
  }
}
*/
/*
func ComputeNetForce(theta float64, currentNode *Node, t Tree)OrderedPair{

    dist := Distance(currentNode.coordinate,t.node.coordinate)

    //if its a leaf node
    if t.child[0] == nil && t.child[1] == nil && t.child[2] == nil && t.child[3] == nil {
      //if its Empty node
      if t.node.mass == 0.0{
        //return no force
        var dummyForce OrderedPair
			  dummyForce.x = 0.0
			  dummyForce.y = 0.0
        fmt.Println("A ",t.node)
			  return dummyForce
      }else{
        if dist != 0.0{
        fmt.Println("B ",t.node)
        bforce := ComputeForce(t.node, currentNode, dist)
        fmt.Println("bforce ", bforce)
        return bforce
      }else{
        var dummyForce OrderedPair
			  dummyForce.x = 0.0
			  dummyForce.y = 0.0
        fmt.Println("C ",t.node)
			  return dummyForce
      }
      }
    }else{ //internal node

      s := t.sec_quad.width
      if dist != 0.0{
      if (s/dist) <= theta{
        fmt.Println("D ",t.node)
        dforce := ComputeForce(t.node, currentNode, dist)
        fmt.Println("dforce ", dforce)
        return dforce
      }else{
        var totalForce OrderedPair
        for _,child := range t.child{
          f := ComputeNetForce(theta, currentNode, child)
          totalForce.x += f.x
          totalForce.y += f.y
        }
        return totalForce
      }
    }else{
      var dummyForce OrderedPair
      dummyForce.x = 0.0
      dummyForce.y = 0.0
      return dummyForce
    }
    }

  }
*/

func ComputeForce(nearbyStar, currentStar *Node, d float64) OrderedPair{

	var force OrderedPair
	//d := Distance(currentStar.position, nearbyStar.position)

	F  := (G * currentStar.mass * nearbyStar.mass) / (d * d)

	deltaX := nearbyStar.coordinate.x - currentStar.coordinate.x
	deltaY := nearbyStar.coordinate.y - currentStar.coordinate.y

	force.x = F * deltaX / d //deltaX/dist = cos theta
	force.y = F * deltaY / d //deltaY/dist = sin theta
	return force

}

func Distance(p1, p2 OrderedPair) float64{

	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)

}

func check(){

  var s []string
  s = append(s, "hello")
  fmt.Println(s)
  s = nil
  fmt.Println(s)
}

func ComputeNetForce(theta float64, currentNode *Node, t Tree)OrderedPair{

    dist := Distance(currentNode.coordinate,t.node.coordinate)

    //if its a leaf node
    if t.child[0] == nil && t.child[1] == nil && t.child[2] == nil && t.child[3] == nil {
      //if its Empty node
      if t.node.mass == 0.0{
        //return no force
        var dummyForce OrderedPair
			  dummyForce.x = 0.0
			  dummyForce.y = 0.0
        fmt.Println("A ",t.node)
			  return dummyForce
      }else{
        if dist != 0.0{
        	bforce := ComputeForce(t.node, currentNode, dist)
          fmt.Println("B ",t.node)
					fmt.Println("bforce ",bforce)
        	return bforce
      	}else{
        	var dummyForce OrderedPair
			  	dummyForce.x = 0.0
			  	dummyForce.y = 0.0
          fmt.Println("C ",t.node)
			  	return dummyForce
      	}
      }
    }else{ //internal node

      s := t.sec_quad.width
      if dist != 0.0{
      	if (s/dist) <= theta{
        	dforce := ComputeForce(t.node, currentNode, dist)
          fmt.Println("D ",t.node)
          fmt.Println("s ",s)
          fmt.Println("dist ",dist)
          fmt.Println("theta ", s/dist)
					fmt.Println("dforce ",dforce)
        	return dforce
      	}else{
        	var totalForce OrderedPair
          fmt.Println("E ",t.node)
          fmt.Println("s ",s)
          fmt.Println("dist ",dist)
          fmt.Println("theta ", s/dist)
        	for _,child := range t.child{
          	f := ComputeNetForce(theta, currentNode, child)
          	totalForce.x += f.x
          	totalForce.y += f.y
        	}
					fmt.Println("totalForce ",totalForce)
        	return totalForce
      	}
    	}else{
        fmt.Println("F ",t.node)
      	var dummyForce OrderedPair
      	dummyForce.x = 0.0
      	dummyForce.y = 0.0
      	return dummyForce
    	}
    }
  }
