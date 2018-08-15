// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンパスの大きさ
	cells         = 100                 // 格子のマスの数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // x ,y軸の角度　angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// そもそも無限大を返さないようにする方法をとってしまった・・・
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0) 引数の二乗和の平方根を返す
	if math.IsInf(r, 0) == true {
	}
	return math.Sin(r) / 10
}

//	func IsInf(f float64, sign int) bool
//	IsInfは、fがsignと一致する無限大であるかを返します。sign > 0のときはfが正の無限大かを返し、sign < 0のときはfが負の無限大かを返します。sign == 0のときはどちらかの無限大であるかを返します。
//	参考ページ　http://y0m0r.hateblo.jp/entry/20140921/1411296684

// func IsInfCheck(f float64, sign int) float64 {
//	if math.IsInf(f, 0) == true {
//	}
//	return f
//}

//!-
