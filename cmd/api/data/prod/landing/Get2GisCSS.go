package site

import "FitroomTests_v2/cmd/api/models"

var Get2GisMap = models.Request{
	URL:    "https://maps.api.2gis.ru/2.0/css/?version=v3.7.3",
	Method: "GET",
	Auth:   new(string),
}

var Get2GisCatalog = models.Request{
	URL:    "https://catalog.api.2gis.ru/2.0/region/list?format=json&key=rubnkm7490&fields=items.bounds,items.zoom_level,items.time_zone,items.code,items.flags,items.country_code,items.domain,items.default_pos",
	Method: "GET",
	Auth:   new(string),
}
