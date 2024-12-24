package etcd

func GlobalClient() *Client {
	return globalClient
}

func AddFunc(key string, initFunc InitFunc, watchFunc WatchFunc) {
	globalClient.AddFunc(key, initFunc, watchFunc)
}

func Start() {
	globalClient.Start()
}

func Restart() {
	globalClient.Restart()
}
