package assert

// errorオブジェクトをチェックし、nilの場合例外を送出
func Assert(err error, msg string) {
	if err != nil {
		panic(err.Error() + ":" + msg)
	}
}
