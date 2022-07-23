// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Repository",
            "url": "https://github.com/Scrip7/imdb-api"
        },
        "license": {
            "name": "MIT license",
            "url": "https://github.com/Scrip7/imdb-api/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/title/{id}": {
            "get": {
                "description": "The ID should start with \"tt\" at the beginning.\n| Title Text | IMDb ID |\n| --- | --- |\n| Spider-Man: No Way Home | ` + "`" + `tt10872600` + "`" + ` |\n| Spider-Man: Far from Home | ` + "`" + `tt6320628` + "`" + ` |\n| Spider-Man: Homecoming | ` + "`" + `tt2250912` + "`" + ` |\n| Avengers: Endgame | ` + "`" + `tt4154796` + "`" + ` |\n| Avengers: Infinity War | ` + "`" + `tt4154756` + "`" + ` |\n| The Dark Knight | ` + "`" + `tt0468569` + "`" + ` |\n| The Godfather | ` + "`" + `tt0068646` + "`" + ` |\n| Friends | ` + "`" + `tt0108778` + "`" + ` |\n| Breaking Bad | ` + "`" + `tt0903747` + "`" + ` |\n| Chernobyl | ` + "`" + `tt7366338` + "`" + ` |",
                "tags": [
                    "Title"
                ],
                "summary": "Get IMDb title by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pipe.IndexTransform"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pipe.Cast": {
            "type": "object",
            "properties": {
                "characters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "episodeCredits": {
                    "$ref": "#/definitions/pipe.CastEpisodeCredits"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "$ref": "#/definitions/pipe.CastImage"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "pipe.CastEpisodeCredits": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "years": {
                    "$ref": "#/definitions/pipe.YearRange"
                }
            }
        },
        "pipe.CastImage": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.ExternalReviews": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                }
            }
        },
        "pipe.FAQ": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "a short list of FAQs as preview",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.FAQItem"
                    }
                },
                "total": {
                    "description": "total number of keywords that are related to this title",
                    "type": "integer"
                }
            }
        },
        "pipe.FAQItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "pipe.FeaturedReviewItem": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/pipe.ReviewAuthor"
                },
                "createdAt": {
                    "type": "string"
                },
                "dislikes": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.Genre": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "slug": {
                    "description": "can directly be used in advances search feature",
                    "type": "string"
                }
            }
        },
        "pipe.ImageItem": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.Images": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.ImageItem"
                    }
                },
                "primary": {
                    "$ref": "#/definitions/pipe.PrimaryImage"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "pipe.IndexTransform": {
            "type": "object",
            "properties": {
                "cast": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.Cast"
                    }
                },
                "faq": {
                    "$ref": "#/definitions/pipe.FAQ"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.Genre"
                    }
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "$ref": "#/definitions/pipe.Images"
                },
                "keywords": {
                    "$ref": "#/definitions/pipe.Keywords"
                },
                "plot": {
                    "type": "string"
                },
                "popularity": {
                    "$ref": "#/definitions/pipe.Popularity"
                },
                "related": {
                    "description": "list of related titles",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.RelatedTitle"
                    }
                },
                "reviews": {
                    "$ref": "#/definitions/pipe.Reviews"
                },
                "series": {
                    "description": "only when viewing an episode of a series",
                    "$ref": "#/definitions/pipe.Series"
                },
                "soundtracks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.Soundtrack"
                    }
                },
                "title": {
                    "$ref": "#/definitions/pipe.Title"
                },
                "trivia": {
                    "$ref": "#/definitions/pipe.Trivia"
                },
                "validate": {
                    "$ref": "#/definitions/pipe.Validate"
                },
                "videos": {
                    "$ref": "#/definitions/pipe.Videos"
                }
            }
        },
        "pipe.Keywords": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "a short list of keywords as preview",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "total": {
                    "description": "total number of keywords that are related to this title",
                    "type": "integer"
                }
            }
        },
        "pipe.PlaybackItem": {
            "type": "object",
            "properties": {
                "mimeType": {
                    "type": "string"
                },
                "quality": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "pipe.Popularity": {
            "type": "object",
            "properties": {
                "difference": {
                    "type": "integer"
                },
                "rank": {
                    "type": "integer"
                }
            }
        },
        "pipe.PrimaryImage": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "thumbnails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.PrimaryImageThumbnail"
                    }
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.PrimaryImageThumbnail": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.Rating": {
            "type": "object",
            "properties": {
                "certificate": {
                    "description": "For example \"TV-14\"",
                    "type": "string"
                },
                "count": {
                    "description": "Voters count",
                    "type": "integer"
                },
                "score": {
                    "description": "IMDb score from 0 to 10 (includes precision)",
                    "type": "number"
                }
            }
        },
        "pipe.RelatedTitle": {
            "type": "object",
            "properties": {
                "canHaveEpisodes": {
                    "type": "boolean"
                },
                "duration": {
                    "description": "unit is seconds",
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.Genre"
                    }
                },
                "id": {
                    "type": "string"
                },
                "poster": {
                    "$ref": "#/definitions/pipe.RelatedTitlePoster"
                },
                "rating": {
                    "$ref": "#/definitions/pipe.Rating"
                },
                "releaseYear": {
                    "$ref": "#/definitions/pipe.YearRange"
                },
                "title": {
                    "$ref": "#/definitions/pipe.RelatedTitleName"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "pipe.RelatedTitleName": {
            "type": "object",
            "properties": {
                "original": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.RelatedTitlePoster": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.ReviewAuthor": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "pipe.Reviews": {
            "type": "object",
            "properties": {
                "external": {
                    "$ref": "#/definitions/pipe.ExternalReviews"
                },
                "featured": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.FeaturedReviewItem"
                    }
                },
                "users": {
                    "$ref": "#/definitions/pipe.UsersReviews"
                }
            }
        },
        "pipe.Series": {
            "type": "object",
            "properties": {
                "current": {
                    "$ref": "#/definitions/pipe.SeriesCurrent"
                },
                "id": {
                    "$ref": "#/definitions/pipe.SeriesID"
                },
                "releaseYear": {
                    "$ref": "#/definitions/pipe.YearRange"
                },
                "title": {
                    "$ref": "#/definitions/pipe.SeriesTitle"
                }
            }
        },
        "pipe.SeriesCurrent": {
            "type": "object",
            "properties": {
                "episode": {
                    "description": "episode number (starts from 1)",
                    "type": "integer"
                },
                "season": {
                    "description": "season number (starts from 1)",
                    "type": "integer"
                }
            }
        },
        "pipe.SeriesID": {
            "type": "object",
            "properties": {
                "episodeNext": {
                    "type": "string"
                },
                "episodePrevious": {
                    "type": "string"
                },
                "parent": {
                    "type": "string"
                }
            }
        },
        "pipe.SeriesTitle": {
            "type": "object",
            "properties": {
                "original": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.Soundtrack": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.SoundtrackComment"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "pipe.SoundtrackComment": {
            "type": "object",
            "properties": {
                "html": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.Thumbnail": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "pipe.Title": {
            "type": "object",
            "properties": {
                "aka": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "original": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.Trivia": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "pipe.UsersReviews": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                }
            }
        },
        "pipe.Validate": {
            "type": "object",
            "properties": {
                "canHaveEpisodes": {
                    "type": "boolean"
                },
                "isAdult": {
                    "type": "boolean"
                },
                "isEpisode": {
                    "type": "boolean"
                },
                "isMovie": {
                    "type": "boolean"
                },
                "isSeries": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "pipe.VideoItem": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "thumbnail": {
                    "$ref": "#/definitions/pipe.Thumbnail"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/pipe.VideoTypeWrapper"
                }
            }
        },
        "pipe.VideoItemPrimary": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "isMature": {
                    "type": "boolean"
                },
                "playback": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.PlaybackItem"
                    }
                },
                "thumbnail": {
                    "$ref": "#/definitions/pipe.Thumbnail"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/pipe.VideoTypeWrapper"
                }
            }
        },
        "pipe.VideoTypeWrapper": {
            "type": "object",
            "properties": {
                "slug": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "pipe.Videos": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.VideoItem"
                    }
                },
                "primaries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipe.VideoItemPrimary"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "pipe.YearRange": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                }
            }
        },
        "server.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "IMDb-API",
	Description:      "Cross-platform microservice to scrape the IMDb website.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}