package api

type HelloReq struct {
    Name    string `json:"name"`
    Message string `json:"message"`
}

type HelloResp struct {
    Reply string `json:"reply"`
}
