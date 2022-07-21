// +build !dev

package config

func GetTemplatePath(fileName string) string {
	return "/app/api/templates/" + fileName
}
