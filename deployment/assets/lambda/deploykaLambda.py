import boto3
import json
import time
import os

tableName = "Deployka"
BLOCK = "block"
UNBLOCK = "unblock"


def lambda_handler(event, context):
    headers = event['headers']
    method = event['httpMethod']
    apikey = headers['x-api-key']

    dynamodb = boto3.resource('dynamodb')
    deploykaTable = dynamodb.Table(tableName)

    client = boto3.client('apigateway')
    apikeyvals = client.get_api_keys(
        includeValues=True
    )

    for item in apikeyvals['items']:
        if item['value'] == apikey:
            key = item
            break

    if not key['enabled']:
        return {
            "statusCode": 401,
            "body": json.dumps("Forbidden. Your key is invalid.")
        }

    if method == "POST":
        body = json.loads(event['body'])
        action = body['action']

        response = ""
        try:
            if action == "get-pipes":
                pipes = json.loads(os.environ['pipenames'])
                response = get_all_pipes(deploykaTable, pipes)

            if action == "block":
                response = update_pipes(deploykaTable, body['block'], key['name'], BLOCK)

            if action == "unblock":
                response = update_pipes(deploykaTable, body['unblock'], key['name'], UNBLOCK)

        except Exception as error:
            print(error)
            return {
                "statusCode": 500,
                "body": json.dumps(
                    f"Something went wrong on the server but the operation may have partially completed.\n Message: {error}")
            }
        print(response)
        return {
            "statusCode": 200,
            "body": json.dumps(response)
        }

    else:
        return {
            "statusCode": 200,
            "body": json.dumps(f"Hello, {key['name']}, says Deployka!")
        }


# Get all pipes from the table
def get_all_pipes(table, pipes):
    pipelist = []
    for pipe in pipes:
        pipelist.append(get_item(table, pipe))
    return pipelist


def get_item(table, pipename):
    res = table.get_item(
        Key={
            'Pipename': pipename
        }
    )
    return res['Item']


# Get a dict with each block list
def get_blocked_dict(table, pipes) -> dict:
    lists = {}
    for pipe in pipes:
        item = get_item(table, pipe)
        lists[pipe] = item['BlockedBy']

    return lists


# Either blocks or unblocks the passed pipes.
def update_pipes(table, pipes, user, operation):
    allowed_pipes = json.loads(os.environ['pipenames'])

    for pipe in pipes:
        if pipe not in allowed_pipes:
            pipes.remove(pipe)


    blockDict = get_blocked_dict(table, pipes)
    updates = []

    for pipe in pipes:
        allowed = False

        if user not in blockDict[pipe] and operation == BLOCK:
            blockDict[pipe].append(user)
        elif user in blockDict[pipe] and operation == UNBLOCK:
            blockDict[pipe].remove(user)
            if not blockDict[pipe]:
                allowed = True

        t = time.gmtime()
        timestamp = {'year':str(t.tm_year), 'month':str(t.tm_mon), 'day':str(t.tm_mday), 'hour':str(t.tm_hour), 'min':str(t.tm_min)}

        table.update_item(
            Key={
                'Pipename': pipe
            },
            UpdateExpression="SET BlockedBy = :b, Allowed = :a, BlockedAt = :t",
            ExpressionAttributeValues={
                ':b': blockDict[pipe],
                ':a': allowed,
                ':t': timestamp
            },
        )

        res = get_item(table, pipe)

        updates.append(res)

    return updates