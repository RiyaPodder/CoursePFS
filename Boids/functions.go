package main

import(
  "math"
  "math/rand"
  "time"
  "fmt"
)

//place your non-drawing functions here.

func InitializeSky(skyWidth, initialSpeed, maxBoidSpeed, proximity, separationFactor, alignmentFactor, cohesionFactor float64, numBoids int) Sky{

  var initialSky Sky

  initialSky.width = skyWidth
  initialSky.boids = make([]Boid, numBoids)
  initialSky.maxBoidSpeed = maxBoidSpeed
  initialSky.proximity = proximity
  initialSky.separationFactor = separationFactor
  initialSky.alignmentFactor = alignmentFactor
  initialSky.cohesionFactor = cohesionFactor

  for i := 0; i<numBoids; i++{

    rand.Seed(time.Now().UnixNano())
    theta := 2 * math.Pi * rand.Float64()

    initialSky.boids[i].velocity.x = initialSpeed * math.Cos(theta)
    initialSky.boids[i].velocity.y = initialSpeed * math.Sin(theta)

    initialSky.boids[i].position.x = float64(rand.Float64() * skyWidth)
    initialSky.boids[i].position.y = float64(rand.Float64() * skyWidth)

    initialSky.boids[i].acceleration.x = 0.0
    initialSky.boids[i].acceleration.y = 0.0

  }

  return initialSky
}

func SimulateBoids(initialSky Sky, numGens int, time float64) []Sky {
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialSky

	//now range over the number of generations and update the sky each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], time)
	}

	return timePoints
}

func UpdateSky(currentSky Sky, time float64) Sky {
	newSky := CopySky(currentSky)

	//range over all boids in the sky and update their acceleration,
	//velocity, and position
	for i := range newSky.boids {
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, newSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(newSky.boids[i], time, newSky)
		newSky.boids[i].position = UpdatePosition(newSky.boids[i], time)
	}

	return newSky
}

func CopySky(currentSky Sky) Sky {
	var newSky Sky

	newSky.width = currentSky.width

	//let's make the new sky's slice of Boid objects
	numBoids := len(currentSky.boids)
	newSky.boids = make([]Boid, numBoids)

	//now, copy all of the boids' fields into our new boids
	for i := range currentSky.boids {

		newSky.boids[i].position.x = currentSky.boids[i].position.x
		newSky.boids[i].position.y = currentSky.boids[i].position.y
		newSky.boids[i].velocity.x = currentSky.boids[i].velocity.x
		newSky.boids[i].velocity.y = currentSky.boids[i].velocity.y
		newSky.boids[i].acceleration.x = currentSky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = currentSky.boids[i].acceleration.y

	}

	return newSky
}

func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(currentSky, b)

	//now, calculate acceleration (F = ma)
	accel.x = force.x
	accel.y = force.y

  //fmt.Println("UpdateAcceleration: ", accel)

	return accel
}

func ComputeNetForce(currentSky Sky, b Boid) OrderedPair {
	var netForce OrderedPair
  var avg_sf, avg_af, avg_cf OrderedPair
  var nearbyBoids int

	for i := range currentSky.boids {
    //d := Distance(b.position, currentSky.boids[i].position)
    //fmt.Println("Distance: ", d)
		//only do a force computation if current body is not the input Body
		if currentSky.boids[i] != b /*&& d <= currentSky.proximity*/ {
      //fmt.Println("Hey!")
      nearbyBoids++

			separation_force := ComputeSeparationForce(b, currentSky.boids[i], currentSky.separationFactor)
      alignment_force := ComputeAlignmentForce(b, currentSky.boids[i], currentSky.alignmentFactor)
      cohesion_force := ComputeCohesionForce(b, currentSky.boids[i], currentSky.cohesionFactor)

      avg_sf.x += separation_force.x
      avg_sf.y += separation_force.y
      avg_af.x += alignment_force.x
      avg_af.y += alignment_force.y
      avg_cf.x += cohesion_force.x
      avg_cf.y += cohesion_force.y
		}
	}

  //write if condition for the case when nearbyBoids is zero, force should be zero
  if nearbyBoids == 0{
    netForce.x = 0
    netForce.y = 0
    fmt.Println("Hey2")
  }else {

    n := float64(nearbyBoids)

    avg_sf.x /= n
    avg_sf.y /= n
    avg_af.x /= n
    avg_af.y /= n
    avg_cf.x /= n
    avg_cf.y /= n

    netForce.x += avg_sf.x + avg_af.x + avg_cf.x
    netForce.y += avg_sf.y + avg_af.y + avg_cf.y

  }


	return netForce
}

func ComputeSeparationForce(b1, b2 Boid, sf float64) OrderedPair{
	var force OrderedPair

  d := Distance(b1.position, b2.position)
  force.x = sf * (b1.position.x - b2.position.x)/(d*d)
  force.y = sf * (b1.position.y - b2.position.y)/(d*d)


	return force
}

func ComputeAlignmentForce(b1, b2 Boid, af float64) OrderedPair {
	var force OrderedPair
  d := Distance(b1.position, b2.position)
  force.x = af * b2.velocity.x/d
  force.y = af * b2.velocity.y/d

	return force
}

func ComputeCohesionForce(b1, b2 Boid, cf float64) OrderedPair {
	var force OrderedPair
  d := Distance(b1.position, b2.position)
  force.x = cf * (b2.position.x - b1.position.x)/d
  force.y = cf * (b2.position.y - b1.position.y)/d

	return force
}

func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func UpdateVelocity(b Boid, time float64, newSky Sky) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

  vel_magnitude := math.Sqrt(vel.x*vel.x + vel.y*vel.y)

  if vel_magnitude > newSky.maxBoidSpeed{
    vel.x  = (vel.x/vel_magnitude) * newSky.maxBoidSpeed
    vel.y  = (vel.y/vel_magnitude) * newSky.maxBoidSpeed
  }

	return vel
}

func UpdatePosition(b Boid, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	pos.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

  //fmt.Println("Update position: ",pos)
	return pos
}
