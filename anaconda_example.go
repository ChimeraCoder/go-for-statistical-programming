//throttledQuery executes and automatically throttles queries according to SECONDS_PER_QUERY
func (c TwitterApi) throttledQuery() {
    for q := range c.queryQueue {
        now := time.Now()

        err := c.execQuery(q.url, q.form, q.data, q.method)

        q.response_ch <- struct {
            data interface{}
            err  error
        }{q.data, err}

        time.Sleep(SECONDS_PER_QUERY - time.Since(now))
    }
}

