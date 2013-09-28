func FindUser(c mgo.Collection, username string) result, err{
    var result User
    err := c.Find(bson.M{"username", username}).One(&result)
    return result, err
}

func main(){
    err := FindUser("JohnDoe")
    if err != nil{
        //Handle error
        //Log error message, repeat action, etc.
    }
}
