package helper

type ReqBody struct{
	SourceLang string `json:"sourceLang"`
	SourceText string `json:"sourceText"`
	TargetLang string `json:"targetLang"`
}
