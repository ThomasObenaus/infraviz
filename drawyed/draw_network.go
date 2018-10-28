package drawyed

import (
	"github.com/thomasobenaus/infraviz/network"
)

func (d *YedDraw) DrawVPC(vpcs []network.VPC) {
	for _, vpc := range vpcs {

		awsVpc := vpc.Vpc.AwsVpc
		tfVpc := vpc.Vpc.TfVpc

		label := NewHTMLLabel()
		name := awsVpc.NameTag
		if len(awsVpc.NameTag) == 0 {
			name = "UNNAMED"
		}
		defaultStr := ""
		if awsVpc.IsDefaultVPC {
			defaultStr = "(default)"
		}
		label.AddStyledText(name+" "+defaultStr, FSBold)
		label.AddText(" [" + awsVpc.CIDR + "]")
		label.NewLine()
		label.AddText(awsVpc.VpcID)

		if tfVpc != nil {
			label.AddText("/ " + tfVpc.Name())
		}

		d.Rectangle(0, 0, 100, 100, label.String())
	}
}
