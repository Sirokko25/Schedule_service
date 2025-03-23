package models

type Shedule struct {
	UserId   int64  `json:"userid" binding:"required"`
	Medicine string `json:"medicine" binding:"required"`
	Period   int64  `json:"period" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

type SheduleWithChart struct {
	UserId   int64  `json:"userid"`
	Medicine string `json:"medicine"`
	Period   int64  `json:"period"`
	Duration string `json:"duration"`
	Chart    JSONB  `json:"chart"`
}

type Medicine struct {
	Medicine string   `json:"medicine"`
	Time     []string `json:"time"`
}

type JSONB struct {
	Time []string `json:"time"`
}
