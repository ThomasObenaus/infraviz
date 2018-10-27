package drawyed

import (
	"strconv"

	"github.com/thomasobenaus/infraviz/network"
)

func (d *YedDraw) DrawVPC(vpcs []network.VPC) {
	//doc := &d.doc

	for idx, _ := range vpcs {

		label := "vpc.CIDR" + strconv.Itoa(idx)

		d.Rectangle(0, 0, 100, 100, label)

	}
}
