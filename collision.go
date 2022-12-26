package main

import (
	"alienGame/entity"
)

// CheckCollision 检查A与B是否发烧碰撞，
// A的4个顶点只要有一个位于B的矩形中，就认为它们碰撞
func CheckCollision(A, B entity.Entity) bool {
	aWidth, aHeight, aX, aY := A.GetInfo()
	bWidth, bHeight, bX, bY := B.GetInfo()

	top, left := aY, aX
	bottom, right := aY+float64(aHeight), aX+float64(aWidth)
	// 左上角
	x, y := bX, bY
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右上角
	x, y = bX+float64(bWidth), bY
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 左下角
	x, y = bX, bY+float64(bHeight)
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右下角
	x, y = bX+float64(bWidth), bY+float64(bHeight)
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	return false
}
