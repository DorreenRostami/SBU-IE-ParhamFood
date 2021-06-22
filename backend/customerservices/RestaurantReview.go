package customerservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurantservices"
	"github.com/labstack/echo/v4"
)

type CustomerReviewReq struct {
	CustomerID   int    `json:"customer_id"`
	RestaurantID int    `json:"restaurant_id"`
	Text         string `json:"text"`
	Stars        int    `json:"stars"`
}

type RestReviewInfo struct {
	RID     int            `json:"id"`
	Reviews []model.Review `json:"reviews"`
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

func PostReview(c echo.Context) error {
	var req CustomerReviewReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var rev model.Review
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == req.RestaurantID {
			rev = model.Review{
				ReviewID:     len(profiles.Profiles[i].Reviews),
				CustomerID:   req.CustomerID,
				RestaurantID: req.RestaurantID,
				Text:         req.Text,
				Stars:        req.Stars,
				Reply:        "",
			}
			profiles.Profiles[i].Reviews = append(profiles.Profiles[i].Reviews, rev)
			model.UpdateRestaurantProfileFile(profiles)
			break
		}
	}
	return c.JSON(http.StatusOK, rev)
}
