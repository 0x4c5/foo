package vipper

import (
	"github.com/spf13/viper"
)

func Init(filename string) (vp *viper.Viper, err error) {
	// config init
	vp = viper.New()
	vp.SetConfigFile(filename)
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		return
	}
	return
}
