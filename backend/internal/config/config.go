package config

type AppConfig struct {
	App struct {
		Host string
		Port string
		Name string
	}

	Watermark struct {
		Type         string
		Text         string
		WatermarkDir string `mapstructure:"watermark_dir"`
		Position     string
		Margin       int32
		Opacity      float32
		Scale        float32
	}

	Image struct {
		RandomSeed int32
		ImageDir   string `mapstructure:"image_dir"`
		OutputDir  string `mapstructure:"output_dir"`
	}

	Constants struct {
		StringLength int32 `mapstructure:"string_length"`
		RandomLimit  int32 `mapstructure:"random_limit"`
	}

	URL struct {
		Root   string
		Output string
		Image  struct {
			Root     string
			Upload   string
			Apply    string
			Download string
			Delete   string
		}
		Watermark struct {
			Root   string
			Upload string
			Apply  string
			Delete string
			List   string
			Get    string
		}
	}

	Database struct {
		Title    string
		Host     string
		Port     string
		User     string
		Password string
		Engine   string
		Name     string
		Sslmode  string
	}
}
