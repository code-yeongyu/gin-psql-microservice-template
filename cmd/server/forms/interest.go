package forms

import (
	"gopkg.in/guregu/null.v4"
)

// AddInterestRequest is a struct for adding interest request
type AddInterestRequest struct {
	ClubhouseUserID null.Int `json:"clubhouse_user_id" valid:"optional"`
	Email           string   `json:"email" valid:"email, required" binding:"required"`
}

// AddInterestResponse is a struct for adding interest response
type AddInterestResponse struct {
	ID              uint     `json:"id"`
	ClubhouseUserID null.Int `json:"clubhouse_user_id" valid:"optional"`
	Email           string   `json:"email" valid:"email, required" binding:"required"`
}

// AddInterestRequestDocument is the same as AddInterestRequest
// but its nullable types are replaced with golang's types
type AddInterestRequestDocument struct {
	ClubhouseUserID int64  `json:"clubhouse_user_id" valid:"optional" example:"1234"`
	Email           string `json:"email" valid:"email, required" binding:"required" example:"example@example.com"`
}

// AddInterestResponseDocument is the same as AddInterestResponse
// but its nullable types are replaced with golang's types
type AddInterestResponseDocument struct {
	ID              uint   `json:"id" example:"0"`
	ClubhouseUserID int64  `json:"clubhouse_user_id" valid:"optional" example:"1234"`
	Email           string `json:"email" valid:"email, required" binding:"required" example:"example@example.com"`
}

// DeleteInterest is a struct for deleting interest request
type DeleteInterest struct {
	Email string `json:"email" valid:"email, required" binding:"required"`
}
