package main

import (
	"fmt"
	"math"
	"regexp"
	"sync"
	"time"

	"evgeni.com/util"
)

const (
	MAXSTEPS            = 100000
	VELOCTY_RANGE_MIN_X = 1
	VELOCTY_RANGE_MAX_X = 100
	VELOCTY_RANGE_MIN_Y = 1
	VELOCTY_RANGE_MAX_Y = 1000
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	maxYReachedChan := make(chan int)
	area := areaNew(input[0])
	go func() {
		couroutinesCounter := 0
		for velX := VELOCTY_RANGE_MIN_X; velX <= VELOCTY_RANGE_MAX_X; velX++ {
			for velY := VELOCTY_RANGE_MIN_Y; velY <= VELOCTY_RANGE_MAX_Y; velY++ {
				probe := newProbe(velX, velY)
				wg.Add(1)
				go probe.Shoot(area, &wg, maxYReachedChan)
				couroutinesCounter++

				if couroutinesCounter%1024 == 0 {
					fmt.Println(couroutinesCounter, " coroutines ran")
					time.Sleep(time.Millisecond)
				}
			}
		}
		wg.Wait()
		fmt.Println("All couroutines are done. Closing channel.")
		close(maxYReachedChan)
	}()

	maxYReached := math.MinInt
	for maxY := range maxYReachedChan {
		if maxY > maxYReached {
			fmt.Println("Received new max ", maxY)
			maxYReached = maxY
		}
	}
	fmt.Println()
	fmt.Println("Max height reached: ", maxYReached)
}

type probe struct {
	posX, posY, velX, velY int
}

func newProbe(velX, velY int) probe {
	var probe probe
	probe.velX = velX
	probe.velY = velY
	return probe
}

func (p *probe) Shoot(targetArea *area, wg *sync.WaitGroup, maxYReachedChan chan int) {
	localMax := math.MinInt
	for step := 0; step < MAXSTEPS; step++ {
		p.Step()
		if p.posY > localMax {
			localMax = p.posY
		}
		if targetArea.IsInside(p.posX, p.posY) {
			maxYReachedChan <- localMax
			break
		}
		if p.posY < targetArea.maxY || p.posX > targetArea.maxX {
			break
		}
	}
	wg.Done()
}

// The probe's x position increases by its x velocity.
// The probe's y position increases by its y velocity.
// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
// Due to gravity, the probe's y velocity decreases by 1.
func (p *probe) Step() {
	p.posX += p.velX
	p.posY += p.velY

	if p.velX > 0 {
		p.velX -= 1
	} else if p.velX < 0 {
		p.velX += 1
	}
	p.velY -= 1
}

type area struct {
	minX, maxX, minY, maxY int
}

func areaNew(input string) *area {
	areaMatcher := regexp.MustCompile(`target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	results := areaMatcher.FindStringSubmatch(input)
	var area area
	area.minX = util.ParseInt(results[1])
	area.maxX = util.ParseInt(results[2])
	area.minY = util.ParseInt(results[3])
	area.maxY = util.ParseInt(results[4])
	return &area
}

func (a *area) IsInside(x, y int) bool {
	return a.minX <= x && x <= a.maxX && a.minY <= y && y <= a.maxY
}
