package main

import (
	"os"

	"github.com/thomasobenaus/inframapper/mappedInfra"
	"github.com/thomasobenaus/inframapper/tfstate"
	"github.com/thomasobenaus/inframapper/trace"
	yed "github.com/thomasobenaus/infraviz/drawyed"
	nw "github.com/thomasobenaus/infraviz/network"
)

func main() {

	profile := "develop"
	region := "eu-central-1"
	tracerMapper := trace.New(os.Stdout)

	keys := make([]string, 2)
	keys[0] = "snapshot/base/networking/terraform.tfstate"
	keys[1] = "snapshot/base/common/terraform.tfstate"
	remoteCfg := tfstate.RemoteConfig{
		BucketName: "741125603121-tfstate",
		Keys:       keys,
		Profile:    "shared",
		Region:     "eu-central-1",
	}

	mInfra, err := mappedInfra.LoadAndMap(profile, region, remoteCfg, tracerMapper)

	tracer := trace.NewInfo(os.Stdout)
	if err != nil {
		tracer.Info("Error loading/ mapping infrastructure: %s", err.Error())
	}

	var mappedResStr string
	var unMappedAwsResStr string
	for _, res := range mInfra.MappedResources() {
		mappedResStr += "\t" + res.String() + "\n"
	}
	for _, res := range mInfra.UnMappedAwsResources() {
		unMappedAwsResStr += "\t" + res.String() + "\n"
	}

	tracer.Info("Mapped Infra: ", mInfra)
	tracer.Info("Mapped Resources [", len(mInfra.MappedResources()), "]:")
	tracer.Info(mappedResStr)
	tracer.Info("UnMapped AWS Resources [", len(mInfra.UnMappedAwsResources()), "]:")
	tracer.Info(unMappedAwsResStr)

	vpcs := make([]nw.VPC, 0)
	for _, mRes := range mInfra.Resources() {
		if mappedInfra.TypeVPC != mRes.ResourceType() {
			continue
		}

		mappedVpc, err := mappedInfra.ToVpc(mRes)
		if err != nil {
			tracer.Error("Unable to cast to Vpc")
			continue
		}
		vpc := nw.VPC{Vpc: mappedVpc}
		vpcs = append(vpcs, vpc)
	}

	file, err := os.Create("testdata/rectangle.graphml")
	if err != nil {
		tracer.Error("Failed to open file: ", err.Error())
	}
	defer file.Close()

	yedd := yed.NewYedDraw(file)

	yedd.DrawVPC(vpcs)

	if err = yedd.Render(); err != nil {
		tracer.Error("Failed to render: ", err.Error())
	}
}
