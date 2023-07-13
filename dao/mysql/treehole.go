package mysql

import "web_app/model"

func CreateTreeHole(treehole *model.TreeHole) (err error) {
	err = DB.Create(treehole).Error
	return
}
func GetTreeHole() (treehole []model.TreeHole, err error) {
	err = DB.Find(&treehole).Error
	return treehole, err
}
