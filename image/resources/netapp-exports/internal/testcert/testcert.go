/*
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package testcert

import "strings"

// generated
// go run gen_certs.go > testcert.go

// PEM encoded TLS certificate with SAN for localhost and 127.0.0.1
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIB6TCCAVKgAwIBAgIBATANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQKEwdBY21l
IENvMCAXDTIxMTIxMDEzMjA1OFoYDzIxMzYwMTA5MDUyMDU4WjASMRAwDgYDVQQK
EwdBY21lIENvMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4jvUVOyDk6Scl
dDCcj6RmTrLcCvCTrBHaYsJ0KHS7HYq4SkRbzSUi3laG7xyP1yXlV5AaQgl4CupK
9/XkG3YrP7QuQNC9M1+VAQPhuXKCT0MHA2SPm1fXEV55mDjWcJCKmGm9OzumevAF
clBe2+7B4s031pCUVXee16AZlMD7PQIDAQABo00wSzAOBgNVHQ8BAf8EBAMCAqQw
HQYDVR0OBBYEFNoGahzSKBFqV3nyHbdMio/rsLDoMBoGA1UdEQQTMBGCCWxvY2Fs
aG9zdIcEfwAAATANBgkqhkiG9w0BAQsFAAOBgQCYJwhmi7hMQHpNcaeALbJIT3vU
r4Vgi/fNEn0/hMV3tUL4Ji8uKpMiiYuX5Pemz3ni51FvPYfVvuH12/MFZ1IFCqVf
yzyl6lJBerYcfuvIQKkLnzhhjaE4ei97RICQ5N11aShkPiY2uiJL/87QF3zdrQZl
xb9jU52NaFqdH5ub1Q==
-----END CERTIFICATE-----`)

// PEM encoded TLS certificate with CN for localhost (no SAN)
var CommonNameCert = []byte(`-----BEGIN CERTIFICATE-----
MIIB9TCCAV6gAwIBAgIBATANBgkqhkiG9w0BAQsFADAmMRAwDgYDVQQKEwdBY21l
IENvMRIwEAYDVQQDEwlsb2NhbGhvc3QwIBcNMjExMjEwMTMyMDU4WhgPMjEzNjAx
MDkwNTIwNThaMCYxEDAOBgNVBAoTB0FjbWUgQ28xEjAQBgNVBAMTCWxvY2FsaG9z
dDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAuI71FTsg5OknJXQwnI+kZk6y
3Arwk6wR2mLCdCh0ux2KuEpEW80lIt5Whu8cj9cl5VeQGkIJeArqSvf15Bt2Kz+0
LkDQvTNflQED4blygk9DBwNkj5tX1xFeeZg41nCQiphpvTs7pnrwBXJQXtvuweLN
N9aQlFV3ntegGZTA+z0CAwEAAaMxMC8wDgYDVR0PAQH/BAQDAgKkMB0GA1UdDgQW
BBTaBmoc0igRald58h23TIqP67Cw6DANBgkqhkiG9w0BAQsFAAOBgQCKBkpEFKVK
tTxwTfFxFbqzg49vUZ4lGMJxukQG19u0yrM8rAyp7fqYY3jn5oqquPrQ7OPjEw4R
+K4BmOgx3OaoqtlQZ5fyEvDIrZFZteqlo679v2LViwWs2K9Wq+SIOdJptCVpRkl5
lnPDq/e/bC+l9dco4VTdkhcMZAuvKRlwHQ==
-----END CERTIFICATE-----`)

// PEM encoded private key for above TLS certificates
var PrivateKey = []byte(testkey(`-----BEGIN TEST KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALiO9RU7IOTpJyV0
MJyPpGZOstwK8JOsEdpiwnQodLsdirhKRFvNJSLeVobvHI/XJeVXkBpCCXgK6kr3
9eQbdis/tC5A0L0zX5UBA+G5coJPQwcDZI+bV9cRXnmYONZwkIqYab07O6Z68AVy
UF7b7sHizTfWkJRVd57XoBmUwPs9AgMBAAECgYBEzYawk4p/zCu72sUEmMhBG7Wy
MqHda5h8QbUceLiLyUedzJIPZzsg9KJtS1bqiNqn1SzznQrKpccSi74ve81hUBbn
MHD7GThr21X22IJskS896SpLXYNiek3RhCPuVtaY00S2cqBpcLIfwN5T0rFx1UVJ
3EXQ1VmHtmgZ9vsqAQJBANWhASUvq2n6LzxfE3jJLKpb+WskdvR3bVFy0SHWYSqG
CsvHIQ/vJM1z6U2inTcUvmAEp+0DT6GvthPUOUXansUCQQDdKeZzZEJ3NhE70rf4
pJZRn/q63iRONBHtA9fi9s0tLV+szsR+4kIgido50pcM6Ax+WZZvK+UGfavorl24
ITIZAkEAh14Wk7G5NNZLyD2W4RrZKrpNOg9JMW/b3Zib3I0z1PZLMQVldetbrrSc
SY2ZgaWrXLyWjCFk/FeTUM1R3WnC4QJAd7f6u3QGVqm54nxKghn2FPFYtFcTqBGf
soFmF8Iphs8M/2peC6FG0n5M1wgcJCxotuyf9kX4j+7vY+EfclDl4QJBAK2C0cdu
XWFpB8IjpwqFGTxZlbdGE2UtwtVW+8KGiaIGahM9KR1fy8TNFNbvKgGsLgScofJ3
eAsfr4ywTM7E/KY=
-----END TEST KEY-----`))

func testkey(key string) string {
	return strings.ReplaceAll(key, "TEST KEY", "PRIVATE KEY")
}
