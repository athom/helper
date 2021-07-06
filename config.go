package helper

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func ReadYamlAsViperConfig(basename string) (r *viper.Viper) {
	rootDir := GetEnv("ROOT_DIR", ".")
	r = viper.New()
	r.SetConfigType("yaml")
	r.SetConfigName("config")
	r.AddConfigPath(fmt.Sprintf("%v/"+basename, rootDir))
	err := r.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return r
}

type SiteVerificationItem struct {
	Name    string `mapstructure:"name"`
	Content string `mapstructure:"content"`
}

func GetSiteVerifications(config *viper.Viper, key string) (r []*SiteVerificationItem, err error) {
	err = config.UnmarshalKey(key, &r)
	if err != nil {
		return
	}

	return
}

func GetSiteVerificationsTemplate(config *viper.Viper, key string) (r template.HTML, err error) {
	items, err := GetSiteVerifications(config, key)
	if err != nil {
		return
	}

	var lines []string
	for _, sv := range items {
		lines = append(lines, fmt.Sprintf(`<meta name="%v" content="%v" />`, sv.Name, sv.Content))
	}
	s := strings.Join(lines, " ")
	r = template.HTML(s)
	return
}
