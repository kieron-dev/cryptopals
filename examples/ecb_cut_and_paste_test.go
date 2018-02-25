package examples_test

import (
	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ECBCutAndPaste", func() {

	It("can parse a kv string", func() {
		str := "foo=bar&baz=qux&zap=zazzle"
		expected := map[string]string{
			"foo": "bar",
			"baz": "qux",
			"zap": "zazzle",
		}
		Expect(examples.ParseKVString(str)).To(Equal(expected))
	})

	It("can encode a hash", func() {
		hash := map[string]string{
			"foo": "bar",
			"baz": "qux",
			"zap": "zazzle",
		}
		enc := examples.EncodeKVs(hash)
		Expect(enc).To(Equal("baz=qux&foo=bar&zap=zazzle"))
	})

	It("always encodes email first", func() {
		hash := map[string]string{
			"foo":   "bar",
			"email": "foo@bar.com",
			"bob":   "elephant",
		}
		Expect(examples.EncodeKVs(hash)).To(MatchRegexp("^email=foo@bar.com"))
	})

	It("always encodes role last", func() {
		hash := map[string]string{
			"role":  "user",
			"email": "foo@bar.com",
			"uid":   "10",
		}
		Expect(examples.EncodeKVs(hash)).To(MatchRegexp("role=user$"))
	})

	It("can generate a profile", func() {
		hash := map[string]string{
			"email": "foo@bar.com",
			"uid":   "10",
			"role":  "user",
		}
		profile := examples.ProfileFor("foo@bar.com")
		Expect(profile).To(Equal(hash))
		Expect(examples.EncodeKVs(profile)).To(Equal("email=foo@bar.com&uid=10&role=user"))
	})

	It("strips dodgy chars from email", func() {
		hash := map[string]string{
			"email": "foo@bar.com",
			"uid":   "10",
			"role":  "user",
		}
		Expect(examples.ProfileFor("foo=@bar.com&")).To(Equal(hash))
	})

	It("generates an encrypted cookie", func() {
		cookie := examples.GetCookie("foo@bar.com")
		Expect(len(cookie)).To(BeNumerically(">", 11))
	})

	It("can decrypt a cookie", func() {
		cookie := examples.GetCookie("foo@bar.com")
		hash := examples.DecryptCookie(cookie)
		Expect(hash["email"]).To(Equal("foo@bar.com"))
	})
})
