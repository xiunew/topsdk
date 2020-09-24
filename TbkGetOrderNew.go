package topsdk

import (
	"strconv"
	"time"
)
//TbkItemInfoGetOrderNew taobao.tbk.order.details.get (淘宝客订单(新版))
//https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
type TbkItemInfoGetOrderNew struct {
	Query_type int `json:"query_type"`
	//查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询
	Position_index string `json:"position_index"`
	//位点，除第一页之外，都需要传递；前端原样返回。
	Page_size int `json:"page_size"`
	//页大小，默认20，1~100
	Member_type int `json:"member_type"`
	//推广者角色类型,2:二方，3:三方，不传，表示所有角色
	Tk_status int `json:"tk_status"`
	//淘客订单状态，12-付款，13-关闭，14-确认收货，15-结算成功;不传，表示所有状态
	End_time string `json:"end_time"`
	//必选   订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	Start_time string `json:"start_time"`
	//必选  订单查询开始时间
	Jump_type int `json:"jump_type"`
	//跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	Page_no int `json:"page_no"`
	//第几页，默认1，1~100
	Order_scene int64 `json:"order_scene"`
	//场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
}

//APIName API名称
func (m TbkItemInfoGetOrderNew) APIName() string {
	return "taobao.tbk.order.details.get"
}

//Params 参数
func (m TbkItemInfoGetOrderNew) Params() map[string]string {
	var params = make(map[string]string)
	if m.Query_type == 0 {
		params["query_type"] = "2"
	}else{
		params["query_type"] = strconv.Itoa(m.Query_type)
	}
	if m.Position_index!=""{
		params["position_index"]=m.Position_index
	}
	if m.Page_size == 0 {
		params["page_size"]="100"
	}else{
		params["page_size"]=strconv.Itoa(m.Page_size)
	}
	if m.Member_type!=0{
		params["member_type"]=strconv.Itoa(m.Member_type)
	}
	if m.Tk_status!=0{
		params["tk_status"]=strconv.Itoa(m.Tk_status)
	}
	if m.End_time==""{
		params["end_time"]=time.Now().Format("2006-01-02 15:04:05")
	}else{
		params["end_time"]=m.End_time
	}
	if m.Start_time==""{
		params["start_time"]=time.Now().Format("2006-01-02 15:04:05")
	}else{
		params["start_time"]=m.Start_time
	}
	if m.Jump_type==0{
		params["jump_type"]="1"
	}else{
		params["jump_type"]=strconv.Itoa(m.Jump_type)
	}
	if m.Page_no==0{
		params["page_no"]="1"
	}else{
		params["page_no"]=strconv.Itoa(m.Page_no)
	}
	if m.Order_scene!=0{
		params["order_scene"]=strconv.FormatInt(m.Order_scene,10)
	}
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkItemInfoGetOrderNew) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkItemInfoGetOrderNew) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkItemInfoGetOrderNew) AddParam(key string, value interface{}) {

}
