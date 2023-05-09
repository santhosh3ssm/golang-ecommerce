package controllers

import(


)
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")