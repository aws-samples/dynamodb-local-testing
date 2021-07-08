// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

import * as cdk from '@aws-cdk/core';
import * as path from 'path';
import * as ec2 from '@aws-cdk/aws-ec2';
import * as ecs from '@aws-cdk/aws-ecs';
import * as iam from '@aws-cdk/aws-iam'
import * as ssm from '@aws-cdk/aws-ssm'
import { LogGroup } from '@aws-cdk/aws-logs'
import { DockerImageAsset } from '@aws-cdk/aws-ecr-assets';
import { ContainerImage, FargatePlatformVersion } from '@aws-cdk/aws-ecs';
import * as ecs_patterns from "@aws-cdk/aws-ecs-patterns";



export class FargateServiceStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const asset = new DockerImageAsset(this, 'ddb-local-image', {
      directory: path.join(__dirname, "..", ".."),
    });

    const vpc = new ec2.Vpc(this, "EcsVpc", {
      maxAzs: 3 // Default is all AZs in region
    });

    const cluster = new ecs.Cluster(this, "TestCluster", {
      vpc: vpc,
      clusterName: "ddb-local-cluster",
      containerInsights: false
    });

    const logGroup = new LogGroup(this, "FargateLogGroup", {
      logGroupName: "/ecs/ddb-local-service"
    })

    const taskDef = new ecs.FargateTaskDefinition(this, "MyTask", {
      cpu: 512,
      memoryLimitMiB: 1024,
    })

    const container = new ecs.ContainerDefinition(this, "MyContainer", {
      image: ContainerImage.fromDockerImageAsset(asset),
      taskDefinition: taskDef,
      environment: {
        PARAM1: "test1"
      },
      logging: new ecs.AwsLogDriver({
        logGroup: logGroup,
        streamPrefix: `ddb-local-service`,
      })
    }
    )

    const service = new ecs_patterns.ApplicationLoadBalancedFargateService(this, "MyFargateService", {
      cluster: cluster, // Required
      cpu: 512, // Default is 256
      desiredCount: 2, // Default is 1
      taskImageOptions: { image: ecs.RepositoryImage.fromDockerImageAsset(asset) },
      memoryLimitMiB: 2048, // Default is 512
      publicLoadBalancer: true // Default is false
      li
    });


    const DDB_TABLE_ARN = ssm.StringParameter.valueForStringParameter(this, "/ddblocal/tableArn")


    service.taskDefinition.addToTaskRolePolicy(new iam.PolicyStatement({
      actions:["dynamodb:PutItem","dynamodb:GetItem"],
      resources:[DDB_TABLE_ARN],
      effect: iam.Effect.ALLOW
    }))

    const serviceDNS = new ssm.StringParameter( this,"serviceDNS",{
      parameterName: "local-ddb-service-dns",
      stringValue: service.loadBalancer.loadBalancerDnsName
    })
  }
}
