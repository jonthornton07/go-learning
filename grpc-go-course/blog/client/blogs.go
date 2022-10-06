package main

import (
	"context"
	pb "grpc/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Jon",
		Title:    "My First Blog",
		Content:  "Contents",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--readBlog was invoked---")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--updateBlog was invoked---")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "Not Jon",
		Title:    "A new title",
		Content:  "New content",
	}

	_, err := c.UpdateBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading %v\n", err)
	}

	log.Printf("Blog was updated")
}

func listBlogs(c pb.BlogServiceClient) {
	log.Println("--listBlog was invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happend %v\n", err)
		}

		log.Println(res)
	}

}

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--deleteBlog was invoked---")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while deleting %v\n", err)
	}

	log.Println("Blog was deleted")
}
