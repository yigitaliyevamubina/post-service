package service

import (
	"context"
	"database/sql"
	"log"
	pbc "template-post-service/genproto/comment_service"
	pb "template-post-service/genproto/post_service"
	pbu "template-post-service/genproto/user_service"
	"template-post-service/pkg/logger"
	grpcclient "template-post-service/service/grpc_client"
	"template-post-service/storage"
)

// PostService
type PostService struct {
	storage storage.IStorage
	logger  logger.Logger
	client  grpcclient.IServiceManager
	pb.UnimplementedPostServiceServer
}

// NewPostService
func NewPostService(db *sql.DB, log logger.Logger, client grpcclient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

//rpc CreatePost(ReqPost) returns (RespPost);
//rpc GetPostById(GetPostId) returns (RespPost);
//rpc GetPostsByOwnerId(GetOwnerId) returns (OwnerPosts);

func (p *PostService) CreatePost(ctx context.Context, post *pb.ReqPost) (*pb.RespPost, error) {
	respPost, err := p.storage.Post().CreatePost(post)
	if err != nil {
		log.Fatal("create post", err.Error())
	}

	return respPost, nil
}

func (p *PostService) GetPostById(ctx context.Context, postId *pb.GetPostId) (*pb.RespPost, error) {
	respPost, err := p.storage.Post().GetPostById(postId)
	if err != nil {
		log.Fatal("get post by id", err.Error())
	}

	user, err := p.client.UserService().GetUserById(ctx, &pbu.GetUserId{
		UserId: respPost.OwnerId,
	})
	if err != nil {
		log.Fatal("cannot get owner of the post", err.Error())
	}
	if respPost.Owner == nil {
		respPost.Owner = &pb.Owner{}
	}
	respPost.Owner.Id = user.Id
	respPost.Owner.FirstName = user.FirstName
	respPost.Owner.LastName = user.LastName
	respPost.Owner.Age = user.Age
	respPost.Owner.Gender = pb.Gender(user.Gender)

	comments, err := p.client.CommentService().GetAllCommentsByPostId(ctx, &pbc.GetPostID{
		PostId: respPost.Id,
	})
	if err != nil {
		log.Fatal("cannot get comments of the post", err.Error())
	}

	var respComments []*pb.Comment
	for _, comment := range comments.Comments {
		c := pb.Comment{
			Id:        comment.Id,
			Content:   comment.Content,
			OwnerId:   comment.OwnerId,
			PostId:    comment.PostId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			DeletedAt: comment.DeletedAt,
		}
		respComments = append(respComments, &c)
	}

	respPost.Comments = respComments

	return respPost, nil
}

func (p *PostService) GetPostsByOwnerId(ctx context.Context, ownerId *pb.GetOwnerId) (*pb.OwnerPosts, error) {
	ownerPosts, err := p.storage.Post().GetPostsByOwnerId(ownerId)
	if err != nil {
		log.Fatal("get posts by owner id", err.Error())
	}

	return ownerPosts, nil
}
