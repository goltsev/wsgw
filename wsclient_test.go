package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignature(t *testing.T) {
	testCases := []struct {
		desc     string
		secret   string
		verb     string
		path     string
		expires  time.Time
		data     []byte
		expected string
	}{
		{
			desc:     "empty string data",
			secret:   "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO",
			verb:     "GET",
			path:     "/api/v1/instrument",
			expires:  time.Unix(1518064236, 0), // 2018-02-08T04:30:36Z
			data:     []byte(""),
			expected: "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00",
		},
		{
			desc:     "nil data",
			secret:   "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO",
			verb:     "GET",
			path:     "/api/v1/instrument",
			expires:  time.Unix(1518064236, 0), // 2018-02-08T04:30:36Z
			data:     nil,
			expected: "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			msg := message(tC.verb, tC.path, tC.expires, tC.data)
			got := signatureString(tC.secret, msg)
			assert.Equal(t, tC.expected, got)
		})
	}
}
