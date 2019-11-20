package commands

import (
	"github.com/dora1998/snail-bot/mock_repository"
	"github.com/dora1998/snail-bot/mock_twitter"
	"github.com/dora1998/snail-bot/repository"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestCommandHandler_remove(t *testing.T) {
	type args struct {
		body     string
		username string
		statusId int64
	}
	tests := []struct {
		name     string
		args     args
		injector func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient)
	}{
		{
			"normal",
			args{
				body:     "test",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				mr := mock_repository.NewMockRepository(c)
				mr.EXPECT().GetTaskByBody(gomock.Eq("test")).Return(&repository.Task{
					Id:        "hoge",
					Body:      "test",
					Deadline:  time.Time{},
					CreatedAt: time.Time{},
					CreatedBy: "testuser",
				}).Times(1)
				mr.EXPECT().Remove(gomock.Eq("hoge")).Return(nil).Times(1)

				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollwing(gomock.Eq("testuser")).Return(true).Times(1)
				mt.EXPECT().CreateFavorite(gomock.Eq(int64(0))).Times(1)
				return mr, mt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo, twitterClient := tt.injector(ctrl)
			h := &CommandHandler{
				repository:    repo,
				twitterClient: twitterClient,
			}
			h.remove(tt.args.body, tt.args.username, tt.args.statusId)
		})
	}
}
