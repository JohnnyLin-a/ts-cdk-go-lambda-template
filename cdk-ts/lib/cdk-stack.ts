import { Stack, type StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import { Architecture, Runtime } from "aws-cdk-lib/aws-lambda";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    new GoFunction(this, "HelloWorldLambdaFunction", {
      entry: "../lambdas/go/hello-world",
      architecture: Architecture.ARM_64,
      runtime: Runtime.PROVIDED_AL2023,
      functionName: "HelloWorldLambdaFunction",
    });
  }
}
