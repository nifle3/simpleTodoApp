package user

import "context"

func (u UseCase) Delete(userID string, ctx context.Context) error {
	return u.storage.DeleteUser(userID, ctx)
}
