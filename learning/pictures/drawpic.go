package pictures

// import "https://github.com/golang/tour"
import "code.google.com/p/go-tour/pic"

func picx(dx, dy int) [][]uint8 {
	pix := [][]uint8{}
	for x := 0; x < dx; x++ {
		pix = append(pix, []uint8{})
		for y := 0; y < dy; y++ {
			pix[x] = append(pix[x], 1)
		}
	}
	return pix
}

func main() {
	Drawpic(256, 256)
}

func Drawpic(dx, dy int) { //draw stuff
	// Draw a pic

	println(string(byte(4)))
	// pic.Show(picx(dx, dy))
	out_pic = picx(dx, dy))
	pic.show(out_pic)
}
