// Code generated by goctl. DO NOT EDIT.
package types

type File struct {
	Title string `form:"title"`
	Id    string `form:"id"`
}

type Posts struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Id      string `form:"id"`
}

type SearchForFilesReq struct {
	Content string `form:"content"`
}

type SearchForFilesResp struct {
	Files []File `form:"files"`
	Error string `form:"error"`
}

type SearchForPostsReq struct {
	Content string `form:"content"`
}

type SearchForPostsResp struct {
	Posts []Posts `form:"posts"`
	Error string  `form:"error"`
}

type SearchForFilesByUserIdReq struct {
	UserId    int64  `form:"userId"`
	TypeMount string `form:"typeMount"`
}

type SearchForFilesByUserIdResp struct {
	Files []File `form:"files"`
	Error string `form:"error"`
}

type SearchForPostsByUserIdReq struct {
	UserId    int64  `form:"userId"`
	TypeMount string `form:"typeMount"`
}

type SearchForPostsByUserIdResp struct {
	Posts []Posts `form:"posts"`
	Error string  `form:"error"`
}

type SearchForFilesRankReq struct {
	Rank      int64  `form:"rank"`
	TypeMount string `form:"typeMount"`
}

type SearchForFilesRankResp struct {
	Files []File `form:"files"`
	Error string `form:"error"`
}

type SearchForPostsRankReq struct {
	Rank      int64  `form:"rank"`
	TypeMount string `form:"typeMount"`
}

type SearchForPostsRankResp struct {
	Posts []Posts `form:"posts"`
	Error string  `form:"error"`
}