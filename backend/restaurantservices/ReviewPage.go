package restaurantservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

type Reviews struct {
	Reviews []model.Review `json:"reviews"`
}

func GetReviews(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := model.GetRestaurantProfilesFromFile()
	var rev []model.Review
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.RID {
			rev = profiles.Profiles[i].Reviews
			break
		}
	}
	return c.JSON(http.StatusOK, Reviews{
		Reviews: rev,
	})
}

func PostReply(c echo.Context) error { //needs restaurant_id, review_id, reply
	var rev model.Review
	if err := c.Bind(&rev); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == rev.RestaurantID {
			reviews := profiles.Profiles[i].Reviews
			for j := 0; j < len(reviews); j++ {
				if reviews[j].ReviewID == rev.ReviewID {
					reviews[j].Reply = rev.Reply
					rev = reviews[j]
					break
				}
				if j == len(reviews)-1 {
					return c.JSON(http.StatusConflict, model.ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "Wrong review ID.",
					})
				}
			}
			break
		}
		if i == len(profiles.Profiles)-1 {
			return c.JSON(http.StatusConflict, model.ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong restaurant ID.",
			})
		}
	}
	model.UpdateRestaurantProfileFile(profiles)
	return c.JSON(http.StatusOK, rev)
}
