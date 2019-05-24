package Core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


/**
	初始配置文件文件，如果需要多个可以自定义添加
 */
type Config struct {
	AppName string `yaml:"appname"`//项目名称
	Port string `yaml:"port"`//端口
	DbName string `yaml:"dbname"`//数据库名称
	DbPwd string `yaml:"dbpwd"`//密码
	DeBug bool `yaml:"debug"` //调试模式
	DbHost string `yaml:"dbhost"`//数据库地址
	DbType string `yaml:"dbtype"`//数据库类型
}

func NewConfig() *Config {

	var con Config

	return 	con.getConf()
}

/**
	返回配置文件
 */
func (this *Config) getConf() *Config  {

	yamFile,err := ioutil.ReadFile("Conf/app")

	if err != nil{

		panic("配置文件地址不存在，请检查配置文件Conf/app 文件是否存在")
	}

	err = yaml.Unmarshal(yamFile,this)

	if err != nil{

		panic(err)
	}

	return this

}