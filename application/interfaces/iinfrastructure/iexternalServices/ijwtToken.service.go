package iexternalServices

import "NetFarm/shared/models/common"

type IJwtTokenService interface {
	GenerateToken(email string, userId string, claim []string) (string, error)
	ValidateToken(tokenString string) (*common.Claims, error)
}
