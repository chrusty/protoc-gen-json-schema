package testdata

const EnumCeption = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name": {
            "properties": {},
            "type": "string"
        },
        "timestamp": {
            "properties": {},
            "type": "string"
        },
        "id": {
            "properties": {},
            "type": "integer"
        },
        "rating": {
            "properties": {},
            "type": "number"
        },
        "complete": {
            "properties": {},
            "type": "boolean"
        },
        "failureMode": {
            "properties": {},
            "enum": [
                "RECURSION_ERROR",
                0,
                "SYNTAX_ERROR",
                1
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        },
        "payload": {
            "$ref": "samples.PayloadMessage",
            "additionalProperties": true,
            "type": "object"
        },
        "payloads": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "$ref": "samples.PayloadMessage"
            },
            "properties": {},
            "type": "array"
        },
        "importedEnum": {
            "properties": {},
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object",
    "definitions": {
        "samples.PayloadMessage": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "properties": {},
                    "type": "string"
                },
                "timestamp": {
                    "properties": {},
                    "type": "string"
                },
                "id": {
                    "properties": {},
                    "type": "integer"
                },
                "rating": {
                    "properties": {},
                    "type": "number"
                },
                "complete": {
                    "properties": {},
                    "type": "boolean"
                },
                "topology": {
                    "properties": {},
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.PayloadMessage"
        }
    }
}`
