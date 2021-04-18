package helper

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadYamlAsViperConfig(basename string) (r *viper.Viper) {
	rootDir := GetEnv("ROOT_DIR", ".")
	r = viper.New()
	r.SetConfigType("yaml")
	r.SetConfigName(basename)
	r.AddConfigPath(fmt.Sprintf("%v/"+basename, rootDir))
	err := r.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return r
}
