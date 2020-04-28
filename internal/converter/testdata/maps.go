package testdata

const Maps = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "map_of_strings": {
            "properties": {},
            "additionalProperties": {
                "properties": {},
                "type": "string"
            },
            "type": "object"
        },
        "map_of_ints": {
            "properties": {},
            "additionalProperties": {
                "properties": {},
                "type": "integer"
            },
            "type": "object"
        },
        "map_of_messages": {
            "properties": {},
            "additionalProperties": {
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
                "type": "object"
            },
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
