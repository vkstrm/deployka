import { expect as expectCDK, matchTemplate, MatchStyle } from '@aws-cdk/assert';
import cdk = require('@aws-cdk/core');
import Deployment = require('../lib/deployment-stack');

test('Empty Stack', () => {
    const app = new cdk.App();
    // WHEN
    const stack = new Deployment.DeploymentStack(app, 'MyTestStack');
    // THEN
    expectCDK(stack).to(matchTemplate({
      "Resources": {}
    }, MatchStyle.EXACT))
});