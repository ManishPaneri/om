package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id             int    `orm:"column(id);auto" json:"id"`
	Name           string `orm:"column(name);size(45)" json:"name"`
	MinAmount      int    `orm:"column(minAmount);null" json:"minAmount"`
	MaxAmount      int    `orm:"column(maxAmount);null" json:"maxAmount"`
	PurchaseAmount int    `orm:"column(purchaseAmount);null" json:"purchaseAmount"`
	Item           int    `orm:"column(item);null" json:"item"`
	Enable         int    `orm:"column(enable);null" json:"enable"`
}

func (t *Product) TableName() string {
	return "product"
}

func init() {
	orm.RegisterModel(new(Product))
}

// AddProduct insert a new Product into database and returns
// last inserted Id on success.
func AddProduct(m *Product) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductById retrieves Product by Id. Returns error if
// Id doesn't exist
func GetProductById(id int) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{}
	err = o.QueryTable(new(Product)).Filter("Id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllProduct() (v []Product, err error) {
	o := orm.NewOrm()
	v = []Product{}
	_, err = o.QueryTable(new(Product)).RelatedSel().All(&v)
	fmt.Println(v)
	return v, err
}

// UpdateProduct updates Product by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductById(m *Product) (err error) {
	o := orm.NewOrm()
	v := &Product{}
	err = o.QueryTable(new(Product)).Filter("Id", m.Id).RelatedSel().One(v)
	// ascertain id exists in the database
	if err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProduct deletes Product by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProduct(id int) (err error) {
	o := orm.NewOrm()
	v := Product{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Product{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
