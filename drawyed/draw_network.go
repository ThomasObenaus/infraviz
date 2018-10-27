package drawyed

import (
	"github.com/thomasobenaus/infraviz/network"
)

func (d *YedDraw) DrawVPC(vpcs []network.VPC) {
	for _, vpc := range vpcs {

		awsVpc := vpc.Vpc.AwsVpc
		tfVpc := vpc.Vpc.TfVpc

		label := awsVpc.CIDR + "\n(" + awsVpc.VpcID

		if tfVpc != nil {
			label += "/" + tfVpc.Name()
		}
		label += ")"
		d.Rectangle(0, 0, 100, 100, label)

	}
}
