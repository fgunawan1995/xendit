package model

var defaultInternalServerError = "Internal Server Error"
var asIs = "AS_IS"
var specialErrors = map[string]string{
	"net/http: request canceled":                    "RTO",
	"marvelStatus:404":                              "Character not found",
	"invalid syntax":                                "characterId must be number",
	"currently fetching all marvel characters data": asIs,
}
