package main


// START OMIT
select {
    case <-ch:
        // a read from ch has occurred
    case <-timeout:
        // the read from ch has timed out
}
// END OMIT
