package groups

import (
	"asira_borrower/handlers"
	"asira_borrower/middlewares"

	"github.com/labstack/echo"
)

func AgentGroup(e *echo.Echo) {
	g := e.Group("/agent")
	middlewares.SetClientJWTmiddlewares(g, "agent")

	// agent's profile endpoints
	g.GET("/profile", handlers.AgentProfile)
	g.PATCH("/profile", handlers.AgentProfileEdit)

	// agent's profile endpoints
	g.POST("/register_borrower", handlers.AgentRegisterBorrower)

	//banks owned by current agent (jti)
	g.GET("/banks", handlers.AgentAllBank)

	// agent's bank Endpoint
	g.GET("/bank_services", handlers.AgentBankService)

	// agent's bank Endpoint
	g.GET("/bank_products", handlers.AgentBankProduct)

	//borrowers owned by current agent (jti) and bank_id
	g.GET("/borrowers/:bank_id", handlers.AgentAllBorrower)

	//check borrower from agent is exist or not
	g.POST("/checks_borrower", handlers.AgentCheckBorrower)

}
