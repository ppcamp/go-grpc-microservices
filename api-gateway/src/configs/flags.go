package configs

import "flag"

var APP_PORT = flag.String("PORT", ":8080", "The port that gRPC server will listen")
