using System;
using Xunit;
using FluentAssertions;
using TripServiceKata.Exception;
using TripServiceKata.Trip;

namespace TripServiceKata.Tests
{
    public class TripServiceTest
    {
        [Fact]
        public void ShouldFailsIfTheUserIsNotLoggedIn()
        {
            var tripService = new TestableTripService();

            Action act = () => tripService.GetTripsByUser(null);

            act.Should().Throw<UserNotLoggedInException>();
        }
    }

    public class TestableTripService : TripService
    {
        protected override User.User GetLoggedUser()
        {
            return null;
        }
    }
}
