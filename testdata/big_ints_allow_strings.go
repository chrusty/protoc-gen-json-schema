package testdata

const BigIntsAllowString = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name1": {
            "oneOf": [
                {
                    "type": "integer"
                },
                {
                    "type": "string"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
