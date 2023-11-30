# Fusion - MPC model service

This is a mini server that mocks the `signature_request_by_id` response from `fusiond`. If the request ID is odd the response will have status `SIGN_REQUEST_STATUS_PENDING`.

```
curl 'localhost:1717/fusionchain/treasury/signature_request_by_id?id=311' (GET)
```
```
{
   "sign_request":{
      "id":"0000000000000000000000000000000000000000000000000000000000000137",
      "creator":"qredo1d652c9nngq5cneak2whyaqa4g9ehr8psyl0t7j",
      "key_id":"0000000000000000000000000000000000000000000000000000000000000002",
      "data_for_signing":"tSR4wa1srbASeiRWjzEKKC1PgSuPBuzuWosOEdj3NB0=",
      "status":"SIGN_REQUEST_STATUS_PENDING",
   }
}
```

If the request ID is even the response will have status `SIGN_REQUEST_STATUS_FULFILLED`.

```
curl 'localhost:1717/fusionchain/treasury/signature_request_by_id?id=312' (GET)
```

```
{
   "sign_request":{
      "id":"0000000000000000000000000000000000000000000000000000000000000138",
      "creator":"qredo1d652c9nngq5cneak2whyaqa4g9ehr8psyl0t7j",
      "key_id":"0000000000000000000000000000000000000000000000000000000000000002",
      "data_for_signing":"tSR4wa1srbASeiRWjzEKKC1PgSuPBuzuWosOEdj3NB0=",
      "status":"SIGN_REQUEST_STATUS_FULFILLED",
      "signed_data":"Lnhyih8OH9e9IA0BkGIC+/ati2xKBoHia6Z9srNnhsQgFnlNJZyn7inUunUZ4lAIGIJ/wV1iBV7FmSzrGWsmXQA="
   }
}
```

## Example usage

```
$ go run .
```

```
$ curl 'localhost:8000/fusionchain/treasury/signature_request_by_id?id=3155'
{"id":"3155","creator":"qredo1d652c9nngq5cneak2whyaqa4g9ehr8psyl0t7j","key_id":"0000000000000000000000000000000000000000000000000000000000000001","data_for_signing":"tSR4wa1srbASeiRWjzEKKC1PgSuPBuzuWosOEdj3NB0=","status":"SIGN_REQUEST_STATUS_PENDING"}

$ curl 'localhost:8000/fusionchain/treasury/signature_request_by_id?id=3156'
{"id":"3156","creator":"qredo1d652c9nngq5cneak2whyaqa4g9ehr8psyl0t7j","key_id":"0000000000000000000000000000000000000000000000000000000000000001","data_for_signing":"tSR4wa1srbASeiRWjzEKKC1PgSuPBuzuWosOEdj3NB0=","status":"SIGN_REQUEST_STATUS_FULFILLED","signed_data":"Lnhyih8OH9e9IA0BkGIC+/ati2xKBoHia6Z9srNnhsQgFnlNJZyn7inUunUZ4lAIGIJ/wV1iBV7FmSzrGWsmXQA="}

```

## Run with docker

TODO

