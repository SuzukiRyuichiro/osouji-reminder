import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { RestApi, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";

export class OsoujiReminderStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myFunction = new lambda.Function(this, "MyLambda", {
      code: lambda.Code.fromAsset("lambdas"),
      handler: "bootstrap",
      runtime: lambda.Runtime.PROVIDED_AL2,
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
