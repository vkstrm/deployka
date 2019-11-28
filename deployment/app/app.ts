#!/usr/bin/env node
import cdk = require('@aws-cdk/core');
import * as config from '../deploy-config.json';
import { DeploykaStack } from '../stacks/deployka.js';

const app = new cdk.App();

const stackEnv = {
    env: {
        account: config.account,
        region: config.region
    }
}

new DeploykaStack(app, 'deployka', stackEnv);