package config

/**
 * @author eyesYeager
 * @date 2023/2/8 15:15
 */

type Nacos struct {
	Ip                  string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port                uint64 `mapstructure:"port" json:"port" yaml:"port"`
	NamespaceId         string `mapstructure:"namespace_id" json:"namespace_id" yaml:"namespace_id"`
	DataId              string `mapstructure:"data_id" json:"data_id" yaml:"data_id"`
	Group               string `mapstructure:"group" json:"group" yaml:"group"`
	Username            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Timeout             uint64 `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	NotLoadCacheAtStart bool   `mapstructure:"not_load_cache_at_start" json:"not_load_cache_at_start" yaml:"not_load_cache_at_start"`
	LogDir              string `mapstructure:"log_dir" json:"log_dir" yaml:"log_dir"`
	CacheDir            string `mapstructure:"cache_dir" json:"cache_dir" yaml:"cache_dir"`
	LogLevel            string `mapstructure:"log_level" json:"log_level" yaml:"log_level"`
}
