package e

var MsgFlags = map[int]string {
	SUCCESS : "OK",
	ERROR : "fail",
	INVALID_PARAMS : "wrong parameters",
	ERROR_EXIST_TAG : "tag already exists",
	ERROR_NOT_EXIST_TAG : "tag don't exists",
	ERROR_NOT_EXIST_ARTICLE : "article not exists",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "token authorized fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token timeout",
    ERROR_AUTH_TOKEN : "Token generates failed",
    ERROR_AUTH : "Token error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}