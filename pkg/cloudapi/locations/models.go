package locations

type Location struct {
	GID          int           `json:"gid"`
	Id           int           `json:"id"`
	Guid         int           `json:"guid"`
	LocationCode string        `json:"locationCode"`
	Name         string        `json:"name"`
	Flag         string        `json:"flag"`
	Meta         []interface{} `json:"_meta"`
	CKey         string        `json:"_ckey"`
}

type LocationsList []Location
