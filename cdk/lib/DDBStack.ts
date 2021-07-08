import * as cdk from '@aws-cdk/core';
import * as dynamodb from '@aws-cdk/aws-dynamodb';
import * as ssm from '@aws-cdk/aws-ssm'

export class DdbLocalTestStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const postTable = new dynamodb.Table(this, 'PostTable', {
      tableName: "blog-post-table",
      partitionKey: { name: 'id', type: dynamodb.AttributeType.STRING },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      //removalPolicy: cdk.RemovalPolicy.RETAIN,
      //pointInTimeRecovery: true
    });

    const ddbNameParam = new ssm.StringParameter(this, 'DDBNameParam', {
      description: 'DDB Table Name',
      parameterName: "/ddblocal/tableName",
      stringValue: postTable.tableName,
      tier: ssm.ParameterTier.STANDARD,
    });

    const ddbArnParam = new ssm.StringParameter(this, 'DDBArnParam', {
      description: 'DDB Table Arn',
      parameterName: "/ddblocal/tableArn",
      stringValue: postTable.tableArn,
      tier: ssm.ParameterTier.STANDARD,
    });



  }
}
