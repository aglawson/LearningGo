package api

func GetRPC(network string) string {
	url := `http://` + network + `.infura.io/v3/` + Conf.INFURA_KEY
	return url
}
