package structs

import "../logUtil"

type CoreCfg struct {
	token string `json:"token"`

	Logger  logUtil.Level
	LogFile string `json:"logFile"`
}

func (c *CoreCfg) SetToken(token string) {
	c.token = token
}
