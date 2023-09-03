package handlers

import (
	"time"

	"github.com/amezianechayer/ssltracker/data"
	"github.com/gofiber/fiber/v2"
)

func HandleGetDashboard(c *fiber.Ctx) error {
	sslTrackings := []data.DomainTracking{
		{
			ID:         1,
			Expires:    time.Now().AddDate(0, 0, 4),
			Issuer:     "Let's encrypt",
			DomainName: "sendit.sh",
			Statut:     "OK",
		},
		{
			ID:         2,
			Expires:    time.Now().AddDate(0, 0, 20),
			Issuer:     "Let's encrypt",
			DomainName: "fulltimegodev.com",
			Statut:     "OK",
		},
		{
			ID:         3,
			Expires:    time.Now().AddDate(0, 1, 20),
			Issuer:     "Let's encrypt",
			DomainName: "thetotalcoder.com",
			Statut:     "OK",
		},
	}
	data := fiber.Map{
		"trackings": sslTrackings,
	}
	return c.Render("dashboard/index", data)
}
