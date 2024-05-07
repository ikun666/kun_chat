package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetMsg(w http.ResponseWriter, r *http.Request) {
	//1.  获取参数 并 检验 token 等合法性

	//uid, tid, start, end int64, isRev bool
	uid, err := strconv.Atoi(r.PostFormValue("uid"))
	if err != nil {
		fmt.Println("类型转换失败", err)
		return
	}
	tid, err := strconv.Atoi(r.PostFormValue("tid"))
	if err != nil {
		fmt.Println("类型转换失败", err)
		return
	}
	start, err := strconv.Atoi(r.PostFormValue("start"))
	if err != nil {
		fmt.Println("类型转换失败", err)
		return
	}
	end, err := strconv.Atoi(r.PostFormValue("end"))
	if err != nil {
		fmt.Println("类型转换失败", err)
		return
	}
	isRev, err := strconv.ParseBool(r.PostFormValue("isRev"))
	if err != nil {
		fmt.Println("类型转换失败", err)
		return
	}
	res := RedisMsg(int64(uid), int64(tid), int64(start), int64(end), isRev)
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(msg)
}
