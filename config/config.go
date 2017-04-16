package config

import "github.com/spf13/viper"

var _instance *viper.Viper

func instance() *viper.Viper {
    if _instance == nil {
        _instance = new()
    }
    return _instance
}

func new() *viper.Viper {
    viper.AddConfigPath("$HOME")
    viper.SetConfigName(".lgtm-capture")

    viper.SetDefault("font.ttf_path", "/Library/Fonts/Impact.ttf")
    viper.SetDefault("font.adjusted_value", 3.5)
    viper.SetDefault("font.base_color", "white")
    viper.SetDefault("font.border_color", "black")

    viper.ReadInConfig()

    return viper.GetViper()
}

func FontTTFPath() string {
    return instance().GetString("font.ttf_path")
}

func FontAdjustedValue() float64 {
    return instance().GetFloat64("font.adjusted_value")
}

func FontBaseColor() string {
    return instance().GetString("font.base_color")
}

func FontBorderColor() string {
    return instance().GetString("font.border_color")
}
