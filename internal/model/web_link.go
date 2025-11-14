package model

type CheckLinksStatusByUrlRequest struct {
	Links []string `json:"links" validate:"required,gt=0,lte=200,dive,gt=0,lte=200"`
}

type CheckLinksStatusByUrlResponse struct {
	Links    map[string]string `json:"links"`
	LinksNum int               `json:"links_num"`
}

type CheckLinksStatusByIDRequest struct {
	LinksList []int `json:"links_list" validate:"required,gt=0,lte=200,dive,gt=0,lte=1000000000"`
}

type PersistentStorageData struct {
	ID          int      `json:"id"`
	LinkedLinks []string `json:"links"`
}
