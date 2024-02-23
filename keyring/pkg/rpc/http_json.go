// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
	return err
}

func DecodeJSON(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}
