package configs

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "cache_host",
		Destination: &CacheHost,
		EnvVars:     []string{"CACHE_HOST"},
		Value:       "localhost",
	},
	&cli.StringFlag{
		Name:        "cache_port",
		Destination: &CachePort,
		EnvVars:     []string{"CACHE_PORT"},
		Value:       "6379",
	},
	&cli.StringFlag{
		Name:        "cache_password",
		Destination: &CachePassword,
		EnvVars:     []string{"CACHE_PASSWORD"},
		Value:       "",
	},
	&cli.IntFlag{
		Name:        "cache_database",
		Destination: &CacheDb,
		EnvVars:     []string{"CACHE_DATABASE"},
		Value:       0,
	},

	&cli.StringFlag{
		Name:        "database_host",
		Destination: &DatabaseHost,
		EnvVars:     []string{"DATABASE_HOST"},
		Value:       "localhost",
	},
	&cli.StringFlag{
		Name:        "database_port",
		Destination: &DatabasePort,
		EnvVars:     []string{"DATABASE_PORT"},
		Value:       "5432",
	},
	&cli.StringFlag{
		Name:        "database_dbname",
		Destination: &DatabaseName,
		EnvVars:     []string{"DATABASE_DBNAME"},
		Value:       "authentication",
	},
	&cli.StringFlag{
		Name:        "database_user",
		Destination: &DatabaseUser,
		EnvVars:     []string{"DATABASE_USER"},
		Value:       "authuser",
	},
	&cli.StringFlag{
		Name:        "database_password",
		Destination: &DatabasePassword,
		EnvVars:     []string{"DATABASE_PASSWORD"},
		Value:       "somepassword",
	},

	&cli.StringFlag{
		Name:        "app_env",
		Destination: &AppEnvironment,
		EnvVars:     []string{"APP_ENV"},
		Value:       "dev",
	},
	&cli.StringFlag{
		Name:        "app_port",
		Destination: &AppPort,
		EnvVars:     []string{"APP_PORT"},
		Value:       ":9000",
	},
	&cli.StringFlag{
		Name:        "app_id",
		Destination: &AppId,
		EnvVars:     []string{"APP_ID"},
		Value:       "3490be09e8904918997b073c460c834c",
	},

	&cli.StringFlag{
		Name:        "jwt_public",
		Destination: &JwtPublic,
		EnvVars:     []string{"JWT_PUBLIC"},
		Value: `ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAFeQDDdKv96vAHOivPGDiEkSt02E8qFmGNz+aFv/hZ0fkP3QLeHxm7HCMZguNcCj3HFTi3HERj8jN0nHvTXbTM6TwEG6evRIzm8qZlzfl3CsBIRtiAFmdKKHHmO9sibd7gZZifU8+emZReFF1ZyYL0v5HuT8M2vs67J2vSsyTxjkyD4ww== ppcamp@DESKTOP-14OV55P
`,
	},
	&cli.StringFlag{
		Name:        "jwt_private",
		Destination: &JwtPrivate,
		EnvVars:     []string{"JWT_PRIVATE"},
		Value: `-----BEGIN OPENSSH PRIVATE KEY-----
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
	},
}
