package handlers

import (
	"go-api/internal/database"
	"go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

// create a map (js object), key is a string and value is a Post struct
var posts = make(map[string]models.Post) 
// can add by doing like: posts["1"] = models.Post{ID: "1", Title: "Title 1", Body: "Body 1", Author: "Author 1"}
// post := posts["1"]     // Gets post with ID "1"
// post, exists := posts["1"]  // Also checks if it exists

func SetupPostRoutes(app *fiber.App) { //fiber.App: pointer to Fiber app instance so we modify it by adding the routes
	api := app.Group("/api/posts")

	api.Get("/", GetPosts)
	api.Get("/:id", GetPost)
	api.Post("/", CreatePost)
	api.Put("/:id", UpdatePost)
	api.Delete("/:id", DeletePost)
}

// In memory Version
// the fn returns an error type, if error->return error, if good return nil
// fiber.Ctx is a struct that contains the request and response objects
// func GetPosts(ctx *fiber.Ctx) error {
// 	fmt.Println("Getting all posts")

// 	// var postsList []models.Post //create a slice
// 	postsList := []models.Post{} //create a empty slice

// 	for _, post := range posts {
// 		postsList = append(postsList, post)
// 	}


// 	return ctx.Status(200).JSON(postsList)
// 	// return ctx.JSON([]models.Post{}) // [] declares a slice, similar to array but can grow or shrink, type, {} empty initializer
// }

// func GetPosts(ctx *fiber.Ctx) error {
// 	var posts []models.Post
// 	result := database.DB.Find(&posts)

// 	if result.Error != nil {
// 		return ctx.Status(500).JSON(fiber.Map{
// 			"error": "Error fetching posts",
// 		})
// 	}

// 	return ctx.Status(200).JSON(posts)
// }

func GetPosts(c *fiber.Ctx) error {
    var posts []models.Post
    result := database.DB.Find(&posts)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to fetch posts",
        })
    }
    return c.JSON(posts)
}



// func GetPost(c *fiber.Ctx) error {
//     var id = c.Params("id") //get id from url params

// 	post, exists := posts[id] //check if post exists

// 	if !exists {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Post not found",
// 		})
// 	}

// 	return c.Status(200).JSON(post)
// }


func GetPost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post
    
    result := database.DB.First(&post, id)
    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Post not found",
        })
    }

    return c.JSON(post)
}




// func CreatePost(c *fiber.Ctx) error {
//     post := new(models.Post) //allocate memory for new post struct, post is the pointetr. would be same as var post = &models.Post{}

// 	// This is equivalent to:
//     // err := c.BodyParser(post)  // First assign error to err
//     // if err != nil {           // Then check if error is not nil
// 	//Body.parser will validate if the body matchs Post struct
// 	//  BodyParser fills in its fields with the values from the request body.
//     if err := c.BodyParser(post); err != nil { //reads json from req.body 
//         return c.Status(400).JSON(fiber.Map{
//             "error": "Invalid request body",
//         })
//     }
    
//     posts[post.ID] = *post
//     return c.Status(201).JSON(post)
// }

func CreatePost(c *fiber.Ctx) error {
    post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

	result := database.DB.Create(&post)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to create post",
        })
    }

    return c.Status(201).JSON(post)


}

// func UpdatePost(c *fiber.Ctx) error {
//     var id = c.Params("id") //get id from url params

// 	_, exists := posts[id] //check if post exists

// 	if !exists {
// 		return c.Status(404).JSON(fiber.Map{
// 			"error": "Post not found",
// 		})
// 	}

//     post := new(models.Post)
// 	// BodyParser fills in its fields with the values from the request body.
// 	err := c.BodyParser(post)

// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	post.ID = id

// 	posts[id] = *post

// 	fmt.Println(err, post, *post)


// 	// return c.Status(200).JSON(fiber.Map{ "post": post })
// 	return c.Status(200).JSON(post)
// }


func UpdatePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post

    // Check if post exists
    if result := database.DB.First(&post, id); result.Error != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Post not found",
        })
    }

    // Parse updated data
    if err := c.BodyParser(&post); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    database.DB.Save(&post)

    return c.JSON(post)
}

// func DeletePost(c *fiber.Ctx) error {
//     var id = c.Params("id") //get id from url params

// 	post, exists := posts[id] //check if post exists
// 	fmt.Println(post)
// 	fmt.Println(exists)

// 	if !exists {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Post not found",
// 		})
// 	}

// 	delete(posts, id)

// 	return c.Status(200).JSON(fiber.Map{})
// }

func DeletePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post

    if result := database.DB.First(&post, id); result.Error != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Post not found",
        })
    }
    

    result := database.DB.Delete(&post, id)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to delete post",
        })
    }

    return c.SendStatus(204)
}