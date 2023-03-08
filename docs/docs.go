// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/kanji/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "summary": "get a kanji",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Kanji ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Kanji"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/meaning_mnemonic": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meaning-Mnemonic"
                ],
                "summary": "update the meaning mnemonic text",
                "parameters": [
                    {
                        "description": "mnemonic id + text",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateMeaningMnemonic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meaning-Mnemonic"
                ],
                "summary": "create a meaning mnemonic for a kanji or vocabulary",
                "parameters": [
                    {
                        "description": "Meaning mnemonic",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateMeaningMnemonic"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.MeaningMnemonic"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meaning-Mnemonic"
                ],
                "summary": "delete a meaning mnemonic",
                "parameters": [
                    {
                        "description": "mnemonic id",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteMeaningMnemonic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/meaning_mnemonic/toggle_favorite": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meaning-Mnemonic"
                ],
                "summary": "toggle favorite on a meaning mnemonic",
                "parameters": [
                    {
                        "description": "mnemonic id",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ToggleFavoriteMeaningMnemonic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/meaning_mnemonic/vote": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meaning-Mnemonic"
                ],
                "summary": "vote on a meaning mnemonic",
                "parameters": [
                    {
                        "description": "vote can be 1 or -1",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VoteMeaningMnemonic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/radical/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "summary": "get a radical",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Radical ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Radical"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "search for subjects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListResponse-dto_SubjectPreview"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/subject/{id}/meaning_mnemonics": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "summary": "get meaning mnemonics optionally with user data if authenticated",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subject ID vocabulary or kanji",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListResponse-dto_MeaningMnemonicWithUserInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vocabulary/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "summary": "get a vocabulary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Vocabulary ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Vocabulary"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateMeaningMnemonic": {
            "type": "object",
            "required": [
                "object",
                "subject_id",
                "text"
            ],
            "properties": {
                "object": {
                    "$ref": "#/definitions/model.Object"
                },
                "subject_id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dto.DeleteMeaningMnemonic": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.Kanji": {
            "type": "object",
            "properties": {
                "amalgamation_subjects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                },
                "characters": {
                    "type": "string"
                },
                "component_subjects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "meanings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Meaning"
                    }
                },
                "object": {
                    "$ref": "#/definitions/model.Object"
                },
                "reading_mnemonic": {
                    "type": "string"
                },
                "readings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.KanjiReading"
                    }
                },
                "slug": {
                    "type": "string"
                },
                "visually_similar_subjects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                }
            }
        },
        "dto.ListResponse-dto_MeaningMnemonicWithUserInfo": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MeaningMnemonicWithUserInfo"
                    }
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "dto.ListResponse-dto_SubjectPreview": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "dto.MeaningMnemonic": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "subject_id": {
                    "description": "must be string ensure that TAG for the index works as expected",
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "voting_count": {
                    "type": "integer"
                }
            }
        },
        "dto.MeaningMnemonicWithUserInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "downvoted": {
                    "type": "boolean"
                },
                "favorite": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "me": {
                    "type": "boolean"
                },
                "subject_id": {
                    "description": "must be string ensure that TAG for the index works as expected",
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "upvoted": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "string"
                },
                "voting_count": {
                    "type": "integer"
                }
            }
        },
        "dto.Radical": {
            "type": "object",
            "properties": {
                "amalgamation_subjects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                },
                "character_image": {
                    "type": "string"
                },
                "characters": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "meaning_mnemonic": {
                    "type": "string"
                },
                "meanings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Meaning"
                    }
                },
                "object": {
                    "$ref": "#/definitions/model.Object"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "dto.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.SubjectPreview": {
            "type": "object",
            "properties": {
                "character_image": {
                    "type": "string"
                },
                "characters": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "meanings": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "object": {
                    "$ref": "#/definitions/model.Object"
                },
                "readings": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "dto.ToggleFavoriteMeaningMnemonic": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateMeaningMnemonic": {
            "type": "object",
            "required": [
                "id",
                "text"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dto.Vocabulary": {
            "type": "object",
            "properties": {
                "characters": {
                    "type": "string"
                },
                "component_subjects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SubjectPreview"
                    }
                },
                "context_sentences": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContextSentence"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "meanings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Meaning"
                    }
                },
                "object": {
                    "$ref": "#/definitions/model.Object"
                },
                "reading_mnemonic": {
                    "type": "string"
                },
                "readings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VocabularyReading"
                    }
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "dto.VoteMeaningMnemonic": {
            "type": "object",
            "required": [
                "meaning_mnemonic_id",
                "vote"
            ],
            "properties": {
                "meaning_mnemonic_id": {
                    "type": "string"
                },
                "vote": {
                    "type": "integer"
                }
            }
        },
        "model.ContextSentence": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string"
                },
                "hiragana": {
                    "type": "string"
                },
                "ja": {
                    "type": "string"
                }
            }
        },
        "model.KanjiReading": {
            "type": "object",
            "properties": {
                "primary": {
                    "type": "boolean"
                },
                "reading": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Meaning": {
            "type": "object",
            "properties": {
                "meaning": {
                    "type": "string"
                },
                "primary": {
                    "type": "boolean"
                }
            }
        },
        "model.Object": {
            "type": "string",
            "enum": [
                "kanji",
                "radical",
                "vocabulary"
            ],
            "x-enum-varnames": [
                "ObjectKanji",
                "ObjectRadical",
                "ObjectVocabulary"
            ]
        },
        "model.VocabularyReading": {
            "type": "object",
            "properties": {
                "primary": {
                    "type": "boolean"
                },
                "reading": {
                    "type": "string"
                },
                "romaji": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "robanohashi.org",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Roba no hashi API",
	Description:      "Query Kanji, Vocabulary, and Radicals with Mnemonics",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
