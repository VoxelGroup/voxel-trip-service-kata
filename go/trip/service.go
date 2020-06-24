package trip

import (
	"github.com/pkg/errors"
	"github.com/sandromancuso/trip-service-kata/go/user"
)

type Service struct {
	tripDAO *Dao
}

// New ...
func NewTripService(tripDao *Dao) *Service {
	return &Service{
		tripDAO: tripDao,
	}
}

func (this *Service) getTripByUser(friend *user.User) ([]Trip, error) {

	var trips []Trip

	friends, err := friend.Friends()
	if err != nil {
		return trips, err
	}
	loggedUser, err := user.CurrentSession().GetLoggedUser()
	if err != nil {
		return trips, err
	}

	var isFriend bool
	if loggedUser != nil {
		for _, friend := range friends {
			if *loggedUser == friend {
				isFriend = true
				break
			}
		}
		if isFriend {
			return this.tripDAO.FindTripByUser(friend)
		}
		return trips, err
	} else {
		return trips, errors.New("user not logged in")
	}
}
