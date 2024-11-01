package generators

type Metadata struct {
	Title         string   `yaml:"title"`
	Description   string   `yaml:"description"`
	Banner        string   `yaml:"banner"`
	Badge         string   `yaml:"badge"`
	IntroQuestion string   `yaml:"intro_question"`
	IntroText     string   `yaml:"intro_text"`
	Purposes      []string `yaml:"purposes"`
	Notice        string   `yaml:"notice"`
}

type Item struct {
	Title       string   `yaml:"title"`
	URL         string   `yaml:"url"`
	Description string   `yaml:"description"`
	Docs        string   `yaml:"docs,omitempty"`
	Thumbnail   string   `yaml:"thumbnail,omitempty"`
	Added       string   `yaml:"added,omitempty"`
	Tags        []string `yaml:"tags,omitempty"`
}

type Section struct {
	Section     string `yaml:"section"`
	Description string `yaml:"description,omitempty"`
	Items       []Item `yaml:"items"`
}

type ListData struct {
	Metadata Metadata  `yaml:"metadata"`
	Contents []Section `yaml:"contents"`
}
