package groups

import (
	"asira_borrower/handlers"
	"asira_borrower/middlewares"

	"github.com/labstack/echo"
)

func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	middlewares.SetClientJWTmiddlewares(g, "client")
	g.GET("/check_unique", handlers.CheckData)
	g.POST("/register_borrower", handlers.RegisterBorrower)
	g.POST("/borrower_login", handlers.BorrowerLogin)

	g.POST("/agent_login", handlers.AgentLogin)

	g.POST("/reset_password", handlers.ClientResetPassword)
	g.POST("/change_password", handlers.ChangePassword)

	//banks
	g.GET("/banks", handlers.ClientBanks)
	g.GET("/banks/:bank_id", handlers.ClientBankbyID)

	//bank service
	g.GET("/bank_services", handlers.ClientBankServices)
	g.GET("/bank_services/:id", handlers.ClientBankServicebyID)

	// loan purposes
	g.GET("/loan_purposes", handlers.LoanPurposeList)
	g.GET("/loan_purposes/:loan_purpose_id", handlers.LoanPurposeDetail)

	//server time & service info
	g.GET("/serviceinfo", handlers.ServiceInfo)

	g.POST("/otp_request", handlers.RequestOTPverifyAccount)

	//FAQ
	g.GET("/faq", handlers.FAQList)
	g.GET("/faq/:faq_id", handlers.FAQDetail)
}
