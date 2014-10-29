package emoticon

type Type string

const (
	Global = Type("global")
	Group  = Type("group")
	All    = Type("all")
)

type Emoticon struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	AudioPath string `json:"audio_path"`
	ID        string `json:"id"`
	Shortcut  string `json:"shortcut"`
}

type emoticonIter struct {
}

type EmoticonCache map[string]Emoticon

/*
  How do I expect people to use the Iterative structure?
  I need to answer that question first. Or should I build a raw version first.
  I'm leaning toward the raw as that is faster.
*/
