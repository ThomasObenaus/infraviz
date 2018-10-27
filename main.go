package main

import (
	"os"

	"github.com/thomasobenaus/inframapper/mappedInfra"
	"github.com/thomasobenaus/inframapper/tfstate"
	"github.com/thomasobenaus/inframapper/trace"
	yed "github.com/thomasobenaus/infraviz/drawyed"
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

	mappedInfra, err := mappedInfra.LoadAndMap(profile, region, remoteCfg, tracerMapper)

	tracer := trace.NewInfo(os.Stdout)
	if err != nil {
		tracer.Info("Error loading/ mapping infrastructure: %s", err.Error())
	}

	var mappedResStr string
	var unMappedAwsResStr string
	for _, res := range mappedInfra.MappedResources() {
		mappedResStr += "\t" + res.String() + "\n"
	}
	for _, res := range mappedInfra.UnMappedAwsResources() {
		unMappedAwsResStr += "\t" + res.String() + "\n"
	}

	tracer.Info("Mapped Infra: ", mappedInfra)
	tracer.Info("Mapped Resources [", len(mappedInfra.MappedResources()), "]:")
	tracer.Info(mappedResStr)
	tracer.Info("UnMapped AWS Resources [", len(mappedInfra.UnMappedAwsResources()), "]:")
	tracer.Info(unMappedAwsResStr)

	doc := yed.NewInitializedDocument()

	nodeKeyID := doc.KeyIDNodeGraphics()

	node := yed.Node{ID: "n0"}
	snode := &yed.ShapeNode{}
	snode.Geometry = yed.Geometry{Height: 30.0, Width: 30.0, X: 800.5, Y: 350}
	snode.Shape = yed.Shape{ShapeType: yed.Rectangle}
	nodeLabel := yed.NewNodeLabel("Hello World")
	snode.NodeLabel = &nodeLabel
	node.Data = []yed.Data{yed.Data{Key: nodeKeyID, ShapeNode: snode}}

	doc.AddNode(node)

	if err = doc.Encode(os.Stdout); err != nil {
		tracer.Error("Failed to encode: ", err.Error())
	}

	file, err := os.Create("testdata/rectangle.graphml")
	if err != nil {
		tracer.Error("Failed to open file: ", err.Error())
	}
	defer file.Close()

	if err = doc.Encode(file); err != nil {
		tracer.Error("Failed to encode: ", err.Error())
	}
}
