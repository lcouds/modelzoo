package consts

import "time"

const (
	LabelBuildName      = "ai.tensorchord.build"
	LabelName           = "ai.tensorchord.name"
	LabelServerResource = "ai.tensorchord.server-resource"

	Domain        = "modelzoo.live"
	DefaultPrefix = "modelzoo-"
	APIKEY_PREFIX = "mzi-"
)
const DefaultAPIServerReadyTimeout = 15 * time.Minute
