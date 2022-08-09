package emailSender

import "testing"

func TestSendEmail(t *testing.T) {
	if err := SendEmail("thanaponpuipui@gmail.com", "testCode.png"); err != nil {
		t.Error(err)
	}
}
