package handlers

import (
	"log"
	"strconv"
	"github.com/davidclb/Yoshu/errors"
	"github.com/gofiber/fiber/v2"
)




func (h *Handler) Accueil(q *fiber.Ctx) error{

	paramUid := q.Params("uid")
	uid, err := strconv.Atoi(paramUid)

	if err != nil {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", q)
		err := apperrors.NewInternal()
		q.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
	}
	u, err:= h.UserService.GetUser(q, uid)
	 if err != nil {
        log.Printf("Unable to find user: %v\n%v", uid, err)
        err := apperrors.NewNotFound("user", "uid")
        q.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
			}

		q.JSON(fiber.Map{

			"user":u,
			  })
		return nil

}
