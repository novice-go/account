package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"wdkj/account/model"
)

func TestInitConfig(t *testing.T) {

	resp := &model.Config{}
	err := InitConfig(os.Getenv("GOPATH")+"/src/wdkj/account/config.yaml", resp)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.DBConfig.Host)
	t.Logf("%+v", resp)
}
