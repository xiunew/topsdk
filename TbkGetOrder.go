package topsdk

//TbkItemInfoGetOrder taobao.tbk.order.get (淘宝客订单)
//https://open.taobao.com/api.htm?docId=24527&docType=2&scopeId=11650
type TbkItemInfoGetOrder struct {
	Fields string `json:"fields"`
	//true	tb_trade_parent_id,tb_trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk3rd_pub_id,tk3rd_site_id,tk3rd_adzone_id,relation_id,tb_trade_parent_id,tb_trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk3rd_pub_id,tk3rd_site_id,tk3rd_adzone_id,special_id,click_time 需返回的字段列表
	Start_time string `json:"start_time"`
	//true  2016-05-23 12:18:22 订单查询开始时间
	Span int `json:"span"`
	//false 600 订单查询时间范围，单位：秒，最小60，最大1200，如不填写，默认60。查询常规订单、三方订单、渠道，及会员订单时均需要设置此参数，直接通过设置page_size,page_no 翻页查询数据即可
	Page_no int `json:"page_no"`
	//false	1 第几页，默认1，1~100
	Page_size int `json:"page_size"`
	//false	20 页大小，默认20，1~100
	Tk_status int `json:"tk_status"`
	//false  订单状态，1: 全部订单，3：订单结算，12：订单付款， 13：订单失效，14：订单成功； 订单查询类型为‘结算时间’时，只能查订单结算状态
	Order_query_type string `json:"order_query_type"`
	//false	 settle_time 订单查询类型，创建时间“create_time”，或结算时间“settle_time”。当查询渠道或会员运营订单时，建议入参创建时间“create_time”进行查询
	Order_scene int `json:"order_scene"`
	//false  1订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1，通过设置订单场景类型，媒体可以查询指定场景下的订单信息，例如不设置，或者设置为1，表示查询常规订单，常规订单包含淘宝客所有的订单数据，含渠道，及会员运营订单，但不包含3方分成，及维权订单
	Order_count_type int `json:"order_count_type"`
	//false  1订单数据统计类型，1: 2方订单，2: 3方订单，如果不设置，或者设置为1，表示2方订单
}

//APIName API名称
func (m TbkItemInfoGetOrder) APIName() string {
	return "taobao.tbk.order.get"
}

//Params 参数
func (m TbkItemInfoGetOrder) Params() map[string]string {
	var params = make(map[string]string)
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkItemInfoGetOrder) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkItemInfoGetOrder) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkItemInfoGetOrder) AddParam(key string, value interface{}) {

}
