package opt

type LeaderOptions struct {
	Conf               string `flag:"conf"`
	SyncQueueCacheChan int    `flag:"sync_queue_cache_chan"`
	TimeTickerInterval int    `flag:"time_ticker_interval"`
}

func NewLeaderOptions() *LeaderOptions {
	return &LeaderOptions{
		Conf:               "conf.yaml",
		SyncQueueCacheChan: 10000,
	}
}
