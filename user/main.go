package user

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
	res "github.com/jirenius/go-res"
	"golang.org/x/oauth2"
)

// user profile info
type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func HandlerUserService() {

	s := res.NewService("user")
	s.Handle("info",
		res.Access(res.AccessGranted),
		res.GetModel(func(r res.ModelRequest) {
			r.Model(struct {
				Message string `json:"message"`
			}{"Hello, World!"})
		}),

		res.Call("get", func(r res.CallRequest) {
			var p struct {
				OauthToken string `json:"oauth_token,omitempty"`
				LoginID    string `json:"login_id,omitempty"`
			}
			r.ParseParams(&p)

			// Check if the message property was changed
			if p.OauthToken != "" {
				// Update the model
				user, _, err := loadUserInfo(p.OauthToken, p.LoginID)

				if err != nil {
					fmt.Println("Error while loading the user info" + err.Error())
				}
				// Send success response
				r.OK(user)
			} else {
				r.NotFound()
			}
		}),
	)

	fmt.Println("user service is registered")
	err := s.ListenAndServe("nats://localhost:4222")
	if err != nil {
		fmt.Println("Error while registering the user service" + err.Error())
	}
}

func loadUserInfo(token string, LoginID string) (*github.User, *github.Response, error) {
	fmt.Println("User_ID " + LoginID)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client.Users.Get(ctx, LoginID)
}

// get Repo service
func HandlerRepoService() {

	s := res.NewService("repos")
	s.Handle("info",
		res.Access(res.AccessGranted),
		res.GetModel(func(r res.ModelRequest) {
			r.Model(struct {
				Message string `json:"message"`
			}{"please try /api/repos/info/get"})
		}),

		res.Call("get", func(r res.CallRequest) {
			var p struct {
				OauthToken string `json:"oauth_token,omitempty"`
				Owner      string `json:"owner,omitempty"`
				Repo       string `json:"repo,omitempty"`
			}
			r.ParseParams(&p)

			// Check if the message property was changed
			if p.OauthToken != "" {
				// Update the model
				repo, _, err := loadRepo(p.OauthToken, p.Owner, p.Repo)

				if err != nil {
					fmt.Println("Error while loading the user info" + err.Error())
				}
				// Send success response
				r.OK(repo)
			} else {
				r.NotFound()
			}
		}),
	)

	fmt.Println("user service is registered")
	err := s.ListenAndServe("nats://localhost:4222")
	if err != nil {
		fmt.Println("Error while registering the user service" + err.Error())
	}
}

func loadRepo(token string, owner, repo string) (*github.Repository, *github.Response, error) {
	fmt.Println("repo " + owner + "/" + repo)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client.Repositories.Get(ctx, owner, repo)
}
