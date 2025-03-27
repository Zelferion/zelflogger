package zelflogger

type Config struct {
	RedisLogging    bool
	PostgresLogging bool
	HttpReqLogging  bool
	HttpRespLogging bool
	HeadersLogging  bool
	BodyLogging     bool
	ErrorLogging    bool

	TimestampsLogging bool
  Debug bool
  DebugOne bool
  DebugTwo bool
  DebugThree bool
  DebugFour bool
}

var Conf = &Config{
	RedisLogging:    true,
	PostgresLogging: true,
	HttpReqLogging:  true,
	HttpRespLogging: true,
	HeadersLogging:  true,
	BodyLogging:     true,
	ErrorLogging:    true,

	TimestampsLogging: true,
  Debug: true,
  DebugOne: true,
  DebugTwo: true,
  DebugThree: true,
  DebugFour: true,
}

// func Init(obj Config) *Config {
// 	return &Config{
// 		RedisLogging:      obj.RedisLogging,
// 		PostgresLogging:   obj.PostgresLogging,
// 		HttpReqLogging:    obj.HttpReqLogging,
// 		ErrorLogging:      obj.ErrorLogging,
// 		TimestampsLogging: obj.TimestampsLogging,
// 	}
// }

