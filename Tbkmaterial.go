package topsdk

type TbkItemInfoMaterial struct {
	//TbkItemInfoMaterial taobao.tbk.sc.material.optional( 通用物料搜索API )
	//https://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991
	Q          string `json:"q"`          // 关键字
	Page_no    int64  `json:"page_no"`    //	false	1	第几页，默认：１
	Page_size  int64  `json:"page_size"`  //	false	20	页大小，默认20，1~100
	Sort       string `json:"sort"`       //tk_rate_des 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
	Adzone_id  int64  `json:"adzone_id"`  //true 12345678 mm_xxx_xxx_12345678三段式的最后一段数字
	Site_id    int64  `json:"site_id"`    //true 22 mm_xxx_22_xxx三段式的第二段数字
	Has_coupon bool   `json:"has_coupon"` //	false	false	优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	Session    string `json:"session"`    //用户登录授权成功后，TOP颁发给应用的授权信息，详细介绍请点击这里。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
}

//APIName API名称
func (m TbkItemInfoMaterial) APIName() string {
	return "taobao.tbk.sc.material.optional"
}

//Params 参数
func (m TbkItemInfoMaterial) Params() map[string]string {
	var params = make(map[string]string)


	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkItemInfoMaterial) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkItemInfoMaterial) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkItemInfoMaterial) AddParam(key string, value interface{}) {

}
