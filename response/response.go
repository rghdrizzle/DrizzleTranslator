package response

type DetectResponse struct {
    Data struct {
        Detections [][]struct {
            Language string `json:"language"`
        } `json:"detections"`
    } `json:"data"`
}
type Translation struct {
    TranslatedText string `json:"translatedText"`
}

type TranslateResponse struct {
    Data struct {
        Translations []Translation `json:"translations"`
    } `json:"data"`
}

type TranslateRequest struct {
    Q      string `json:"q"`
    Source string `json:"source"`
    Target string `json:"target"`
    Format string `json:"format"`
}