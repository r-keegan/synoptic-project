package Session_test

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Services/Session"
	"testing"
	"time"
)

var sessionService Session.SessionService

func TestSessionService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Session Services Suite")
}

var _ = Describe("UserService", func() {
	const SessionTimeoutInSeconds = 1
	BeforeEach(func() {
		sessionService = Session.SessionService{MaxSessionLengthInSeconds: SessionTimeoutInSeconds}
	})

	Context("Create Session", func() {
		const cardId = "123"

		It("can create a session", func() {
			sessionService.CreateSession(cardId)

			Expect(sessionService.HasSession(cardId)).To(Equal(true))
		})

		It("can destroy a session", func() {
			sessionService.CreateSession(cardId)
			Expect(sessionService.HasSession(cardId)).To(Equal(true))

			sessionService.DestroySession(cardId)

			Expect(sessionService.HasSession(cardId)).To(Equal(false))
		})

		It("will invalidate a session after the configured timeout", func() {
			sessionService.CreateSession(cardId)
			Expect(sessionService.HasSession(cardId)).To(Equal(true))

			time.Sleep((SessionTimeoutInSeconds + 1) * time.Second) //Wait for session to timeout

			Expect(sessionService.HasSession(cardId)).To(Equal(false))
		})
	})
})
