/*
 * @Author: Jecosine
 * @Date: 2021-01-04 02:45:03
 * @LastEditTime: 2021-01-05 00:49:59
 * @LastEditors: Jecosine
 * @Description: model of category
 */
package models

type Category struct {
	Model

	ID       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"column:username" json:"name"`
	ParentID string `json:"parent_id"`
}

func GetCategory(id string) (category Category) {

}

func GetCategoryByName(name string) (category Category) {

}

func EditCategory(id string) bool {

}

func 