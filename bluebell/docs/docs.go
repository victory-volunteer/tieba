// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "liwenzhou",
            "url": "http://www.liwenzhou.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口(api分组展示使用的)"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空,若为空，则默认按照所有id排序，不涉及社区",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "score",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._ResponsePostList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "作者id",
                    "type": "integer"
                },
                "author_name": {
                    "description": "作者",
                    "type": "string"
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "description": "omitempty表示当该字段为空时就不显示了",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                },
                "vote_num": {
                    "description": "投票数",
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "bluebell项目接口文档",
	Description: "Go web开发进阶项目实战课程bluebell",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
