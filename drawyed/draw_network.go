package drawyed

import (
	"github.com/thomasobenaus/infraviz/network"
)

func (d *YedDraw) DrawVPC(vpcs []network.VPC) {

	style := defaultNodeLabelStyle
	style.LabelModelPosition = LMPTopLeft
	d.doc.PushNodeLabelStyle(&style)

	width := float32(500)

	for idx, vpc := range vpcs {
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

		x := width*float32(idx) + 80.0*float32(idx)
		d.Rectangle(x, 0, width, 150, label.String())
	}

	d.doc.PopNodeLabelStyle()
}
