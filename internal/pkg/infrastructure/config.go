package infrastructure

type Config struct {
	Port        string `env:"PORT,required"`
	Environment string `env:"ENV,required"`
	GRPCPort    string `env:"GRPC_PORT,required"`
}
