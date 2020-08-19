package pkgx

// import "https://github.com/golang/tour"
import (
	"log"

	"golang.org/x/tour/pic"
)

func picx(dx, dy int) [][]uint8 {
	pix := [][]uint8{}
	for x := 0; x < dx; x++ {
		pix = append(pix, []uint8{})
		for y := 0; y < dy; y++ {
			pix[x] = append(pix[x], uint8(float32(x)/float32(y)))
		}
	}
	return pix
}

// Drawpic stuff
func Drawpic(dx, dy int) {

	log.Println("Drawpic(dx,dy) of learning2/importfolder/drawpic.go")
	// outpic := picx(dx, dy)
	pic.Show(picx)
}
