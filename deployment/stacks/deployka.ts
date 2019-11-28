import cdk = require('@aws-cdk/core')
import lambda = require('@aws-cdk/aws-lambda')
import apigateway = require('@aws-cdk/aws-apigateway')
import certmanager = require('@aws-cdk/aws-certificatemanager')
import route53 = require('@aws-cdk/aws-route53')
import route53target = require('@aws-cdk/aws-route53-targets')
import dynamodb = require('@aws-cdk/aws-dynamodb')

import * as config from '../deploy-config.json';
import { ServicePrincipal, PolicyStatement } from '@aws-cdk/aws-iam'


/**
 * Deployka stack with various stuff
 */
export class DeploykaStack extends cdk.Stack {
    constructor(
        scope: cdk.Construct,
        id: string,
        props?: cdk.StackProps) {
        
        super(scope, id, props);


        // The lambda behind the gateway, doing the work
        const deploykaLambda = new lambda.Function(this, "deploykalambda", {
            code: lambda.AssetCode.fromAsset("./assets/lambda"),
            runtime: lambda.Runtime.PYTHON_3_7,
            handler: "deploykaLambda.lambda_handler",
            functionName: "deployka"
        });

        // Custom domain name stuff
        // const hostedZoneName = config.hostedZoneName;
        // const recordsetName = config.recordsetName;
        // const deploykaDomainName = `${recordsetName}.${hostedZoneName}`;

        // The API Gateway
        const deploykaGateway = new apigateway.RestApi(this, 'deploykaGateway', {
            restApiName: 'DeploykaAPIGateway',
            endpointTypes: [apigateway.EndpointType.REGIONAL],
            deploy: true,
            // domainName: {
            //     domainName: deploykaDomainName,
            //     endpointType: apigateway.EndpointType.REGIONAL,
            //     certificate: certmanager.Certificate.fromCertificateArn(this, 'deploykaApiCert', config.certificateArn)
            // }
        });

        deploykaLambda.addToRolePolicy(new PolicyStatement({
            actions: ["apigateway:GET"],
            resources: ["*"]
        }));

        const integration = new apigateway.LambdaIntegration(deploykaLambda);

        // Stuff for the API Gateway
        const resource = deploykaGateway.root.addResource('dep');
        resource.addMethod('GET', integration, {
            apiKeyRequired: true
        });
        resource.addMethod('POST', integration, {
            apiKeyRequired: true
        });

        const plan = deploykaGateway.addUsagePlan('deploykaUsagePlan', {
            name: 'Deployka - Standard',
        });

        plan.addApiStage({
            stage: deploykaGateway.deploymentStage
        });

        // More custom domain name stuff
        // const hostedZone = route53.HostedZone.fromLookup(this, 'depHostedZone', {
        //     domainName: hostedZoneName
        // })

        // new route53.RecordSet(this, 'deplyokaRecordset', {
        //     recordName: recordsetName,
        //     recordType: route53.RecordType.A,
        //     zone: hostedZone,
        //     target: route53.RecordTarget.fromAlias(new route53target.ApiGateway(deploykaGateway))
        // });

        // Dynamodb table for storing stuff
        const deploykaTable = new dynamodb.Table(this, 'deploykaTable', {
            partitionKey: {
                name: 'Pipename',
                type: dynamodb.AttributeType.STRING
            },
            tableName: 'Deployka'
        });            

        deploykaTable.grantReadWriteData(deploykaLambda)
    }
}
