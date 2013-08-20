package main


// START OMIT
select {
    case <-algorithm1:
        // a read from ch has occurred
    case <-algorithm2:
        // the read from ch has timed out
}
// END OMIT
