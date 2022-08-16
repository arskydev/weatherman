package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getMock(w http.ResponseWriter, r *http.Request) {
	mock := map[string]interface{}{
		"location": map[string]string{"city": "London", "country": "GB", "flag": "🇬🇧"},
		"weather":  map[string]interface{}{"temperature": 24, "unit": "℃", "weathersymbol": "☁️", "weathertype": "Clouds"},
		"suncycle": map[string]string{"sunrise": "20 Jul 22 05:26 +0400", "sunset": "20 Jul 22 20:38 +0400"},
	}

	resp, _ := json.Marshal(mock)
	w.Write(resp)
}
