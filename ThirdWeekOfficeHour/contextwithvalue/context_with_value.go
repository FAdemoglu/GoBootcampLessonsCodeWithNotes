package main

import (
	"context"
	"fmt"
)

//context paketi bir yerden bir yere veri taşımanız

func main() {
	ProcessRequest("patika", "patika123")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "patikaID", userID) //context.Background empty bir context,diğer parametre key, diğeri ise value
	ctx = context.WithValue(ctx, "patikaToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"context içerisinde gelen veriler : %v (%v)",
		ctx.Value("patikaID"),
		ctx.Value("patikaToken"),
	)
}
