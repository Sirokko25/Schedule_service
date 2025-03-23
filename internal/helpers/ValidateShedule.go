package helpers

import "sheduler/internal/models"

func ValidateShedule(model models.Shedule) bool {
	return !(model.UserId == 0 || model.Medicine == "" || model.Duration == "" || model.Period == 0)
}
