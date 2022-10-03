package locations

type Location struct {
	GID          uint64        `json:"gid"`
	ID           uint64        `json:"id"`
	GUID         uint64        `json:"guid"`
	LocationCode string        `json:"locationCode"`
	Name         string        `json:"name"`
	Flag         string        `json:"flag"`
	Meta         []interface{} `json:"_meta"`
	CKey         string        `json:"_ckey"`
}

type LocationsList []Location
