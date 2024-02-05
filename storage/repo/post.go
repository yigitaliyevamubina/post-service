package repo

import pb "template-post-service/genproto/post_service"

//rpc CreatePost(ReqPost) returns (RespPost);
//rpc GetPostById(GetPostId) returns (RespPost);
//rpc GetPostsByOwnerId(GetOwnerId) returns (OwnerPosts);

// PostStorageI
type PostStorageI interface {
	CreatePost(*pb.ReqPost) (*pb.RespPost, error)
	GetPostById(*pb.GetPostId) (*pb.RespPost, error)
	GetPostsByOwnerId(*pb.GetOwnerId) (*pb.OwnerPosts, error)
}
