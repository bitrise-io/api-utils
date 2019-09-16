package security

// JWTMock ...
type JWTMock struct {
	SignFn   func(authToken string) (string, error)
	VerifyFn func(jwtToken string) (bool, error)
}

// Sign ...
func (j *JWTMock) Sign(authToken string) (string, error) {
	if j.SignFn == nil {
		panic("You have to override JWTService.Sign function in tests")
	}
	return j.SignFn(authToken)
}

// Verify ...
func (j *JWTMock) Verify(jwtToken string) (bool, error) {
	if j.VerifyFn == nil {
		panic("You have to override JWTService.Verify function in tests")
	}
	return j.VerifyFn(jwtToken)
}
