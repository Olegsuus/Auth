// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/google/callback": {
            "get": {
                "description": "Обрабатывает ответ от Google после аутентификации и генерирует JWT-токен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Аутентификация"
                ],
                "summary": "Обратный вызов после аутентификации через Google",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Код авторизации от Google",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Состояние запроса",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/auth/google/login": {
            "get": {
                "description": "Перенаправляет пользователя на страницу авторизации Google",
                "tags": [
                    "Аутентификация"
                ],
                "summary": "Вход через Google",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Состояние запроса",
                        "name": "state",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Код авторизации от Google",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirect to Google"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/chats": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создает новый чат с указанными участниками",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Чаты"
                ],
                "summary": "Создание нового чата",
                "parameters": [
                    {
                        "description": "ID участников чата",
                        "name": "chat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_handlers_dto.AddChatDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/chats/ws": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Устанавливает WebSocket соединение для обмена сообщениями в реальном времени",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Чаты"
                ],
                "summary": "Установление WebSocket соединения для чата",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID чата",
                        "name": "chat_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/chats/{:id}": {
            "get": {
                "description": "Возвращает информацию о пользователе по email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Получение информации о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id чата",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Проверяет учетные данные пользователя и возвращает JWT-токен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Аутентификация пользователя",
                "parameters": [
                    {
                        "description": "Учетные данные пользователя",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_handlers_dto.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/messages": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Отправляет сообщение в указанный чат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Сообщения"
                ],
                "summary": "Отправка сообщения в чат",
                "parameters": [
                    {
                        "description": "Данные сообщения",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_handlers_dto.SendMessageDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/messages/chat/{chat_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает все сообщения из указанного чата",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Сообщения"
                ],
                "summary": "Получение сообщений из чата",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID чата",
                        "name": "chat_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Создает нового пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Регистрация нового пользователя",
                "parameters": [
                    {
                        "description": "Данные нового пользователя",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_handlers_dto.RegisterNewUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/user/password/reset": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Позволяет пользователю сбросить пароль, используя секретное слово",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Сброс пароля пользователя",
                "parameters": [
                    {
                        "description": "Данные для сброса пароля",
                        "name": "reset",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_handlers_dto.ResetPasswordDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/user/profile": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Обновляет данные профиля пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Обновление профиля пользователя",
                "parameters": [
                    {
                        "description": "Данные для обновления профиля",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Olegsuus_GoChat_internal_models.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/user/{email}": {
            "get": {
                "description": "Возвращает информацию о пользователе по email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Получение информации о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удаляет пользователя по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователи"
                ],
                "summary": "Удаление пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Неверные данные запроса"
                    },
                    "500": {
                        "description": "Ошибка на сервере"
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Olegsuus_GoChat_internal_handlers_dto.AddChatDTO": {
            "type": "object",
            "required": [
                "participant_ids"
            ],
            "properties": {
                "participant_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "github_com_Olegsuus_GoChat_internal_handlers_dto.LoginDTO": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "github_com_Olegsuus_GoChat_internal_handlers_dto.RegisterNewUserDTO": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "github_com_Olegsuus_GoChat_internal_handlers_dto.ResetPasswordDTO": {
            "type": "object",
            "required": [
                "email",
                "new_password",
                "secret_word"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string",
                    "minLength": 6
                },
                "secret_word": {
                    "type": "string"
                }
            }
        },
        "github_com_Olegsuus_GoChat_internal_handlers_dto.SendMessageDTO": {
            "type": "object",
            "required": [
                "chat_id",
                "content"
            ],
            "properties": {
                "chat_id": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                }
            }
        },
        "github_com_Olegsuus_GoChat_internal_models.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8765",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Chat API",
	Description:      "API для аутентификации пользователей и онлайн-чата с WebSocket.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
