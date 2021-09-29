package main

type LinkRequest struct {
	Link string `form:"link" json:"link" binding:"required"`
}

type LinkResponse struct {
	ID   string `form:"id" json:"id"`
	Link string `form:"link" json:"link"`
}
