package configs

import "flag"

var (
	APP_PORT = flag.String(
		"PORT",
		":9001",
		"The port that gRPC server will listen",
	)

	JWT_PRIVATE = flag.String(
		"JWT_PRIVATE",
		`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAArAAAABNlY2RzYS
1zaGEyLW5pc3RwNTIxAAAACG5pc3RwNTIxAAAAhQQBXkAw3Sr/erwBzorzxg4hJErdNhPK
hZhjc/mhb/4WdH5D90C3h8ZuxwjGYLjXAo9xxU4txxEY/IzdJx70120zOk8BBunr0SM5vK
mZc35dwrASEbYgBZnSihx5jvbIm3e4GWYn1PPnpmUXhRdWcmC9L+R7k/DNr7Ouydr0rMk8
Y5Mg+MMAAAEYbnyzrm58s64AAAATZWNkc2Etc2hhMi1uaXN0cDUyMQAAAAhuaXN0cDUyMQ
AAAIUEAV5AMN0q/3q8Ac6K88YOISRK3TYTyoWYY3P5oW/+FnR+Q/dAt4fGbscIxmC41wKP
ccVOLccRGPyM3Sce9NdtMzpPAQbp69EjObypmXN+XcKwEhG2IAWZ0ooceY72yJt3uBlmJ9
Tz56ZlF4UXVnJgvS/ke5Pwza+zrsna9KzJPGOTIPjDAAAAQgFbCxbLXNkxBcdk+46SXOwr
x8tIUfjKNd+LZoiu7vFTk+V2L8jvaKlj3anxhcyrvSf28D8Jna1LZ5Ru+AaXgFLJCgAAAB
ZwcGNhbXBAREVTS1RPUC0xNE9WNTVQAQIDBA==
-----END OPENSSH PRIVATE KEY-----
`,
		"JWT secret token",
	)

	JWT_PUBLIC = flag.String(
		"JWT_PUBLIC",
		`ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAFeQDDdKv96vAHOivPGDiEkSt02E8qFmGNz+aFv/hZ0fkP3QLeHxm7HCMZguNcCj3HFTi3HERj8jN0nHvTXbTM6TwEG6evRIzm8qZlzfl3CsBIRtiAFmdKKHHmO9sibd7gZZifU8+emZReFF1ZyYL0v5HuT8M2vs67J2vSsyTxjkyD4ww== ppcamp@DESKTOP-14OV55P
`,
		"JWT public token",
	)

	DATABASE_QUERY = flag.String(
		"DATABASE_QUERY",
		"host=localhost port=5431 user=authuser password=somepassword dbname=users sslmode=disable application_name=users",
		"Query used to connect with psql driver",
	)

	APP_ID = flag.String(
		"APP_ID",
		"3490be09e8904918997b073c460c834c",
		"The app unique ID, it can be user to store/retrieve data from cache")
)
