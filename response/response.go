package response

type DetectResponse struct {
    Data struct {
        Detections [][]struct {
            Language string `json:"language"`
        } `json:"detections"`
    } `json:"data"`
}

type TranslateResponse struct{
	TargetText string `json:"targetText"`
}

type TranslateRequest struct {
    Q      string `json:"q"`
    Source string `json:"source"`
    Target string `json:"target"`
    Format string `json:"format"`
}