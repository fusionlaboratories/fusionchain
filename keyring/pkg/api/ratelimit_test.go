// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
	n := "test"
	log, err := logger.NewLogger("fatal", "plain", false, n)
	if err != nil {
		t.Fatal(err)
	}
	m := newMockKeyRing("", log, database.NewMemory(), &mockModule{})
	apiTests := []struct {
		name             string
		endpoint         string
		ratelimit        int
		requests         int
		expectedResponse any
		expectedCode     int
	}{
		{
			"too many (low)",
			PubKeysEndPnt,
			1,
			2,
			map[string]string{"error": errTooManyRequests.Error()},
			http.StatusBadRequest,
		},
		{
			"high",
			PubKeysEndPnt,
			100,
			100,
			&Response{Message: "OK", Version: common.FullVersion, Service: "test"},
			http.StatusOK,
		},
		{
			"too many (high)",
			PubKeysEndPnt,
			100,
			101,
			map[string]string{"error": errTooManyRequests.Error()},
			http.StatusBadRequest,
		},
	}

	for _, tt := range apiTests {
		t.Run(tt.name, func(t *testing.T) {
			method := WithRateLimit(tt.ratelimit, time.Second, PasswordProtected(m.password, HandlePubKeyRequest(m.log, m.db, n)))
			httpReq := httptest.NewRequest(http.MethodGet, tt.endpoint, nil)
			respRecorder := httptest.NewRecorder()
			for i := 0; i < tt.requests; i++ {
				respRecorder = httptest.NewRecorder()
				method(respRecorder, httpReq)
			}
			if g, w := respRecorder.Code, tt.expectedCode; g != w {
				t.Errorf("unexpected response code, want %v got %v", w, g)
			}
			expectedJSON, _ := json.Marshal(tt.expectedResponse)

			if g, w := respRecorder.Body.Bytes(), expectedJSON; !bytes.Equal(g, w) {
				t.Fatalf("unexpected response, want %s, got %s", w, g)
			}
		})
	}
}

func Test_TokenRefill(t *testing.T) {
	n := "test"
	log, err := logger.NewLogger("fatal", "plain", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	m := newMockKeyRing("", log, database.NewMemory(), &mockModule{})

	duration := 10 * time.Millisecond
	pause := 11 * time.Millisecond

	method := WithRateLimit(1, duration, HandleStatusRequest(m.log, n))
	httpReq := httptest.NewRequest(http.MethodGet, PubKeysEndPnt, nil)
	respRecorder := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		respRecorder = httptest.NewRecorder()
		method(respRecorder, httpReq)
		time.Sleep(pause) // Ensure time elapses for the tokens to refill
	}

	if g, w := respRecorder.Code, http.StatusOK; g != w {
		t.Errorf("unexpected response code, want %v got %v", w, g)
	}

	expectedJSON, _ := json.Marshal(&Response{Message: "OK", Version: common.FullVersion, Service: "test"})

	if g, w := respRecorder.Body.Bytes(), expectedJSON; !bytes.Equal(g, w) {
		t.Fatalf("unexpected response, want %s, got %s", w, g)
	}
}
