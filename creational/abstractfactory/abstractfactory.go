// 抽象工厂模式

package abstractfactory

import "fmt"

// SaveArticle 抽象模式工厂接口
type SaveArticle interface {
	CreateProse() Prose
	CreateAncientPoetry() AncientPoetry
}

type SaveRedis struct{}

func (*SaveRedis) CreateProse() Prose {
	return &RedisProse{}
}

func (*SaveRedis) CreateAncientPoetry() AncientPoetry {
	return &RedisProse{}
}

type SaveMysql struct{}

func (*SaveMysql) CreateProse() Prose {
	return &MysqlProse{}
}

func (*SaveMysql) CreateAncientPoetry() AncientPoetry {
	return &MysqlProse{}
}

// Prose 散文
type Prose interface {
	SaveProse()
}

// AncientPoetry 古诗
type AncientPoetry interface {
	SaveAncientPoetry()
}

type RedisProse struct{}

func (*RedisProse) SaveProse() {
	fmt.Println("Redis Save Prose")
}

func (*RedisProse) SaveAncientPoetry() {
	fmt.Println("Redis Save Ancient Poetry")
}

type MysqlProse struct{}

func (*MysqlProse) SaveProse() {
	fmt.Println("Mysql Save Prose")
}

func (*MysqlProse) SaveAncientPoetry() {
	fmt.Println("Mysql Save Ancient Poetry")
}

func CreateAndSave(saveArticle SaveArticle) {
	saveArticle.CreateProse().SaveProse()
	saveArticle.CreateAncientPoetry().SaveAncientPoetry()
}
