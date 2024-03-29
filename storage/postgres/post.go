package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	pb "template-post-service/genproto/post_service"
)

type postRepo struct {
	db *sql.DB
}

// NewPostRepo
func NewPostRepo(db *sql.DB) *postRepo {
	return &postRepo{db: db}
}

//rpc CreatePost(ReqPost) returns (RespPost);
//rpc UpdatePost(ReqPost) returns (ReqPost);
//rpc DeletePost(GetPostId) returns (ReqPost);
//rpc GetPostById(GetPostId) returns (RespPost);
//rpc GetPostsByOwnerId(GetOwnerId) returns (OwnerPosts);

func (p *postRepo) CreatePost(post *pb.ReqPost) (*pb.RespPost, error) {
	if post.Id == "" {
		post.Id = uuid.NewString()
	}
	query := `INSERT INTO posts(id, title, image_url, owner_id) VALUES($1, $2, $3, $4) RETURNING id, title, image_url, owner_id`

	var respPost pb.RespPost
	rowComment := p.db.QueryRow(query, post.Id, post.Title, post.ImageUrl, post.OwnerId)

	if err := rowComment.Scan(&respPost.Id, &respPost.Title, &respPost.ImageUrl, &respPost.OwnerId); err != nil {
		return nil, err
	}

	return &respPost, nil
}

func (p *postRepo) GetPostById(postId *pb.GetPostId) (*pb.RespPost, error) {
	query := `SELECT id, title, image_url, owner_id FROM posts WHERE id = $1`

	var respPost pb.RespPost
	rowPost := p.db.QueryRow(query, postId.PostId)

	if err := rowPost.Scan(&respPost.Id, &respPost.Title, &respPost.ImageUrl, &respPost.OwnerId); err != nil {
		return nil, err
	}

	return &respPost, nil
}

func (p *postRepo) GetPostsByOwnerId(ownerId *pb.GetOwnerId) (*pb.OwnerPosts, error) {
	query := `SELECT id, title, image_url, owner_id FROM posts WHERE owner_id = $1`

	var respPosts pb.OwnerPosts
	rows, err := p.db.Query(query, ownerId.OwnerId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respPost pb.ReqPost
		if err := rows.Scan(&respPost.Id, &respPost.Title, &respPost.ImageUrl, &respPost.OwnerId); err != nil {
			return nil, err
		}

		respPosts.Posts = append(respPosts.Posts, &respPost)
	}

	return &respPosts, nil
}

func (p *postRepo) UpdatePost(reqPost *pb.ReqPost) (*pb.ReqPost, error) {
	query := `UPDATE posts SET title = $1, image_url = $2, owner_id = $3 WHERE id = $4 RETURNING id, title, image_url, owner_id`

	rowPost := p.db.QueryRow(query, reqPost.Title, reqPost.ImageUrl, reqPost.OwnerId, reqPost.Id)
	if err := rowPost.Scan(&reqPost.Id, &reqPost.Title, &reqPost.ImageUrl, &reqPost.OwnerId); err != nil {
		return nil, err
	}

	return reqPost, nil
}

func (p *postRepo) DeletePost(postId *pb.GetPostId) (*pb.ReqPost, error) {
	query := `DELETE FROM posts WHERE id = $1 RETURNING id, title, image_url, owner_id`

	var deletedPost pb.ReqPost
	rowPost := p.db.QueryRow(query, postId.PostId)
	if err := rowPost.Scan(&deletedPost.Id, &deletedPost.Title, &deletedPost.ImageUrl, &deletedPost.OwnerId); err != nil {
		return nil, err 
	}

	return &deletedPost, nil
}