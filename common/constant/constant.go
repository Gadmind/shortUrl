package constant

const (

	// SSOPrefix 短链接-》原始链接缓存前缀
	SSOPrefix = "su:short:origin:"

	// SOSPrefix 原始链接-》短链接缓存前缀
	SOSPrefix = "su:origin:short:"

	// RandPrefix 随机数缓存前缀
	RandPrefix = "su:rand:"

	// MaxRandomRange 短链接最大随机范围
	MaxRandomRange = 56800235583

	// MinRandomRange 短链接最小随机范围
	MinRandomRange = 14776336
)
