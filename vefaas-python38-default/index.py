import json


def handler(event, context):
    result = {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json'
        },
        'body': json.dumps({
            'message': 'hello world'
        })
    }
    return result
