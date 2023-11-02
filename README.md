
create user.proto file:   go to proto foler and run this command
protoc --go-grpc_out=. --go_out=. *.proto



start server:
go to server folder and run 
go run server.go


start client
run ----    go run client.go



open postman for API's run
localhost:8000/user          using get --   this will give all users list.




localhost:8000/user          using post and send user in data will create new user.

data should be like this.
 {
            "Id": 7,
            "FName": "xyz",
            "City":"LA",
            "Phone":"454545",
            "Height":6.4,
            "Married":false
}

validation: Id should be greate than 0.
and Id, fname, city, phone, height is required.




localhost:8000/user/1        using get will give user which have id 1



validation:  i have applied validation on model of user.

