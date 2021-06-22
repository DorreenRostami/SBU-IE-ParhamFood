package customerservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurantservices"
	"github.com/labstack/echo/v4"
)

type RestReviewInfo struct {
	RID     int            `json:"id"`
	Reviews []model.Review `json:"review"`
}

func GetRestaurantReviews(c echo.Context) error {
	var req rst.RestID
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var restInfo RestReviewInfo
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == req.RID {
			restInfo = RestReviewInfo{
				RID:     profiles.Profiles[i].ID,
				Reviews: profiles.Profiles[i].Reviews,
			}
			break
		}
	}
	return c.JSON(http.StatusOK, restInfo)
}
