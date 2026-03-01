package routes

import (
	"fmt"
	domain_stream "go-backend-stream/internal/domains/stream"
	domain_upload "go-backend-stream/internal/domains/upload"
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/infrastructure/factory"

	"go-backend-stream/internal/utilities/logger"
	"go-backend-stream/internal/utilities/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, db *database.StreamDB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", middlewares.NewJwtSessionMiddleware(db), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1 := app.Group("v1")

	controllers := factory.NewAppControllers(db)

	//====================================================================================
	// (v1) Auth
	//====================================================================================
	{
		v1Auth := v1.Group("auth")
		// v1Auth.Post("/sign-in", authHandler.SignIn)
		// v1Auth.Post("/sign-up", authHandler.SinUp)

		//====================================================================================
		// (v1) Session Auth
		//====================================================================================
		v1SessionAuth := v1Auth.Group("/session")
		v1SessionAuth.Post("/sign-in", controllers.AuthController.SignInWithSession)
		v1SessionAuth.Post("/sign-up", controllers.AuthController.SignUpWithSessionController)
		v1SessionAuth.Get("/", middlewares.NewJwtSessionMiddleware(db), func(c *fiber.Ctx) error {
			userId := c.Locals("userId").(string)

			return c.JSON(fiber.Map{
				"message": "sessão válida",
				"user_id": userId,
			})
		})
		v1SessionAuth.Get("/all", middlewares.NewJwtSessionMiddleware(db), controllers.AuthController.GetAllSessionsController)
		v1SessionAuth.Delete("/:sessionId", middlewares.NewJwtSessionMiddleware(db), controllers.AuthController.DeleteSessionController)

		//====================================================================================
		// (v1) Recovery Auth
		//====================================================================================
		v1AuthRecovery := v1Auth.Group("/recovery")
		v1AuthRecovery.Post("/request/:email", controllers.AuthController.RequestRecoveryController)
		v1AuthRecovery.Post("/change-password", controllers.AuthController.ChangePasswordController)

		//====================================================================================
		// (v1) Recovery Auth
		//====================================================================================
		{
			v1Upload := v1.Group("upload")
			v1Upload.Post("/video", domain_upload.VideoUploadHandler)
		}

		{
			v1Stream := v1.Group("stream")
			v1Stream.Get("/videos/:filename", domain_stream.VideoStream)
		}

		{
			v1Organization := v1.Group("organization")
			v1Organization.Post("/", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.CreateOrganizationController)
			v1Organization.Get("/", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.GetUserOrganizationsController)
			v1Organization.Get("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.GetOrganizationController)
			v1Organization.Patch("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.UpdateOrganizationController)
			v1Organization.Delete("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.DeleteOrganizationController)
			v1Organization.Post("/:id/members", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.AddMemberController)
			v1Organization.Delete("/:id/members/:memberId", middlewares.NewJwtSessionMiddleware(db), controllers.OrganizationController.RemoveMemberController)
		}

		//====================================================================================
		// (v1) Game
		//====================================================================================
		fmt.Println("teste")
		{
			v1Game := v1.Group("game")
			v1Game.Post("/", middlewares.NewJwtSessionMiddleware(db), controllers.GameController.CreateGameController)
			v1Game.Get("/", controllers.GameController.GetAllGamesController)
			v1Game.Get("/:id", controllers.GameController.GetGameController)
			v1Game.Patch("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.GameController.UpdateGameController)
			v1Game.Delete("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.GameController.DeleteGameController)
		}

		//====================================================================================
		// (v1) Member
		//====================================================================================
		{
			v1Member := v1.Group("member")
			v1Member.Get("/", middlewares.NewJwtSessionMiddleware(db), controllers.MemberController.GetUserMembersController)
			v1Member.Get("/:id", middlewares.NewJwtSessionMiddleware(db), controllers.MemberController.GetMemberController)
			v1Member.Post("/:id/games", middlewares.NewJwtSessionMiddleware(db), controllers.MemberController.AddGameToMemberController)
			v1Member.Delete("/:id/games/:gameId", middlewares.NewJwtSessionMiddleware(db), controllers.MemberController.RemoveGameFromMemberController)
		}

		logger.Info.Printf("log all routes mapped...")
		for _, r := range app.GetRoutes() {
			if r.Method == "HEAD" || r.Method == "OPTIONS" || r.Method == "CONNECT" || r.Method == "TRACE" {
				continue
			}

			logger.Debug.Printf("[ROUTE] %s %s\n", r.Method, r.Path)

		}
	}

}
