package routes

import (
	"net/http"
	"strings"

	"github.com/ImPedro29/rinha-backend-2024/api/controllers"
	"github.com/ImPedro29/rinha-backend-2024/db/interfaces"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/valyala/fasthttp"
)

type Routes struct {
	pb.BankServiceServer
	DB interfaces.DB
}

var controller = controllers.NewController()

func InitRoutes(ctx *fasthttp.RequestCtx) {
	pathSplit := strings.Split(string(ctx.Path()), "/")
	if (len(pathSplit) != 5 && len(pathSplit) != 4) || pathSplit[1] != "clientes" {
		ctx.Response.Header.SetStatusCode(http.StatusNotFound)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")

	if pathSplit[3] == "transacoes" && ctx.IsPost() {
		controller.CreateTransaction(ctx, pathSplit[2])
		return
	}

	if pathSplit[3] == "extrato" && ctx.IsGet() {
		controller.Statements(ctx, pathSplit[2])
		return
	}

	ctx.Response.Header.SetStatusCode(http.StatusNotFound)
}
