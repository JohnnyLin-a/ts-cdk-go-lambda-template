package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	lambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type WorkshopStackProps struct {
	awscdk.StackProps
	// Add additional stack props if needed here
}

func NewWorkshopStack(scope constructs.Construct, id string, props *WorkshopStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	lambdago.NewGoFunction(stack, new("HelloWorldLambdaFunction"), &lambdago.GoFunctionProps{
		FunctionName: new("HelloWorldLambdaFunction"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Entry:        new("../lambdas/go/hello-world"),
	})
	return stack
}
