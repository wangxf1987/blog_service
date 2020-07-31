package errcode

var ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
var ErrorCreateTagFail = NewError(20010002, "创建标签失败")
var ErrorUpdateTagFail = NewError(20010003, "更新标签失败")
var ErrorDeleteTagFail = NewError(20010004, "删除标签失败")
var ErrorCountTagFail = NewError(20010005, "统计标签失败")
