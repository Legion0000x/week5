package unit

import (
	// "fmt"
	"testing"

	"example.com/SE-LAB/entity"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "12345678900", // ผิดตรงนี้ มี 11 ตัว
			Password:    "12345678",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits."))

	})
}

func TestUrl(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`check_url`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1234567890", //ถูก
			Password:    "12345678",
			Url:         "https:///www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Url is invalid"))

	})
}

func TestAllTrue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`all_true_case`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1234567890", //ถูก
			Password:    "12345678",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		//g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits."))

	})
}