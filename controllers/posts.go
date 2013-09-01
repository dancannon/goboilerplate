package controllers

import (
	"fmt"
	"github.com/dancannon/gonews/models"
	repository "github.com/dancannon/gonews/repositories/rethinkdb"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

type PostController struct {
}

func (c *PostController) Init(r *mux.Router) {
	r.HandleFunc("/new", c.newHandler)
}

func (c *PostController) newHandler(w http.ResponseWriter, r *http.Request) {
	post := models.Post{
		Author: "User 1",
		Type:   models.PostTypeLink,
		Title:  "Test Post",
		Meta: map[string]string{
			"type":    models.LinkTypeArticle,
			"title":   "Test Link Title",
			"content": "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Nam liber tempor cum soluta nobis eleifend option congue nihil imperdiet doming id quod mazim placerat facer possim assum. Typi non habent claritatem insitam; est usus legentis in iis qui facit eorum claritatem. Investigationes demonstraverunt lectores legere me lius quod ii legunt saepius. Claritas est etiam processus dynamicus, qui sequitur mutationem consuetudium lectorum. Mirum est notare quam littera gothica, quam nunc putamus parum claram, anteposuerit litterarum formas humanitatis per seacula quarta decima et quinta decima. Eodem modo typi, qui nunc nobis videntur parum clari, fiant sollemnes in futurum.",
		},
		UpVotes:   rand.Intn(1000),
		DownVotes: rand.Intn(1000),
		Tags:      []string{"News"},
		Created:   time.Now(),
		Modified:  time.Now(),
	}
	repository.Store(post)

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, "Added new post.")
}
