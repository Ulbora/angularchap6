package main

import (
	"time"
)

// {
// 	"_id": 1,
// 	"date": 1400623623107,
// 	"introText": "This is a blog post about AngularJS. We will cover how to build",
// 	"blogText": "This is a blog post about AngularJS. We will cover how to build a blog and how to add comments to the blog post."
// },

// "comments" :[
//                 {
//                     "commentText" : "Very good post. I love it."
//                 },
//                 {
//                     "commentText" : "When can we learn services."
//                 }

//             ]
// {
// 	"_id": 2,
// 	"date": 1400267723107,
// 	"introText": "In this blog post we will learn how to build applications based on REST",
// 	"blogText": "In this blog post we will learn how to build applications based on REST web services that contain most of the business logic needed for the application."
// }

// "comments" :[
//                 {
//                     "commentText" : "REST is great. I want to know more."
//                 },
//                 {
//                     "commentText" : "Will we use Node.js for REST services?"
//                 }

//             ]

var blogs []Blog

//Blog Blog
type Blog struct {
	ID        int64      `json:"_id"`
	Bdate     time.Time  `json:"date"`
	InteoText string     `json:"introText"`
	BlogText  string     `json:"blogText"`
	Comments  *[]Comment `json:"comments"`
}

//Comment Comment
type Comment struct {
	CommentText string `json:"commentText"`
}

//Init Init
func Init() {
	var b1 Blog
	b1.ID = 1
	b1.Bdate = time.Now()
	b1.InteoText = "This is a blog post about AngularJS. We will cover how to build"
	b1.BlogText = "In this blog post we will learn how to build applications based on REST web services that contain most of the business logic needed for the application."

	var c11 Comment
	c11.CommentText = "Very good post. I love it."

	var c12 Comment
	c12.CommentText = "When can we learn services."
	var cmts1 []Comment
	cmts1 = append(cmts1, c11)
	cmts1 = append(cmts1, c12)
	b1.Comments = &cmts1

	var b2 Blog
	b2.ID = 2
	b2.Bdate = time.Now()
	b2.InteoText = "In this blog post we will learn how to build applications based on REST"
	b2.BlogText = "This is a blog post about AngularJS. We will cover how to build a blog and how to add comments to the blog post."

	var c21 Comment
	c21.CommentText = "REST is great. I want to know more."

	var c22 Comment
	c22.CommentText = "Will we use Node.js for REST services?"
	var cmts2 []Comment
	cmts2 = append(cmts2, c21)
	cmts2 = append(cmts2, c22)
	b2.Comments = &cmts2

	blogs = append(blogs, b1)
	blogs = append(blogs, b2)
}

//GetBlogs GetBlogs
func GetBlogs() *[]Blog {
	return &blogs
}

//GetBlog GetBlog
func GetBlog(id int64) *Blog {
	var rtn Blog
	if id == 1 {
		rtn = blogs[0]
	} else if id == 2 {
		rtn = blogs[1]
	}
	return &rtn
}
