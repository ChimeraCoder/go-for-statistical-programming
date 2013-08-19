//Execute a query that will automatically be throttled
func throttledQuery(queryQueue chan queryChan) {
    for q := range queryQueue {

endpoint_path := q.endpoint_path
                   method := q.method
                   keyVals := q.keyVals
                   response_ch := q.response_ch
                   result, err := Query(endpoint_path, method, keyVals)
                   response_ch <- struct {
                       result []byte
                           err    error
                   }{result, err}

               time.Sleep(SECONDS_PER_QUERY)
    }
}
