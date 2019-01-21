package actions

import (
	"github.com/JDWardle/buffalo_spelunking/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Topic)
// DB Table: Plural (topics)
// Resource: Plural (Topics)
// Path: Plural (/topics)
// View Template Folder: Plural (/templates/topics/)

// TopicsResource is the resource for the Topic model
type TopicsResource struct {
	buffalo.Resource
}

// List gets all Topics. This function is mapped to the path
// GET /topics
func (v TopicsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	topics := &models.Topics{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Topics from the DB
	if err := q.All(topics); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, topics))
}

// Show gets the data for one Topic. This function is mapped to
// the path GET /topics/{topic_id}
func (v TopicsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Topic
	topic := &models.Topic{}

	// To find the Topic the parameter topic_id is used.
	if err := tx.Find(topic, c.Param("topic_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, topic))
}

// New renders the form for creating a new Topic.
// This function is mapped to the path GET /topics/new
func (v TopicsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Topic{}))
}

// Create adds a Topic to the DB. This function is mapped to the
// path POST /topics
func (v TopicsResource) Create(c buffalo.Context) error {
	// Allocate an empty Topic
	topic := &models.Topic{}

	// Bind topic to the html form elements
	if err := c.Bind(topic); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(topic)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, topic))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Topic was created successfully")

	// and redirect to the topics index page
	return c.Render(201, r.Auto(c, topic))
}

// Edit renders a edit form for a Topic. This function is
// mapped to the path GET /topics/{topic_id}/edit
func (v TopicsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Topic
	topic := &models.Topic{}

	if err := tx.Find(topic, c.Param("topic_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, topic))
}

// Update changes a Topic in the DB. This function is mapped to
// the path PUT /topics/{topic_id}
func (v TopicsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Topic
	topic := &models.Topic{}

	if err := tx.Find(topic, c.Param("topic_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Topic to the html form elements
	if err := c.Bind(topic); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(topic)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, topic))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Topic was updated successfully")

	// and redirect to the topics index page
	return c.Render(200, r.Auto(c, topic))
}

// Destroy deletes a Topic from the DB. This function is mapped
// to the path DELETE /topics/{topic_id}
func (v TopicsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Topic
	topic := &models.Topic{}

	// To find the Topic the parameter topic_id is used.
	if err := tx.Find(topic, c.Param("topic_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(topic); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Topic was destroyed successfully")

	// Redirect to the topics index page
	return c.Render(200, r.Auto(c, topic))
}
