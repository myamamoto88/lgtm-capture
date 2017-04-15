package config

import "github.com/spf13/viper"

var _instance *viper.Viper

func Instance() *viper.Viper {
    if _instance == nil {
        _instance = new()
    }
    return _instance
}

func new() *viper.Viper {
    viper.AddConfigPath("$HOME")
    viper.SetConfigName(".lgtm-capture")
    viper.SetConfigType("yml")

    viper.SetDefault("font.ttf", "/Library/Fonts/Impact.ttf")
    viper.SetDefault("font.adjusted_value", 3.5)
    viper.SetDefault("font.base_color", "white")
    viper.SetDefault("font.border_color", "black")

    viper.ReadInConfig()

    return viper.GetViper()
}
