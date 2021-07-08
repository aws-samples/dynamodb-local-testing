#!/usr/bin/env node

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { DdbLocalTestStack } from '../lib/DDBStack';
import {FargateServiceStack} from "../lib/FargateServiceStack";

const app = new cdk.App();
new DdbLocalTestStack(app, 'ddblocal-ddb-stack');
new FargateServiceStack(app, 'ddblocal-fargate-stack');

