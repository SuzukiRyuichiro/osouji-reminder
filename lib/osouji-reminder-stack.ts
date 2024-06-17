import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { RestApi, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";
import { config } from "dotenv";
config();

export class OsoujiReminderStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myFunction = new lambda.Function(this, "MyLambda", {
      code: lambda.Code.fromAsset("lambdas"),
      handler: "bootstrap",
      runtime: lambda.Runtime.PROVIDED_AL2,
      architecture: lambda.Architecture.ARM_64,
      environment: {
        LINE_CHANNEL_SECRET: process.env.LINE_CHANNEL_SECRET!,
        LINE_CHANNEL_ACCESS_TOKEN: process.env.LINE_CHANNEL_ACCESS_TOKEN!,
      },
    });

    const gateway = new RestApi(this, "MyGateway", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTION"],
      },
    });

    const integration = new LambdaIntegration(myFunction);
    const testResource = gateway.root.addResource("test");

    testResource.addMethod("GET", integration);
  }
}
