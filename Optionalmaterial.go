package topsdk

import (
	"fmt"
	"strconv"
)

//物料搜索
type Optionalmaterial struct {
	//Optionalmaterial    taobao.tbk.dg.material.optional ( 淘宝客-推广者-物料搜索 )
	//https://open.taobao.com/api.htm?docId=35896&docType=2
	//PasswordContent string `json:"password_content"`
	Q          string `json:"q"`          // 关键字
	Page_no    int64  `json:"page_no"`    //	false	1	第几页，默认：１
	Page_size  int64  `json:"page_size"`  //	false	20	页大小，默认20，1~100
	Platform   int64  `json:"platform"`   //链接形式：1：PC，2：无线，默认：１
	Sort       string `json:"sort"`       //排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
	Has_coupon bool   `json:"has_coupon"` //	false	false	优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	Adzone_id  int64  `json:"adzone_id"`  //true 12345678 mm_xxx_xxx_12345678三段式的最后一段数字

}

//APIName API名称
func (m Optionalmaterial) APIName() string {
	return "taobao.tbk.dg.material.optional"
}

//Params 参数
func (m Optionalmaterial) Params() map[string]string {
	var params = make(map[string]string)
	params["q"] = m.Q
	params["page_no"] = strconv.FormatInt(m.Page_no, 10)
	params["page_size"] = strconv.FormatInt(m.Page_size, 10)
	params["platform"] = strconv.FormatInt(m.Platform, 10)
	params["sort"] = m.Sort
	fmt.Println(params)
	if m.Has_coupon == true{
		params["has_coupon"] = "true"
	}else{
		params["has_coupon"] = "false"
	}
	params["adzone_id"] = strconv.FormatInt(m.Adzone_id, 10)
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m Optionalmaterial) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m Optionalmaterial) ExtJSONParamValue() string {
	return marshal(m)
}
