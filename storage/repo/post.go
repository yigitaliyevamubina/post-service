package repo

import pb "template-post-service/genproto/post_service"

//rpc CreatePost(ReqPost) returns (RespPost);
//rpc UpdatePost(ReqPost) returns (ReqPost);
//rpc DeletePost(GetPostId) returns (ReqPost);
//rpc GetPostById(GetPostId) returns (RespPost);
//rpc GetPostsByOwnerId(GetOwnerId) returns (OwnerPosts);

// PostStorageI
type PostStorageI interface {
	CreatePost(*pb.ReqPost) (*pb.RespPost, error)
	UpdatePost(*pb.ReqPost) (*pb.ReqPost, error)
	DeletePost(*pb.GetPostId) (*pb.ReqPost, error)
	GetPostById(*pb.GetPostId) (*pb.RespPost, error)
	GetPostsByOwnerId(*pb.GetOwnerId) (*pb.OwnerPosts, error)
}
