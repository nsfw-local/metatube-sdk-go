package gcolle

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcolle_GetMovieInfoByID(t *testing.T) {
	provider := New()
	for _, item := range []string{
		"847256",
		"848234",
		"845371",
		"839979",
		"848315",
	} {
		info, err := provider.GetMovieInfoByID(item)
		data, _ := json.MarshalIndent(info, "", "\t")
		assert.True(t, assert.NoError(t, err) && assert.True(t, info.Valid()))
		t.Logf("%s", data)
	}
}
