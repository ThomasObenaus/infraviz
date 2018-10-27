package network

import aws "github.com/thomasobenaus/inframapper/aws"

type VPC struct {
	AwsVpc aws.Vpc
	Label  string
}
