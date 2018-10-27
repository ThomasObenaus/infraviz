package main

import (
	"os"

	"github.com/thomasobenaus/inframapper/mappedInfra"
	"github.com/thomasobenaus/inframapper/tfstate"
	"github.com/thomasobenaus/inframapper/trace"
	gml "github.com/thomasobenaus/infraviz/graphml"
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

	nodeKeyID := "d6"
	nodeKey := gml.NewYedNodeKey(nodeKeyID)

	node := gml.Node{ID: "n0"}
	snode := &gml.ShapeNode{}
	snode.Geometry = gml.Geometry{Height: 30.0, Width: 30.0, X: 800.5, Y: 350}
	snode.Shape = gml.Shape{ShapeType: gml.Rectangle}
	nodeLabel := gml.NewNodeLabel("Hello World")
	snode.NodeLabel = &nodeLabel
	node.Data = []gml.Data{gml.Data{Key: nodeKeyID, ShapeNode: snode}}

	graph := gml.Graph{EdgeDefault: gml.Directed, ID: "G"}
	graph.Nodes = append(graph.Nodes, node)
	doc := gml.NewEmptyDocument()
	doc.Graphs = append(doc.Graphs, graph)
	doc.Keys = append(doc.Keys, nodeKey)

	if err = gml.Encode(os.Stdout, doc); err != nil {
		tracer.Error("Failed to encode: ", err.Error())
	}
}
