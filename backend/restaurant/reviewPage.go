package restaurant

import (
	"net/http"

	fh "github.com/DorreenRostami/IE_ParhamFood/filehandler"
	"github.com/labstack/echo/v4"
)

func PostReply(c echo.Context) error { //needs restaurant_id, review_id, reply
	var rev fh.Review
	if err := c.Bind(&rev); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()
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
					return c.JSON(http.StatusConflict, ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "Wrong review ID.",
					})
				}
			}
			break
		}
		if i == len(profiles.Profiles)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong restaurant ID.",
			})
		}
	}
	fh.UpdateProfileFile(profiles)
	return c.JSON(http.StatusOK, rev)
}
